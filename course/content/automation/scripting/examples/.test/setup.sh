#!/bin/bash

echo "setting up SSH server on the VM..."

sudo apt-get update
sudo apt-get install -y openssh-server

sudo service ssh status
sudo service ssh start
sudo service ssh status

# enable SSH
sudo systemctl enable ssh

# now we can see its enabled:
sudo service ssh status

#TODO install git, runner, nginx 