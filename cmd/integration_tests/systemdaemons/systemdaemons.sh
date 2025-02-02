#!/bin/sh
set -eu

# Start a system bus
dbus-daemon --config-file=/dbus.conf
sleep 1
export DBUS_SYSTEM_BUS_ADDRESS=unix:path=/dbus/system_bus_socket

mode=${1:-default}

# Set polkit allowance/denial for adsys
polkit_mode=${1}
if [ "${polkit_mode}" != "default" -a "$polkit_mode" != "no"  ]; then
    polkit_mode=yes
fi
if [ "${polkit_mode}" != "default" ]; then
    sed -e "s#<allow_any>.*#<allow_any>${polkit_mode}</allow_any>#" \
        -e "s#<allow_inactive>.*#<allow_inactive>${polkit_mode}</allow_inactive>#" \
        -e "s#<allow_active>.*#<allow_active>${polkit_mode}</allow_active>#" \
    /usr/share/polkit-1/actions.orig/com.ubuntu.adsys.policy > /usr/share/polkit-1/actions/com.ubuntu.adsys.policy
fi

# Handle systemd objects depending on the mode
python3 -m dbusmock --system org.freedesktop.systemd1 /org/freedesktop/systemd1 org.freedesktop.systemd1.Manager &
sleep 1

time=""
case "${mode}" in
  "no_startup_time")
    ;;

  "invalid_startup_time")
    time="<string 'invalid'>"
    ;;

  *)
    time="<uint64 1621860927000000>"
    ;;
esac
if [ -n "${time}" ]; then
    gdbus call --system -d org.freedesktop.systemd1 -o /org/freedesktop/systemd1 -m org.freedesktop.DBus.Mock.AddProperty org.freedesktop.systemd1.Manager GeneratorsStartTimestamp "${time}"
fi

# adsys refresh timer unit
gdbus call --system -d org.freedesktop.systemd1 -o /org/freedesktop/systemd1 -m org.freedesktop.DBus.Mock.AddObject /org/freedesktop/systemd1/unit/adsys_2dgpo_2drefresh_2etimer org.freedesktop.systemd1.Timer "{}" "[]"
time=""
case "${mode}" in
  "no_nextrefresh_time")
    ;;

  "invalid_nextrefresh_time")
    time="<string 'invalid'>"
    ;;

  *)
    time="<uint64 86400000000>"
    ;;
esac
if [ -n "${time}" ]; then
    gdbus call --system -d org.freedesktop.systemd1 -o /org/freedesktop/systemd1/unit/adsys_2dgpo_2drefresh_2etimer  -m org.freedesktop.DBus.Mock.AddProperty org.freedesktop.systemd1.Timer NextElapseUSecMonotonic "${time}"
fi

/usr/lib/policykit-1/polkitd