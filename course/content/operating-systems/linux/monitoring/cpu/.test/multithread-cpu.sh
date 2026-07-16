#!/bin/bash

# Function to simulate CPU usage
simulate_cpu_usage() {
  periodInSeconds=$1
  core=$2
  echo "Simulating CPU usage for $periodInSeconds seconds on core $core..."
  end=$((SECONDS + $periodInSeconds))
  while [ $SECONDS -lt $end ]; do
    # Perform a CPU-intensive task directly in the script
    for i in {1..10000}; do
      : # No-op command, just to keep the CPU busy
    done
  done
  echo "CPU usage simulation complete on core $core."
}

# Number of CPU cores
num_cores=$(nproc)

# Duration for the simulation
duration=500

# Run the function in the background for each core
for core in $(seq 1 $num_cores); do
  simulate_cpu_usage $duration $core &
done

# Wait for all background jobs to complete
wait