package adsys_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ubuntu/adsys/cmd/adsysd/client"
	"github.com/ubuntu/adsys/cmd/adsysd/daemon"
)

const dockerPolkitdImage = "docker.pkg.github.com/ubuntu/adsys/polkitd:0.1"

func TestStartAndStopDaemon(t *testing.T) {
	defer polkitAnswer(t, "yes")()

	_, quit := runDaemon(t, true)
	quit()
}

// runDaemon generates the configuration file and the daemon lifecycle.
// It returns the configuration file path and a quit() function.
func runDaemon(t *testing.T, startDaemon bool) (conf string, quit func()) {
	t.Helper()

	dir := t.TempDir()

	// Create config
	confFile := filepath.Join(dir, "adsys.yaml")
	err := os.WriteFile(confFile, []byte(fmt.Sprintf(`
# Service and client configuration
verbose: 2
socket: %s/socket

# Service only configuration
cache_dir: %s/cache
run_dir: %s/run
servicetimeout: 30
ad_server: warthogs.biz
ad_domain: ldap://adc.warthogs.biz
`, dir, dir, dir)), 0644)
	require.NoError(t, err, "Setup: config file should be created")

	if !startDaemon {
		return confFile, func() {}
	}

	var wg sync.WaitGroup
	d := daemon.New()
	defer changeOsArgs(t, confFile)()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := d.Run()
		require.NoError(t, err, "daemon should exit with no error")
	}()

	d.WaitReady()

	return confFile, func() {
		done := make(chan struct{})
		go func() {
			d.Quit()
			close(done)
		}()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("daemon should have stopped within second")
		}

		wg.Wait()
	}
}

// runClient instantiates a client using conf from the given args.
// It returns the stdout content and error from client.
func runClient(t *testing.T, conf string, args ...string) (stdout string, err error) {
	t.Helper()

	c := client.New()
	defer changeOsArgs(t, conf, args...)()

	// capture stdout
	r, w, err := os.Pipe()
	require.NoError(t, err, "Setup: pipe shouldn’t fail")
	orig := os.Stdout
	os.Stdout = w

	err = c.Run()

	// restore and collect
	os.Stdout = orig
	w.Close()
	var out bytes.Buffer
	_, errCopy := io.Copy(&out, r)
	require.NoError(t, errCopy, "Couldn’t copy stdout to buffer")

	return out.String(), err
}

// changeOsArgs modifies the os Args for cobra to parse them successfully.
// It returns a teardown funciton to restore original args.
// As os.Args is global, calling it prevents any parallell testing.
func changeOsArgs(t *testing.T, conf string, args ...string) (teardown func()) {
	t.Helper()

	origArgs := os.Args

	os.Args = []string{"tests", "-vv"}
	if conf != "" {
		os.Args = append(os.Args, "-c", conf)
	}
	if args != nil {
		os.Args = append(os.Args, args...)
	}

	return func() {
		os.Args = origArgs
	}
}

func TestMain(m *testing.M) {
	// Start local polkitd in container with our policy (one for always yes, one for always no)
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		defer runPolkitd()()
	}

	m.Run()
}

var (
	yesSocket string
	noSocket  string
)

// runPolkitd is a helper to start polkit and a system dbus session in two containers:
// - one giving all permissions to any actions
// - one giving no permissions to every actions
// The current branch .policy file is used.
// you can then select the correct daemon via the system dbus socket with polkitAnswer().
// teardown will ensure the containers are stopped.
func runPolkitd() (teardown func()) {
	r, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		log.Fatalf("Setup: couldn't set a random name for docker container: %v", err)
	}
	containerName := fmt.Sprintf("adsys-tests-%06d", r.Int64())

	adsysActionsDir, err := filepath.Abs("../../../internal/adsysservice/actions")
	if err != nil {
		log.Fatalf("Setup: couldn't get absolute path for actions: %v", err)
	}

	dir, err := os.MkdirTemp("/tmp", "adsys-polkitd.*")
	if err != nil {
		log.Fatalf("Setup: failed to create temporary directory: %v", err)
	}

	answers := map[string]string{
		"yes": filepath.Join(dir, "yes"),
		"no":  filepath.Join(dir, "no"),
	}

	var wg sync.WaitGroup
	for answer, socketDir := range answers {
		answer := answer
		socketDir := socketDir
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := os.MkdirAll(socketDir, 0755); err != nil {
				log.Fatalf("Setup: can’t create %s socket directory: %v", answer, err)
			}

			cmd := exec.Command("docker",
				"run", "--rm", "--pid", "host",
				"--name", containerName+answer,
				"--volume", fmt.Sprintf("%s:%s:ro", adsysActionsDir, "/usr/share/polkit-1/actions.orig"),
				"--volume", `/etc/group:/etc/group:ro`,
				"--volume", `/etc/passwd:/etc/passwd:ro`,
				"-v", fmt.Sprintf("%s:/dbus/", socketDir),
				dockerPolkitdImage,
				answer,
			)
			out, _ := cmd.CombinedOutput()
			// Docker stop -t 0 will kill it anyway the container with exit code 143
			if cmd.ProcessState.ExitCode() > 0 && cmd.ProcessState.ExitCode() != 143 {
				log.Fatalf("Error running polkit %s container: %v", answer, string(out))
			}
		}()
	}

	yesSocket = fmt.Sprintf("unix:path=%s/system_bus_socket", answers["yes"])
	noSocket = fmt.Sprintf("unix:path=%s/system_bus_socket", answers["no"])

	// give time for polkit containers to start
	time.Sleep(2 * time.Second)

	return func() {
		defer func() {
			err := os.RemoveAll(dir)
			if err != nil {
				log.Fatalf("Teardown: failed to delete temporary directory: %v", err)
			}
		}()

		for answer := range answers {
			out, err := exec.Command("docker", "stop", "-t", "0", containerName+answer).CombinedOutput()
			if err != nil {
				log.Fatalf("Teardown: can’t stop polkitd container: %v", string(out))
			}
		}
		wg.Wait()
	}
}

// polkitAnswer will flip to which polkit to communicate to:
// - yes for polkit always authorizing our actions.
// - no for polkit always denying our actions.
// Note that this modify the environment variable, and so, tests using them can’t run in parallel.
func polkitAnswer(t *testing.T, answer string) func() {
	t.Helper()

	var socket string
	switch answer {
	case "yes":
		socket = yesSocket
	case "no":
		socket = noSocket
	case "":
		return func() {}
	default:
		t.Fatalf("Setup: unknown polkit answer to support: %s", answer)
	}

	old := os.Getenv("DBUS_SYSTEM_BUS_ADDRESS")
	if err := os.Setenv("DBUS_SYSTEM_BUS_ADDRESS", socket); err != nil {
		t.Fatalf("Setup: couldn't set DBUS_SYSTEM_BUS_ADDRESS: %v", err)
	}

	return func() {
		if err := os.Setenv("DBUS_SYSTEM_BUS_ADDRESS", old); err != nil {
			t.Fatalf("Setup: couldn't set DBUS_SYSTEM_BUS_ADDRESS: %v", err)
		}
	}
}

type runner interface {
	Run() error
}

func TestExecuteCommand(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] != "--" {
			args = args[1:]
			continue
		}
		args = args[1:]
		break
	}

	// let cobra knows what we want to execute
	os.Args = args

	var app runner
	switch args[0] {
	case "adsysctl":
		app = client.New()
	case "adsysd":
		app = daemon.New()
	default:
		fmt.Fprintf(os.Stderr, "UNKNOWN command: %s", args[0])
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

var testCmdName = os.Args[0]

func startCmd(t *testing.T, wait bool, args ...string) (out func() string, teardown func() error, err error) {
	t.Helper()

	cmdArgs := []string{"env", "GO_WANT_HELPER_PROCESS=1", testCmdName, "-test.run=TestExecuteCommand", "--"}
	cmdArgs = append(cmdArgs, args...)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cmd := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...)

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b

	err = cmd.Start()
	if wait {
		err := cmd.Wait()
		cancel()
		return func() string { return b.String() }, func() error { return nil }, err
	}

	return func() string { return b.String() },
		func() error {
			if err := cmd.Process.Kill(); err != nil {
				t.Fatal("Failed to kill process: ", err)
			}
			err := cmd.Wait()
			cancel()
			return err
		}, err
}
