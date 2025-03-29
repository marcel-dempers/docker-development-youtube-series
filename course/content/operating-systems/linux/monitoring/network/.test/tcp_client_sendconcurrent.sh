#!/bin/bash

echo "Connecting to TCP server on localhost port 12345..."

echo "Sending messages one after another!"
while true; do
  echo "Sending message and closing connection!" | nc -q 5 localhost 12345
done