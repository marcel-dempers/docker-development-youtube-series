#!/bin/bash

echo "Installing kind from skill script..."

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.25.0/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/local/bin/kind
