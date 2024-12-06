# 🎬 Introduction to Servers & Virtualization

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 2](../../../chapters/chapter-2-operating-systems/README.md)

## What is a Server

A server is a dedicated computer that computes, stores and manages data or serves a specific purpose </br>
For example, a server could be running a database that stores important data, run a website for a company, or a set of services </br>

Types of servers include:
* Mail server
* Web server
* Database server
* DNS server
* Proxy servers
* Application servers


Servers can either be physical or virtual. </br>
Physical servers are ones that are made up of physical hardware, like the ones in a datacenter. </br>
You can also setup a physical server by using your old desktop hardware, or when you buy a Rasberry Pi or a mini PC like an NUC. </br>

When setting up a server, hardware is important </br> We want hardware that is reliable and long lasting to minimise cost and increase reliability and longevity </br>

For DevOps, SRE and platform engineers the most important part of setting up servers is choosing the hardware resources such as CPU, RAM memory and Hard disk types and sizes. </br>

## Why Servers are important for DevOps , SRE and Cloud Engineers

Throughout a DevOps career, you will be working either directly or indirectly with servers, whether you are working with real on-prem hardware, or servers in the cloud, or even with serverless technology, you will still be indirectly working with servers and will need to consider important aspects of how servers operate. </br>

For example: 
* "How much traffic can this server handle?" (does not matter if its physical or virtual!)
* "How much data can this server store?" 
* "Why is my application on this server so slow?"


<b>"The cloud is just someone else's computer"</b> </br>

Even if this is a cloud product or "server-less" technology where you cannot interact with a server directly, there is always a server somewhere and we will need to consider performance and reliability of our system running on that server </br>

## What is Virtualisation

Not all servers are physical. </br>
A Virtual server (or Virtual Machine) is a software-based server that acts as a real server. </br>
We can run this software on a physical server and run multiple virtual server </br>
So it allows us to slice up and share physical hardware among multiple virtual servers </br> 
If you have a physical server with 24 cores and 64GB of RAM, you could host a few 4 core 8GB RAM virtual servers </br>

### Why run Virtual machines 

There are a number of benefits why Virtual machines are great. </br>
* <b>Isolation</b>: Having seperate servers with seperate functions, like running a Web server and a Database server
* <b>Automation</b>: Using automation software or tools can help recreate servers and roll out updates
* <b>Scalability</b>: Ability to quickly add more servers when demands increase 

## Virtualisation software

There are many types of virtualization software that allows us to create Virtual machines </br>
For simplicity and compatibility, in this course we will use 
[Virtual Box](https://www.virtualbox.org/) as it can run on Windows, Linux and MacOS and its the simplest one to start with </br>

### Requirements

In order to follow along with this module, you will need an operating system with hypervisor capabilities. </br>

For Windows: check if Virtualization is enabled by: 
* Right click your taskbar and open `Task Manager`
* Select the `Performance` window
* Check the right side of the panel, see `Virtualization` if its enabled or not.

For Linux: check if Virtualization is enabled by:
* run `sudo cat /proc/cpuinfo`
* If you see the below flags, Virtualization is enabled in the BIOS
  * `vmx`: Intel VT-x, virtualization support enabled in BIOS.
  * `svm`: AMD SVM, virtualization enabled in BIOS.

You will need at least 2GB extra RAM and 25GB of hard disk space to create a Virtual machine for this module </br>

<i>These numbers are early access guide only and may be reviewed and adjusted based on course feedback</i>

### Installation

### Windows: 

In this module we will walk through the installation on both Windows & Linux. </br>
for Windows its as simple as following the installer prompts </br>

### Linux

For Linux, the installation can be slightly more complicated. </br>
* Virtualization needs to be enabled as per above Requirements </br>
* You need latest kernel modules for your computer
* You either need to turn off Secure boot in the BIOS, or you need to manually sign the VirtualBox kernel modules by using `mokutil`

Install Virtual Box on Ubuntu:

```
# install a few dependencies
sudo apt-get update 
sudo apt-get install -y ca-certificates curl wget gnupg

# download oracle public key
wget -O- https://www.virtualbox.org/download/oracle_vbox_2016.asc | sudo gpg --yes --output /usr/share/keyrings/oracle-virtualbox-2016.gpg --dearmor

# get your ubuntu codename for your source list
cat /etc/os-release | grep CODENAME

# update apt sources
sudo echo "deb [arch=amd64 signed-by=/usr/share/keyrings/oracle-virtualbox-2016.gpg] https://download.virtualbox.org/virtualbox/debian jammy contrib" >> /etc/apt/sources.list.d/virtualbox.list

sudo apt-get update 
sudo apt-get install -y virtualbox-7.1

```

#### Preferences

On the Preferences page, the most important thing is to select a folder where all our Virtual Machine data will be stored </br>

## Creating our first Server 

In order to create a virtual server, we need to get an [ISO](https://simple.wikipedia.org/wiki/ISO_image) file </br>

What is an ISO image ? 

Traditionally, an ISO image file is an image of an optical disk. In the past, we could create copies of CD's or DVD disks by creating ISO images of them and store them as files </br>

For Virtual machines, ISO image files are basically the disks that contains the operating system files that we use to install the operating system onto a Virtual Server. </br>

### Download the virtual machine ISO file 

In this course we will get Ubuntu Linux from [ubuntu.com](https://ubuntu.com/) </br>

### Setup Network

One very important component of environments both for in data centers and in the cloud is the network. </br> 
Servers communicate with one another over the network. </br>
Networks can be either <b>public</b> (like sitting on the internet) or <b>private</b> or a combination of public and private </br>

* Host-only Networks
  `This can be used to create a network containing the host and a set of virtual machines, without the need for the host's physical network interface`
* NAT Networks `Network Address Translation (NAT) is the simplest way of accessing an external network from a virtual machine. Usually, it does not require any configuration on the host network and guest system. For this reason, it is the default networking mode in Oracle VM VirtualBox.`
* Cloud Networks 
`This can be used to connect a local VM to a subnet on a remote cloud service.`

In order for a server to communicate with another server it needs to belong to a network, and in order to belong to a network, a server needs an IP address </br>
An IP is a series of numbers and dots that identify a device that is part of a network. </br>
Networks have an IP range. </br> 

We can specify a network range using whats called a [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) <br>
It's a way for us to "size" our network and know how many IP addresses are available in that network. </br>

There are many online network size calculators out there, but I like to use [mxtoolbox](https://mxtoolbox.com/subnetcalculator.aspx)

For example, we can create a network with IP starting with `10.0.0.0` and let's say we want 32 IPs , we select `/27` </br>

Our calculator will show us that this gives us 32 IPs </br>
Our first IP is 10.0.0.0 and we can increment the number at the end all the way up to 31, which in total means that we have 32 IP addresses. </br>

Private IPv4 address ranges include:
* 10.0.0.0 to 10.255.255.255
* 172.16.0.0 to 172.31.255.255
* 192.168.0.0 to 192.168.255.255

Private network ranges are blocked from being used on public networks. </br>

A network can have a really large IP range, because you may want to divide networks into many smaller networks, for example a network for each department or subset of servers. </br>
These divides are called [Subnets](https://en.wikipedia.org/wiki/Subnet). </br>

When a device connects to a network its assigned an IP address by [DHCP](https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol) </br>

`Port Forwarding` allows us to forward certain ports from our machine to this virtual machine network. This is because the NAT Network is isolated so we need to setup any connections between us and the virtual machine </br>

Now to create the network, we can briefly walk through the VirtualBox UI and look at the default NAT Network that is created. </br>
In this walk through, I'd like you to get a bit more of an understanding of networks, so I would prefer if we create our own and not rely on the default one </br> The default network could be subject to change which may cause it to break in the future. </br>
To create a network and setup a DHCP for it, we will use the `VBoxManage` command line. </br>

On Linux, if you have VirtualBox installed you can simply type VBoxManage in a terminal window and you should see its output. </br>

Let's create our NAT Network and a DHCP for it using the following commands: 

<i>For Windows, you will need to add the following directory to your PATH environemnt variable: `C:\Program Files\Oracle\VirtualBox`</i>
<i>Once done, open a new PowerShell terminal</i>
```
VBoxManage natnetwork add --netname "my-website" --network "10.0.0.0/27" --enable --ipv6 off
VBoxManage dhcpserver add --netname "my-website" --ip "10.0.0.1" --netmask "255.255.255.224" --lowerip "10.0.0.4" --upperip "10.0.0.30" --enable
VBoxManage natnetwork modify --netname "my-website" --dhcp on
```
   
### Creating Virtual Machine 

On the main panel, we can click on `Machine` and `New` to get to the server creation page. </br>
give our server a `Name` and `Folder` where our virtual server data will be created. </br>
We select the `ISO Image` that we downloaded previously. </br> Once we select our ISO file, the rest gets filled out because the software gets the data from the ISO image. </br>

Next important section is `Hardware` </br>

#### Hardware 

We select our `Base Memory` (RAM Memory) and `Processors` (CPU) </br>
We learnt about the importance of these resources in our Operating system introduction module </br>

When we choose these resources, we have to consider how much resources our host computer has. </br>

Next we create a hard disk, by filling out `Hard Disk File Location and Size`

Then we press `Finish` to create the virtual server </br>

Now before starting the virtual server, we need to add it to our network we created. </br> We can do this from the settings page </br>

## Setup Server Access

Now Virtual Box allows us to access our server by giving us a screen interface. </br>
Now if we wanted to automate things later and remotly manage this server, we need to setup [SSH](https://en.wikipedia.org/wiki/Secure_Shell) access

Get the local IP to setup port forwarding in Virtual Box.
For this demo, we can set a non-standard port just incase you already have port `22` used on your system. </br>
Therefore the host port we will set it to `2222` and the Guest port (which is the virtual server) will set it as `22`

```
# grab the IP address from the `inet` section
ip addr
```
The IP address of the virtual server will be the "Guest IP" in the port forward rule. </br>
This means all traffic on our system from port `2222` will not be forwarded to port `22` of our virtual server.

On our Virtual server, we will need to install SSH server. </br>

```
sudo apt-get update
sudo apt-get install -y openssh-server
```

Ensure that the SSH server is up and running 

```
sudo service ssh status
sudo service ssh start
sudo service ssh status
```

We can now see the SSH service is running, but is disabled. So when our VM reboots it will not automatically start up. To make the SSH service start up automatically, we need to enable it. Let's do that

```
# enable SSH
sudo systemctl enable ssh

# now we can see its enabled:
sudo service ssh status
```

Now on our system, we can open up Visual Studio code terminal and SSH to our Virtual server:

```
ssh -p 2222 <username>@127.0.0.1
```
