# 🎬 Scripting Example: Automate our Infrastructure with Bash

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in automation and scripting. </br>
The module aims not to teach the scripting fundamentals but to teach you my process I follow when developing an infrastructure script, the pros and cons, and various other tips. </br>

Scripting provides engineers with a lot of **power, but also has its challenges.**
As mentioned in our overview of this chapter, we can create a full CI/CD pipeline as well as provision infrastructure without **ANY** tools. </br>
This is powerful, but does not mean it's efficient. </br>
Scripts tend to grow in exponential levels of complexity. </br>
In this module, we learn to strike a balance. </br>

There are ***tools** available that help reduce the amounts of scripts we have to write
* Automation tools
* Configuration management tools
* Infrastructure as code tools 

It's important to understand that not all tools solve all complexity either. </br>
My point is that I still use scripting to this day to create cloud infrastructure when it makes sense to do so. </br>
For example, Terraform is a great IaC code especially when the infrastructure becomes more and more complex. </br>
However, in a cloud environment where I only need a small set of servers and very little cloud resources, I would prefer to resort to a script. </br>

The point here is that it's very much dependent on the situation and how much infrastructure you need and how much it will grow by.

Understanding **when to stop scripting and avoiding complexity** </br>
Us engineers tend to get excited when we are solving problems and we may not always realise that we are introducing complexities somewhere. We become so fixated on solving a problem, and at the same time, we may introduce a large script that requires a lot of cognitive load to maintain and understand. 
There are times when I would prefer to write a short document and have a process be fully manual, instead of introducing a very complex script with dependencies and complexities that an engineering team has to nurture in an ongoing manner.
 
We will also focus on the process I take when designing and building a script. Taking everything we have learned from the command line modules, recording those commands and coming up with a structure for our script. Then iterate on it. A Script is almost never complete. It will forever go through an iteration of building and improvement. </br> 

## Scripting our Server Infrastructure

We still take the steps to create a script that allows us to re-create our virtual server from this course. </br>

This makes our server more **immutable**. 

**Immutable infrastructure** is a term used mostly in cloud computing where servers are never modified after being deployed. </br>
Instead of updating a server in place with patches or new software, a brand new server is created. The old server is then completely replaced by this new one. This method eliminates the risk of **configuration drift** and inconsistencies, making deployments more reliable & predictable. 

### Requirements

In order to create virtual servers, we need the CLI tool for it. When we installed VirtualBox, it also installs a command line tool called `vboxmanage`

To test out the command, we can run `vboxmanage --help` </br>
We can also navigate the command tool using `man vboxmanage` </br>

To start scripting, let's `cd` to the [location](.test/create-vm.sh) of this script. Remember the location can change, so use the link to see the location and use `cd` to follow it in the terminal. 

To run this script we need a bash terminal. </br>
In linux, you can use the default terminal </br>
In Windows, gitbash will do.

<i>**Note**: For most if not all of this course, we will run commands in our Virtual Server. This means the course and what runs on your machine will produce consistent outcomes. However this is a rare case where we will run outside of the Virtual Server because we need to provision a Virtual servers.</i>

## Using the Command line

Throughout Chapters 1 up until now, we've been learning about the importance of command-line. The cool thing about command-line is that when you run commands in the terminal to create infrastructure, you have an opportunity to record these commands in a file. </br>

This file then becomes our script. </br>
Now you have a playbook of how to automatically do what you just did in a quicker way. 

The biggest benefit for me is that it makes my infrastructure completely immutable, meaning I can throw my server away and instantly recreate it with my script. If I spend hours creating my VM using the UI, and then using the UI to configure it and set up applications, when something goes wrong, I have to fix it or it will take hours to recreate it. 

Recording the commands also help if two different people need to create a server, they will create the **SAME server**

So command-line and scripting have great synergy. 

## Process of building a script: structure

When I start scripting, I firstly think about "what is the outcome I am trying to achieve with this script?"

In this module, the outcome is "Create a virtual server". </br>

To do this, we know our server will need a name. Our virtual network will need a name.
We can define some important parameters at the top of the script. </br>

**Parameters** </br>

Some variables may be "settings" that we can tune over time. </br>
Something like the memory and CPU sizes:

```
VM_MEMORY="4096"
VM_CPUS="4"
VM_DISK_SIZE_MB="25600"
```

Other variables, like "VM_NAME" could be an input parameter. </br>

```
VM_NAME=$1
```

This will allow us to run a script and create a few servers and make our script **reusable**: 

```
./create-server "my-website-01"
./create-server "my-website-02"
./create-server "my-website-03"
```

**Idempotent Structure** </br>

We discussed the importance of being able to ensure our script is rerunnable. </br>
There are some obvious structures that our script will need:

* if the server does not exists
  * Create it
* else
  * Update it

We can also apply this flow to things like the virtual network which may or may not exist. </br>

```
# network 

if network_exists "$NETWORK_NAME"; then
    # update network if needed
    echo "NAT network '$NETWORK_NAME' already exists."
else 
    # create network
    echo "NAT network '$NETWORK_NAME' not found. creating..."
fi

# server 

if vm_exists "$VM_NAME"; then
    # update vm
else
    # create vm
fi
```

**Functions** </br>

Our `if` and `else` structures are very easy to read and to achieve that, we had to define two functions. `network_exists` and `vm_exists` </br>

Let's define our functions:

```
network_exists() {
    VBoxManage natnetwork list | grep -q "Name: *$1"
}

# Function to check if a VirtualBox VM exists
vm_exists() {
    VBoxManage list vms | grep -q "\"$1\""
}

```

The cool thing about functions is we can paste them into the terminal and trigger them which helps us test these functions to make sure they work in isolation without having to rerun our entire script. </br>

## Process of building a script: iteration

Now we can go ahead and run our script because we have an idempotent structure with inputs and our script is now ready to run in the terminal. </br>

```
./create-server.sh "my-website-02"
```
Now we have something we can iterate on. We can now add more and more functionality and test it. The quicker we define our above structure, the quicker we can start running our script. </br>

This is important because we want more iterations. The more we run our script from top to bottom, over and over, the more our script becomes "tested". </br>

The more we run commands manually against our infrastructure, the more "out-of-sync" our script becomes. We want to delete our infrastructure and recreate it using our script as often as possible. </br> This is called iteration. </br>
It's the equivalent of "reps" when you are trying to build muscle in the gym. </br>

Let's go ahead and add functionality. </br>

<i>**Important Note**: In the video, I will walk through the remainder of this script. </br>
At the time you are viewing this video, the script may have slight changes as it may become more improved over time. 
</i>

## The final test

The final test is to run our script with input parameters and end up with the virtual server in VirtualBox 

Please be aware, sometimes VirtualBox needs to be restarted to see the new server or network changes in the app 

```
./create-server.sh "my-website-02"
```

## Key Takeaways and Challenges

* **Growing Complexity** </br>
  The more we add to our script, the more complex it becomes. Adding more is not always good. Sometimes knowing when to add or when to remove code is better than just focussing on solving a problem. Not all problems need to be solved.
* **Structure & Iterate** </br>
  The importance of idempotence and structure in a script so it performs a clearly defined purpose. Then start executing the script, add functionality while you run it as often as possible to keep testing by iterating. 