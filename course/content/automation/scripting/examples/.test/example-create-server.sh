#!/bin/bash

# Tell bash to exit immediately if a command exits with a non-zero status
set -e

# --- Input Variables ---
VM_NAME=$1

# --- Configuration Variables ---
VM_MEMORY="4096" # MB
VM_CPUS="4" # Number of CPUs
VM_DISK_SIZE_MB="25600" # 25 GB in MB
VM_DIR="/c/temp/vms/$VM_NAME"
ISO_PATH="/c/temp/vms/ubuntu-24.04.1-live-server-amd64.iso"
HOST_SHARE_PATH="/c/gitrepos"
NETWORK_NAME="my-website-test"
NAT_NETWORK_CIDR="10.0.0.0/27"
NAT_DHCP_IP="10.0.0.1"
NAT_DHCP_NETMASK="255.255.255.224"
NAT_DHCP_LOWERIP="10.0.0.4"
NAT_DHCP_UPPERIP="10.0.0.30"

# Function to check if a VirtualBox network exists
network_exists() {
    VBoxManage natnetwork list | grep -q "Name: *$1"
}

# Function to check if a VirtualBox VM exists
vm_exists() {
    VBoxManage list vms | grep -q "\"$1\""
}

# Create NAT Network if it doesn't exist
echo "checking NAT network '$NETWORK_NAME'..."
if network_exists "$NETWORK_NAME"; then
    echo "NAT network '$NETWORK_NAME' already exists."
else
    echo "NAT network '$NETWORK_NAME' not found. creating..."
    VBoxManage natnetwork add --netname "$NETWORK_NAME" --network "$NAT_NETWORK_CIDR" --enable --ipv6 off
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create NAT network '$NETWORK_NAME'."
        exit 1
    fi
    VBoxManage dhcpserver add --netname "$NETWORK_NAME" --ip "$NAT_DHCP_IP" --netmask "$NAT_DHCP_NETMASK" --lowerip "$NAT_DHCP_LOWERIP" --upperip "$NAT_DHCP_UPPERIP" --enable
    if [ $? -ne 0 ]; then
        echo "Error: Failed to add DHCP server to '$NETWORK_NAME'."
        exit 1
    fi
    VBoxManage natnetwork modify --netname "$NETWORK_NAME" --dhcp on
    if [ $? -ne 0 ]; then
        echo "Error: Failed to enable DHCP on '$NETWORK_NAME'."
        exit 1
    fi
    
    echo "NAT network '$NETWORK_NAME' created and configured successfully."
fi

if VBoxManage natnetwork list $NETWORK_NAME | grep 'portforward_'> /dev/null; then
    echo "removing existing port forwarding rules..."
    VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4 delete "portforward_ssh"
    VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4 delete "portforward_tls"
    VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4 delete "portforward_web"
fi

echo "adding port forwarding rules..." 
VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4="portforward_ssh:tcp:[]:2222:[10.0.0.5]:22"
VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4="portforward_tls:tcp:[]:8081:[10.0.0.5]:443"
VBoxManage natnetwork modify --netname "$NETWORK_NAME" --port-forward-4="portforward_web:tcp:[]:8080:[10.0.0.5]:80"

echo "checking for VM '$VM_NAME'..."
if vm_exists "$VM_NAME"; then
    echo "VM '$VM_NAME' already exists. updating..."

    # main settings
    VBoxManage modifyvm "$VM_NAME" --memory "$VM_MEMORY" --cpus "$VM_CPUS"
    if [ $? -ne 0 ]; then echo "Warning: Failed to update VM memory/CPUs."; fi

    # network setttings
    VBoxManage modifyvm "$VM_NAME" --nic1 natnetwork --nat-network1 "$NETWORK_NAME"
    if [ $? -ne 0 ]; then echo "Warning: Failed to set VM network adapter."; fi

    echo "removing and re-adding shared folder..."
    VBoxManage sharedfolder remove "$VM_NAME" --name "gitrepos"
    VBoxManage sharedfolder add "$VM_NAME" --name "gitrepos" --hostpath "$HOST_SHARE_PATH" --automount

    echo "VM '$VM_NAME' settings updated successfully."

else
    echo "VM '$VM_NAME' does not exist. creating..."

    mkdir -p "$VM_DIR"
    
    VBoxManage createvm --name "$VM_NAME" --ostype "Linux_64" --register --basefolder "$VM_DIR"
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create VM '$VM_NAME'."
        exit 1
    fi

    VBoxManage modifyvm "$VM_NAME" --memory "$VM_MEMORY" --cpus "$VM_CPUS" --boot1 dvd --boot2 disk --boot3 none --boot4 none
    if [ $? -ne 0 ]; then echo "Warning: Failed to set VM memory/CPUs/boot order."; fi


    VM_DISK_PATH="$VM_DIR/$VM_NAME.vdi"
    VBoxManage createmedium disk --filename "$VM_DISK_PATH" --size "$VM_DISK_SIZE_MB" --format VDI
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create virtual disk '$VM_DISK_PATH'."
        exit 1
    fi

    VBoxManage storagectl "$VM_NAME" --name "SATA Controller" --add sata --controller IntelAhci
    if [ $? -ne 0 ]; then echo "Warning: Failed to add SATA controller."; fi
    VBoxManage storageattach "$VM_NAME" --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium "$VM_DISK_PATH"
    if [ $? -ne 0 ]; then echo "Warning: Failed to attach virtual disk."; fi

    # Attach ISO
    if [ -f "$ISO_PATH" ]; then
        VBoxManage storagectl "$VM_NAME" --name "IDE Controller" --add ide --controller PIIX4
        if [ $? -ne 0 ]; then echo "Warning: Failed to add IDE controller."; fi
        VBoxManage storageattach "$VM_NAME" --storagectl "IDE Controller" --port 0 --device 0 --type dvddrive --medium "$ISO_PATH"
        if [ $? -ne 0 ]; then echo "Warning: Failed to attach ISO '$ISO_PATH'. Please check the path."; fi
    else
        echo "Warning: ISO file '$ISO_PATH' not found. VM created without an attached ISO."
        echo "Please attach the ISO manually via VirtualBox Manager if you intend to install an OS."
    fi

    # Configure Network Adapter
    VBoxManage modifyvm "$VM_NAME" --nic1 natnetwork --nat-network1 "$NETWORK_NAME"
    if [ $? -ne 0 ]; then echo "Warning: Failed to set VM network adapter."; fi

    echo "adding shared folder..."
    VBoxManage sharedfolder add "$VM_NAME" --name "gitrepos" --hostpath "$HOST_SHARE_PATH" --automount

    echo "VM '$VM_NAME' created successfully."
    echo "You can now start the VM: VBoxManage startvm \"$VM_NAME\" --type gui"
    echo "Remember to install your operating system from the attached ISO."
fi

echo "script finished."