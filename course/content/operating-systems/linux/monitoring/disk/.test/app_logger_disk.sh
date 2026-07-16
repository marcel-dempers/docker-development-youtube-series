#!/bin/bash

# Directory to create large files in
TARGET_DIR="/var/log/app_logs"
mkdir -p "$TARGET_DIR"

# Number of files to create
NUM_FILES=1000

# Size of each file in MB
FILE_SIZE_MB=1

# Create large files
for i in $(seq 1 $NUM_FILES); do
    TIME=$(date  +%Y%m%d%H%M%s%N)
    mkdir -p "$TARGET_DIR/$TIME"
    dd if=/dev/zero of="$TARGET_DIR/$TIME/app_${i}.log" bs=1M count=$FILE_SIZE_MB
done

echo "Disk usage simulation complete. Large files created in $TARGET_DIR."