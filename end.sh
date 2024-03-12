#!/bin/bash

startPort=10110
endPort=10120

for port in $(seq $startPort $endPort); do
    pid=$(lsof -ti tcp:$port)

    if [[ ! -z "$pid" ]]; then
        echo "Killing process on port $port with PID $pid"
        kill -9 $pid
    else
        echo "No process found on port $port"
    fi
done
