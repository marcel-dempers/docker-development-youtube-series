# 🎬 Declarative Tools Example: Automate our Infrastructure (Vagrant)

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in automation & deployments as well as provisioning infrastructure over many years. </br>

In this course, so far, you've lived through a bit of automation using scripts and we've highlighted several times the downsides and challenges of writing scripts. </br>

Scripting is a sequence of "Imperative" commands, which are executed line by line. Each line changes the state of the system. </br>

Scripting is extremely flexibe and powerful because you can change a system in many ways but comes with drawbacks because of its imeperative nature. </br>

If a command fails, the system may be left in a broken state. This does not scale well when we run commands over many servers in our infrastructure. This is where automation tools help as you will learn in this module. </br>

Automation tools use concepts we will learn about today which attempts to overcome the challenges of Imperative automation scripting. 

Instead of using Imperative scripting which focuses on the "How", We use Declarative tools which focuses on the "What". So "What" do you want to create instead of "How" do you want to create it. 


## Declarative vs Imperative Automation

The simplest way to think about it is that imperative automation focuses on "how you do something", while declarative automation focuses on "what you want to achieve."

**Imperative Script:**

Let's take a look at an Imperative script:

("How to create our virtual server")

```
#!/bin/bash

# 1. Create the virtual machine
VBoxManage createvm --name "my-vm" --ostype "Ubuntu_64" --register

# 2. Add a virtual hard disk with a specific size (e.g., 10GB)
VBoxManage createmedium disk --filename "my-vm-disk.vdi" --size 10240 --format VDI

# 3. Attach the disk to the VM
VBoxManage storageattach "my-vm" --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium "my-vm-disk.vdi"

# 4. Start the VM
VBoxManage startvm "my-vm"
```

There are some key disadvantages with this.

* **Not Idempotent**: Running the script twice can cause an error or a different outcome. The virtual server may already exist.

* **State Management**: The script doesn't know the current state of the server or infrastructure, so it can't fix things that have changed.

* **Complexity**: As workflows get more complex, imperative scripts become long and hard to maintain.

The outcome of Imperative scripts are not garanteed. </br>

Let's see what that would look like if we used a Declarative tool like Vagrant:

```
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/noble"
  
  config.vm.hostname = "my-vm"

  config.disksize.size = "10GB"

  config.vm.network "private_network", ip: "10.0.0.5"
end
```

![diagram](./declarative.drawio.svg)

Declarative tools provide some key advantages over Imperative scripting. 

* **Repeatable** and complex deployments where you need consistency.

* **Infrastructure as Code (IaC)**: Declarative patterns provide capabilities for managing servers, networks, and cloud resources.

* **Ensuring Idempotence**: You can run the automation multiple times, and it will only make changes if the system doesn't match the **desired state**.

Two important concepts come out of this, Desired state and Infrastructure as Code

## Desired state

With Imperative scripting, I mentioned it focuses on the "How". </br> "How do I create the server with a series of commands?" When we look at scripting, we realized that it cannot manage the outcome or the "state" of a system. </br>
Looking back at the server creation script example, there are 4-5 different states in that script with each line changing the state of our system. There are also unknown states, if any command fails, therefore there could be 8 states as a possible outcome. </br>
This is not great. 

Our requirement when automating our server creation is that we want: 

* A repeatable process that creates our virtual server that has everything installed to run our web site with full CI/CD deployment capabilities. 
* If we run our automation several times, we want the exact same outcome. 
* IF we ran our automation tool for 5 servers, we want them all exactly THE SAME

In our scripting example we see there could be 8 states as an outcome. We only desire one, and that is our "Desired State".

Declarative tools aim to achieve desired state that we define in a file declaratively. </br>

![diagram](./declarative-state.drawio.svg)

## Infrastructure as code

The core idea of IaC is that you define the desired end state of your infrastructure in a configuration file, not the step-by-step process of how to get there. </br>
This could be achieved with a programming language like C#, Python or Javascript, or it could be achieved by a configuration language. </br>
For example, Terraform, uses its own language called HCL which stands for HashiCorp Configuration Language. </br>

Declarative IaC (e.g., Terraform, AWS CloudFormation) focuses on the "What" instead of the "How" </br> 

You write a file that says, "I want one virtual machine with this type, this operating system, and this network configuration." </br>
The IaC tool is then responsible for figuring out the exact sequence of API calls and commands needed to create that machine, handling things like dependencies, state management, idempotence, and updates. 

**State management**

This approach is more reliable because the tool can detect and correct for "configuration drift" and provides a single source of truth for your infrastructure's state. To achieve this, an IaC tool generally stores some state as files on storage somewhere. 

It stores the last updated state as the "known" state that was applied. </br>
When the code changes, it may detect the differences between the last "known" state and the "actual state" to ensure the actual infrastructure has not changed. </br>
This way, IaC tools not only detect the changes needed by the code, but also any changes that are in the actual infrastructure but are out of sync with the code. IaC tools will generally revert changes that have been made outside of the code. 
For example if the last "known" state is not the same as the "actual" state and not in-line with the desired state (code), the IaC will revert those changes to ensure the "known" state matches the "actual" state. </br>

The IaC will propose to make the desired state changes to known state and actual state. </br>

## Simple declarative automation using Vagrant

In this module, I will demonstrate a simple declarative approach to creating our infrastructure using a tool called Vagrant. </br>

Now the motivation behind the tool is all about choosing a tool that fits the role of the task at hand. </br>

As we are demonstrating declarative tools and local virtual servers, I chose Vagrant as its very popular and extremely simple for local provisioning of servers. The learning curve to demonstrate the course content is much lower than other Infrastructure as Code tools. </br>

Remember this course is all about learning the concepts in small chunks and applying these concepts with real tools without becoming attached to the tools themselves. If we were to provision cloud based VM's, I would likely show you Terraform or Pulumi. 

The important take-away here is that you learn the concept and how the tools achieve the desired outcomes as well as see the Pro's and Cons of these tools. </br>

### Install Vagrant

Let's start with the [Vagrant Documentation](https://developer.hashicorp.com/vagrant/install) </br>

We can fire up a terminal and test the vagrant cli by running `vagrant --help`

### SSH and Server Access

In this module, we will discuss the importance of `ssh` for accessing servers and importance of using `ssh` keys to access VMs. So far in this module we kept things really simple because our private server is not on a public network, therefore we've used username and password authentication with SSH. </br>
This is also done to keep the course simple and portable for mutliple operating systems. </br>

We can do an entire module on SSH, therefore I'd prefer to keep it simple in this module by sticking to what we have, but highlighting that we will be using and setting up SSH keys when working on public accessible servers and cloud environments. 

When we SSH to a server, the `ssh` client checks the host name we are connecting to and does what is called "strict checking". It maintains a list of known hosts and host keys that it has connected to in the past. </br>
When these match, it allows the `ssh` connection </br>
When things don't match, its a RED flag and the `ssh` client will refuse to connect. </br>
This is to prevent man in the middle attacks and ensures that we are connecting to known servers </br>

The **challenge** </br>
When working with local VMs which we will constantly create, destroy and re-create, the hosts and keys will constantly change, therefore we need to tell `ssh` not to worry about our `Host 127.0.0.1` as its local anyway and poses no threat. 

To do this, create a new file called `config` with no extension, in your Users directoty under `.ssh` folder. 

On **Linux**: `~\.ssh\config` </br>
On **Windows**: `C:\Users\<your-user>\.ssh\config`

Add the following content to the file: 

```
HostName 127.0.0.1
StrictHostKeyChecking no
UserKnownHostsFile /dev/null
```

### Initialise our Vagrant environment

The `vagrant init` command will initialise the current directory and generate a Vagrant file for a specific box. A "box" in Vagrant refers to a virtual server and in our case it will be a virtual box. </br>

```
vagrant init --help
```

This `init` command takes a `--box-version` so we would want to specify the version of the server image we want to use. </br> This means we no longer use the `.iso` image file we downloaded in Chapter 2. </br>

We'll also have to point to a `URL` which points to a ISO image in the [HashiCorp Vagrant Registry](https://portal.cloud.hashicorp.com/vagrant/discover)

I found a good set of Ubuntu Server images we could use if you search the registry for `cloud-image` you will find it. 

Firstly, we should open up the terminal and navigate to the Vagrant file in [this](./.test/Vagrantfile) location

**Important Note:**<i> Links in the video can change, so always click the link to follow the up-to-date path</i>

You can click the link above and either manually `cd` to the path or right click and choose  "Copy Relative Path" in VSCode. 

Once you `cd` to the directory, you should see the Vagrant file when running `ls` 

```
ls -l
total 4
-rwxrwxrwx 1 marcel marcel 1795 Aug 11 09:01 Vagrantfile
```

**Environment Variables**</br>
Before we run `vagrant init`, we want to be able to control where Vagrant will be creating all the virtual server files. There are environment variables that controls many settings and default behaviors. 

We can view them here in the  [Environment Variables Documentation](https://developer.hashicorp.com/vagrant/docs/other/environmental-variables)

Two variables that interest us:

* `VAGRANT_HOME` 
  *  Can be set to change the directory where Vagrant stores global state
* `VAGRANT_DOTFILE_PATH` 
  *  Can be set to change the directory where Vagrant stores VM-specific state, such as the VirtualBox VM UUID

**Important note:**<i>Please note that the commands you see in video footage may be slightly different from the document over here. This is because I may improve commands to ensure stability of the course.</i>

Let's set these variables! </br> 

Windows:
```
$ENV:VAGRANT_HOME = "C:\temp\vms\vagrant"
$ENV:VAGRANT_DOTFILE_PATH = "C:\temp\vms\vagrant\.vagrant"
```

Linux :
```
export VAGRANT_HOME="~/vagrant"
export VAGRANT_DOTFILE_PATH="~/vagrant/.vagrant"
```

Here is the full command I used to initialise our directory:
```
vagrant init cloud-image/ubuntu-24.04 --box-version 20250805.0.0
```

This will output
```
`Vagrantfile` already exists in this directory. Remove it before
running `vagrant init`.
```

This is because I already created a Vagrant file for us we can walk through. `vagrant init` will generate create a new file for you when starting from scratch. </br>

### Creating our Server

1. Let's review the [Vagrant file](./.test/Vagrantfile) contents

2. Configure our settings (Full variable list)

```
# windows
$ENV:VAGRANT_HOME = "C:\temp\vms\vagrant"
$ENV:VAGRANT_DOTFILE_PATH = "C:\temp\vms\vagrant\.vagrant"
$ENV:SERVER_SHAREDFOLDER = "C:\gitrepos"
$ENV:SERVER_USERNAME = "devopsguy"
$ENV:SERVER_PASSWORD = "devopsguy"
$ENV:SERVER_NAME = "my-website-02"

# ci settings
$ENV:PAT_TOKEN = ""
$ENV:GITHUB_ORG = "<your-github-username>"
$ENV:GITHUB_REPO = "my-website"
$ENV:RUNNER_NAME = "test-runner-01"

# linux
export VAGRANT_HOME="~/vagrant"
export VAGRANT_DOTFILE_PATH="~/vagrant/.vagrant"
export SERVER_SHAREDFOLDER=~/gitrepos
export SERVER_USERNAME="devopsguy"
export SERVER_PASSWORD="devopsguy"
export SERVER_NAME="my-website-02"

# ci settings
export PAT_TOKEN=""
export GITHUB_ORG="<your-github-username>"
export GITHUB_REPO="my-website"
export RUNNER_NAME="test-runner-01"

```

To run our server, we run

```
vagrant up --provision
```

Stop our server

```
vagrant halt
```

Destroy server: 

```
vagrant destroy
```




