#!/bin/bash

WORKSPACE=$(cd "$(dirname "$0")" || exit; pwd)
cd "$WORKSPACE" || exit

readonly app='sagooiot'
readonly pidfile="var/$app.pid"
readonly logfile="var/$app.log"

mkdir -p var

check_pid() {
    if [[ -f $pidfile ]]; then
        local pid=$(cat "$pidfile")
        if [[ -n $pid ]]; then
            local running=$(ps -p "$pid" | grep -c -v "PID TTY")
            return "$running"
        fi
    fi
    return 0
}

start() {
    check_pid
    local running=$?
    if [[ $running -gt 0 ]]; then
        printf "%s now is running already, pid: %s\n" "$app" "$(cat "$pidfile")"
        return 1
    fi

    nohup "./$app" >> "$logfile" 2>&1 &
    sleep 1
    running=$(ps -p $! | grep -c -v "PID TTY")
    if [[ $running -gt 0 ]]; then
        echo $! > "$pidfile"
        printf "%s started... pid: %s\n" "$app" "$!"
    else
        printf "%s failed to start.\n" "$app"
        return 1
    fi
}

stop() {
    check_pid
    local running=$?
    if [[ $running -gt 0 ]]; then
        local pid=$(cat "$pidfile")
        kill "$pid"
        rm -f "$pidfile"
        printf "%s stopped.\n" "$app"
    else
        printf "%s is not running.\n" "$app"
    fi
}

restart() {
    stop
    sleep 1
    start
}

status() {
    check_pid
    local running=$?
    if [[ $running -gt 0 ]]; then
        printf "%s is started.\n" "$app"
    else
        printf "%s is stopped.\n" "$app"
    fi
}

tailf() {
    tail -f var/*
}

print_help() {
    printf "Usage: %s {start|stop|restart|status|tail|pid}.\n" "$0"
}

print_pid() {
    cat "$pidfile"
}

main() {
    case "$1" in
        "start") start ;;
        "stop") stop ;;
        "restart") restart ;;
        "status") status ;;
        "tail") tailf ;;
        "pid") print_pid ;;
        *) print_help ;;
    esac
}

main "$@"