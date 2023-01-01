#!/bin/bash
 
WORKSPACE=$(cd $(dirname $0)/ || exit; pwd)
cd "$WORKSPACE" || exit

mkdir -p var

app=sagoo-admin
pidfile=var/$app.pid
logfile=var/$app.log

function check_pid() {
    if [ -f $pidfile ];then
        pid=$(cat $pidfile)
        if [[ -n $pid ]]; then
            running=$(ps -p "$pid"|grep -c -v "PID TTY")
            return "$running"
        fi
    fi
    return 0
}


 function start(){
 	check_pid
 	running=$?
 	if [ $running -gt 0 ]; then
 		echo -n "$app now is running already,pid="
 		cat $pidfile
 		return
 	fi

    nohup ./$app &> $logfile &
    sleep 1
    running=$(ps -p $! | grep -c -v "PID TTY")
    if [ "$running" -gt 0 ];then
        echo $! > $pidfile
        echo "$app started..., pid=$!"
    else
        echo "$app failed to start"
        return 1
    fi

 }

function stop() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        pid=$(cat $pidfile)
        kill "$pid"
        rm -f $pidfile
        echo "$app stoped"
    else
        echo "$app already stoped"
    fi
}

function restart() {
    stop
    sleep 1
    start
}

function status() {
    check_pid
    running=$?
    if [ $running -gt 0 ];then
        echo "started"
    else
        echo "stoped"
    fi
}

function tailf() {
    tail -f var/*
}

function help() {
    echo "$0 pid|start|stop|restart|status|tail"
}

function pid() {
    cat $pidfile
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "stop" ];then
    stop
elif [ "$1" == "start" ];then
    start
elif [ "$1" == "restart" ];then
    restart
elif [ "$1" == "status" ];then
    status
elif [ "$1" == "tail" ];then
    tailf
elif [ "$1" == "pid" ];then
	pid
else
    help
fi
