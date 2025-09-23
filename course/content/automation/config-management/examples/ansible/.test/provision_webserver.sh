#!/bin/bash

#Install NGINX
###################################################################################

# Check if NGINX is already installed.
if dpkg -s nginx >/dev/null 2>&1; then
    echo "NGINX is already installed. Skipping installation."
else
  sudo apt update
  sudo apt install -y curl gnupg2 ca-certificates lsb-release ubuntu-keyring

  curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor \
      | sudo tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null

  gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg

  echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] \
  http://nginx.org/packages/ubuntu `lsb_release -cs` nginx" \
      | sudo tee /etc/apt/sources.list.d/nginx.list

  echo -e "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900\n" \
      | sudo tee /etc/apt/preferences.d/99nginx

  sudo apt update
  sudo apt install -y nginx
fi

# Setup Systemd override for NGINX
###################################################################################

# Define the directory for the override file
OVERRIDE_DIR="/etc/systemd/system/nginx.service.d"
sudo mkdir -p "$OVERRIDE_DIR"

sudo tee "$OVERRIDE_DIR/override.conf" > /dev/null <<EOF
[Service]
Environment="CONFFILE=/websites/my-website/nginx.conf"
EOF

sudo systemctl daemon-reload
echo "Systemd override file for NGINX created successfully."

# Configure NGINX
###################################################################################

echo "creating website folder..."
sudo mkdir -p /websites/my-website/
sudo cp -r /home/vagrant/tmp/my-website/** /websites/my-website/

echo "setting permissions for files..."
sudo chown -R nginx:nginx /websites/my-website

# Start NGINX
###################################################################################
sudo systemctl start nginx
echo "NGINX started successfully."