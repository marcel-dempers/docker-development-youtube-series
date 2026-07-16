#!/bin/bash

# Memory leak simulation script

echo "Memory leak simulation script started with PID $$"

# Function to simulate memory leak
simulate_memory_leak() {
    local count=0
    local data=()
    while true; do
        # Append random data to the array
        data+=($(head -c 1M < /dev/urandom | base64))
        count=$((count + 1))
        echo "Allocated $count MB of memory, consuming more memory..."
        # Sleep for a short period to slow down the leak
        sleep 1
    done
}


# Start the memory leak simulation
simulate_memory_leak
