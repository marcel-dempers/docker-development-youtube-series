#!/bin/bash

# settings
DEPLOYMENT_DEST_DIR="/websites/my-website"

#inputs
PACKAGE_NUMBER=$1

echo "starting our deployment script..."

# Validate Inputs
if [ -z "$PACKAGE_NUMBER" ]; then
  echo "ERROR: PACKAGE_NUMBER is not passed in!"
  exit 1
fi

if [ -z "$DEPLOYMENT_DEST_DIR" ]; then
  echo "ERROR: destination directory is not set!"
  exit 1
fi

echo "extracting package /home/github/packages/package-${PACKAGE_NUMBER}.tar.gz"

tar -C ${PWD} -xzf /home/github/packages/package-${PACKAGE_NUMBER}.tar.gz
ls -la $PWD
DEPLOYMENT_SOURCE_DIR="$PWD/src"

if [ ! -d "$DEPLOYMENT_DEST_DIR" ]; then
  echo "destination directory does not exist, creating... "
  mkdir -p $DEPLOYMENT_DEST_DIR
  echo "destination directory created."  
fi

echo "deploying latest changes..."
NEW_DEPLOYMENT_DIR="$DEPLOYMENT_DEST_DIR/$PACKAGE_NUMBER"

mkdir -p "$NEW_DEPLOYMENT_DIR"

# Copy source files to the new directory
cp -r "$DEPLOYMENT_SOURCE_DIR"/* "$NEW_DEPLOYMENT_DIR"
echo "deployed website to $NEW_DEPLOYMENT_DIR"

echo "deployment complete"