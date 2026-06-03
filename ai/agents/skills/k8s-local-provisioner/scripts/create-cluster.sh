#!/bin/bash

echo "Creating Kubernetes cluster from skill script..."

kind create cluster --name technical-writer --image kindest/node:v1.34.3
