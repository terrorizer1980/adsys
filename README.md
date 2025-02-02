# adsys
Active Directory bridging toolset

[![Code quality](https://github.com/ubuntu/adsys/workflows/QA/badge.svg)](https://github.com/ubuntu/adsys/actions?query=workflow%3AQA)
[![Code coverage](https://codecov.io/gh/ubuntu/adsys/branch/master/graph/badge.svg)](https://codecov.io/gh/ubuntu/adsys)
[![Go Reference](https://pkg.go.dev/badge/github.com/ubuntu/adsys.svg)](https://pkg.go.dev/github.com/ubuntu/adsys)
[![Go Report Card](https://goreportcard.com/badge/ubuntu/adsys)](https://goreportcard.com/report/ubuntu/adsys)
[![License](https://img.shields.io/badge/License-GPL3.0-blue.svg)](https://github.com/ubuntu/adsys/blob/master/LICENSE)

## Usage

### User commands

#### adsysctl

AD integration client

##### Synopsis

Active Directory integration bridging toolset command line tool.

```
adsysctl COMMAND [flags]
```

##### Options

```
  -c, --config string   use a specific configuration file
  -h, --help            help for adsysctl
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl applied

Print last applied GPOs for current or given user/machine

##### Synopsis

Alias of "policy applied"

```
adsysctl applied [USER_NAME] [flags]
```

##### Options

```
  -a, --all        show overridden rules in each GPOs.
      --details    show applied rules in addition to GPOs.
  -h, --help       help for applied
      --no-color   don't display colorized version.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl completion

Generates bash completion scripts

##### Synopsis

To load completion run

. <(adsysctl completion)

To configure your bash shell to load completions for each session add to your ~/.bashrc or ~/.profile:

. <(adsysctl completion)


```
adsysctl completion [flags]
```

##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl doc

Documentation

```
adsysctl doc [CHAPTER] [flags]
```

##### Options

```
  -d, --dest string     Write documentation file(s) to this directory.
  -f, --format string   Format type (markdown, raw or html). (default "markdown")
  -h, --help            help for doc
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy

Policy management

```
adsysctl policy COMMAND [flags]
```

##### Options

```
  -h, --help   help for policy
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy admx

Dump windows policy definitions

```
adsysctl policy admx lts-only|all [flags]
```

##### Options

```
      --distro string   distro for which to retrieve policy definition. (default "Ubuntu")
  -h, --help            help for admx
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy applied

Print last applied GPOs for current or given user/machine

```
adsysctl policy applied [USER_NAME] [flags]
```

##### Options

```
  -a, --all        show overridden rules in each GPOs.
      --details    show applied rules in addition to GPOs.
  -h, --help       help for applied
      --no-color   don't display colorized version.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy update

Updates/Create a policy for current user or given user with its kerberos ticket

```
adsysctl policy update [USER_NAME KERBEROS_TICKET_PATH] [flags]
```

##### Options

```
  -a, --all       all updates the policy of the computer and all the logged in users. -m or USER_NAME/TICKET cannot be used with this option.
  -h, --help      help for update
  -m, --machine   machine updates the policy of the computer.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service

Service management

```
adsysctl service COMMAND [flags]
```

##### Options

```
  -h, --help   help for service
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service cat

Print service logs

```
adsysctl service cat [flags]
```

##### Options

```
  -h, --help   help for cat
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service status

Print service status

```
adsysctl service status [flags]
```

##### Options

```
  -h, --help   help for status
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl service stop

Requests to stop the service once all connections are done

```
adsysctl service stop [flags]
```

##### Options

```
  -f, --force   force will shut it down immediately and drop existing connections.
  -h, --help    help for stop
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl update

Updates/Create a policy for current user or given user with its kerberos ticket

##### Synopsis

Alias of "policy update"

```
adsysctl update [USER_NAME KERBEROS_TICKET_PATH] [flags]
```

##### Options

```
  -a, --all       all updates the policy of the computer and all the logged in users. -m or USER_NAME/TICKET cannot be used with this option.
  -h, --help      help for update
  -m, --machine   machine updates the policy of the computer.
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl version

Returns version of client and service

```
adsysctl version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd

AD integration daemon

##### Synopsis

Active Directory integration bridging toolset daemon.

```
adsysd COMMAND [flags]
```

##### Options

```
  -D, --ad-domain string   AD domain to use. Empty to let ADSys parsing sssd.conf.
  -S, --ad-server string   URL of the Active Directory server. Empty to let ADSys parsing sssd.conf.
      --cache-dir string   directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string      use a specific configuration file
  -h, --help               help for adsysd
      --run-dir string     directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string      socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int        time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count      issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd completion

Generates bash completion scripts

##### Synopsis

To load completion run

. <(adsysd completion)

To configure your bash shell to load completions for each session add to your ~/.bashrc or ~/.profile:

. <(adsysd completion)


```
adsysd completion [flags]
```

##### Options

```
  -h, --help   help for completion
```

##### Options inherited from parent commands

```
  -D, --ad-domain string   AD domain to use. Empty to let ADSys parsing sssd.conf.
  -S, --ad-server string   URL of the Active Directory server. Empty to let ADSys parsing sssd.conf.
      --cache-dir string   directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string      use a specific configuration file
      --run-dir string     directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string      socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int        time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count      issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysd version

Returns version of service and exits

```
adsysd version [flags]
```

##### Options

```
  -h, --help   help for version
```

##### Options inherited from parent commands

```
  -D, --ad-domain string   AD domain to use. Empty to let ADSys parsing sssd.conf.
  -S, --ad-server string   URL of the Active Directory server. Empty to let ADSys parsing sssd.conf.
      --cache-dir string   directory where ADsys caches GPOs downloads and policies. (default "/var/cache/adsys")
  -c, --config string      use a specific configuration file
      --run-dir string     directory where ADsys stores transient information erased on reboot. (default "/run/adsys")
  -s, --socket string      socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int        time in seconds without activity before the service exists. 0 for no timeout. (default 120)
  -v, --verbose count      issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

### Hidden commands

Those commands are hidden from help and should primarily be used by the system or for debugging.

#### adsysctl policy debug

Debug various policy infos

```
adsysctl policy debug [flags]
```

##### Options

```
  -h, --help   help for debug
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

#### adsysctl policy debug gpolist-script

Write GPO list python embeeded script in current directory

```
adsysctl policy debug gpolist-script [flags]
```

##### Options

```
  -h, --help   help for gpolist-script
```

##### Options inherited from parent commands

```
  -c, --config string   use a specific configuration file
  -s, --socket string   socket path to use between daemon and client. Can be overridden by systemd socket activation. (default "/run/adsysd.sock")
  -t, --timeout int     time in seconds before cancelling the client request when the server gives no result. 0 for no timeout. (default 30)
  -v, --verbose count   issue INFO (-v), DEBUG (-vv) or DEBUG with caller (-vvv) output
```

