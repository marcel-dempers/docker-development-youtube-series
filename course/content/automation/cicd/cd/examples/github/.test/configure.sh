#!/bin/bash

# settings
DEPLOYMENT_DEST_DIR="/websites/my-website"

#inputs
PACKAGE_NUMBER=$1

echo "starting our configuration script..."

# Validate Inputs
if [ -z "$PACKAGE_NUMBER" ]; then
  echo "ERROR: PACKAGE_NUMBER is not passed in!"
  exit 1
fi

if [ -z "$DEPLOYMENT_DEST_DIR" ]; then
  echo "ERROR: destination directory is not set!"
  exit 1
fi


NEW_DEPLOYMENT_DIR="$DEPLOYMENT_DEST_DIR/$PACKAGE_NUMBER"

echo "updating website configuration..."

CONFIG_FILE="$DEPLOYMENT_DEST_DIR/nginx.conf"

if [ -f "$CONFIG_FILE" ]; then
  echo "updating website configuration..."
  sed -i "s|root $DEPLOYMENT_DEST_DIR.*;|root $NEW_DEPLOYMENT_DIR;|" "$CONFIG_FILE"
  echo "website configuration updated."
fi
echo "configuration complete"