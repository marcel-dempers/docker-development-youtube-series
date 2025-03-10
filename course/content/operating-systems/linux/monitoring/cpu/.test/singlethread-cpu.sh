#!/bin/bash

# Function to simulate CPU usage
simulate_cpu_usage() {
  periodInSeconds=$1
  echo "Simulating CPU usage for $periodInSeconds seconds..."
  end=$((SECONDS + $periodInSeconds))
  while [ $SECONDS -lt $end ]; do
    # Perform a CPU-intensive task directly in the script
    for i in {1..10000}; do
      : # No-op command, just to keep the CPU busy
    done
  done
  echo "CPU usage simulation complete."
}

# Call the function with the provided duration
simulate_cpu_usage 500


#ps -o pid,ppid,cmd --forest -p <PID>