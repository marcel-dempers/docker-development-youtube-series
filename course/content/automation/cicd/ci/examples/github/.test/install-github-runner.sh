#!/bin/bash

#inputs
PAT_TOKEN=$1
GITHUB_ORG=$2
GITHUB_REPO=$3
RUNNER_NAME=$4
RUNNER_FOLDER=$RUNNER_NAME

# settings 
REPO_URL="https://github.com/$GITHUB_ORG/$GITHUB_REPO"
WEBSITE_LOCATION="/websites/my-website" #we need to ensure user has permission for this location

echo "setting up github user if it doesnt exist..."

if sudo id "github" &>/dev/null; then
    echo "user 'github' already exists."
else
    echo "user 'github' does not exist. Creating..."
    sudo useradd -m -s "$SHELL" -c "github actions runner" -U "github"
fi

echo "setup user permissions for website location $WEBSITE_LOCATION..."
sudo mkdir -p $WEBSITE_LOCATION
sudo apt-get update && sudo apt-get install -y acl
sudo setfacl -m u:github:rwx $WEBSITE_LOCATION

echo "allow github user to reload NGINX" 
SUDOERS_FILE="/etc/sudoers.d/github"
SUDOERS_CONTENT="
# Permissions for the GitHub Actions runner user
github ALL=(root) NOPASSWD: /usr/bin/sed, /usr/bin/grep, /usr/sbin/nginx -s reload
"

# Write the content to a temporary file
sudo echo "$SUDOERS_CONTENT" > /tmp/github
sudo chown root:root /tmp/github

# Set the correct permissions on the temporary file
sudo chmod 0440 /tmp/github

# Move the file into the sudoers.d directory
sudo mv -f /tmp/github "$SUDOERS_FILE"

RUNNER_ARCH="linux-x64"
LATEST_RELEASE_DATA=$(curl -s "https://api.github.com/repos/actions/runner/releases/latest")
RUNNER_VERSION=$(echo "$LATEST_RELEASE_DATA" | jq -r '.tag_name')
DOWNLOAD_URL=$(echo "$LATEST_RELEASE_DATA" | jq -r ".assets[] | select(.name | contains(\"${RUNNER_ARCH}\") and contains(\"tar.gz\")) | .browser_download_url")

# Basic error checking
if [ -z "$RUNNER_VERSION" ] || [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not retrieve latest runner version or download URL."
    exit 1
fi

RUNNER_FILENAME="actions-runner-${RUNNER_ARCH}-${RUNNER_VERSION#v}.tar.gz"

echo "Latest Runner Version: $RUNNER_VERSION"
echo "Download URL: $DOWNLOAD_URL"
echo "Filename: $RUNNER_FILENAME"

#create the runner directory
sudo -u github mkdir -p /home/github/$RUNNER_FOLDER

# Check if the file already exists before downloading
if [[ -f "/home/github/$RUNNER_FOLDER/$RUNNER_FILENAME" ]]; then
    echo "File $RUNNER_FILENAME already exists. Skipping download."
else
    echo "Downloading $RUNNER_FILENAME..."
    sudo -u github curl -o "/home/github/$RUNNER_FOLDER/$RUNNER_FILENAME" -L "$DOWNLOAD_URL"
fi

echo "Extracting $RUNNER_FILENAME..."
sudo -u github tar xzf "/home/github/$RUNNER_FOLDER/$RUNNER_FILENAME" -C "/home/github/$RUNNER_FOLDER/"

if [[ -f "/home/github/.installed_dependencies" ]]; then
  echo "github dependencies script already executed. Skipping."
else 
  echo "github .installed_dependencies not found. Proceeding to install dependencies..."

  sudo /home/github/$RUNNER_FOLDER/bin/installdependencies.sh
  if [ $? -eq 0 ]; then
      echo "github runner dependencies installed successfully."
      sudo touch /home/github/.installed_dependencies
  else
      echo "Error installing dependencies. Exiting."
      exit 1
  fi
fi

sudo chown -R github:github /home/github/$RUNNER_FOLDER

if [[ -f "/home/github/.installed_runner" ]]; then
  echo "github runner installed. Skipping."
else 
  echo "github .installed_runner not found. Proceeding to install runner..."

  REG_TOKEN=$(curl -X POST -H "Authorization: token $PAT_TOKEN" \
     -H "Accept: application/vnd.github+json" \
     "https://api.github.com/repos/$GITHUB_ORG/$GITHUB_REPO/actions/runners/registration-token" | jq -r .token)
  
  if [[ -z "$REG_TOKEN" || "$REG_TOKEN" == "null" ]]; then
    echo "ERROR: The GitHub runner registration token could not be fetched. The variable is empty." >&2
    exit 1
  fi

  sudo -u github bash /home/github/$RUNNER_FOLDER/config.sh --unattended \
    --token "$REG_TOKEN" \
    --url "$REPO_URL" \
    --name "$RUNNER_NAME" \
    --labels "self-hosted" "linux" \
    --work "/home/github/$RUNNER_FOLDER/_work" \
    --replace \
    --runasservice \
    --runnergroup "Default"
  
  if [[ -f "/home/github/.installed_runner_service" ]]; then
    echo "github runner systemd service already installed. Skipping."
  else
    cd /home/github/$RUNNER_FOLDER && sudo ./svc.sh install github && sudo ./svc.sh start
    sudo touch /home/github/.installed_runner_service
  fi
  
  sudo touch /home/github/.installed_runner

fi