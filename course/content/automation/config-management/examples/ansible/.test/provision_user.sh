#!/bin/bash

SERVER_USERNAME=$1
SERVER_PASSWORD=$2

if ! id "$SERVER_USERNAME" >/dev/null 2>&1; then
  echo "User $SERVER_USERNAME not found. Creating..."
  sudo useradd -m -s /bin/bash "$SERVER_USERNAME"
  echo "$SERVER_USERNAME:$SERVER_PASSWORD" | sudo chpasswd
fi

echo "Adding user to sudoers group..."
usermod -aG sudo $SERVER_USERNAME