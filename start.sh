#!/bin/bash

start_port=10110
end_port=10120

for port in $(seq $start_port $end_port)
do
   echo "Starting go program on port $port"
   go run . "$port" &
done

echo "All programs have been started."