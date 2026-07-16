# 🎬 Introduction to the Automation: A overview guide

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 5](../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience delivering software, CI/CD pipelines, automation, keeping servers up and running and up to date. </br>

Automation comes in many shapes and forms. The focus of this module is to conceptually understand automation in its various forms and tooling that can assist you and how it fundamentally solves what it is set out to do. </br>

### Too many Tools!

In DevOps, there are too many tools! </br>
The aim of this module is to provide clarity on the problems that are often required to be solved and how automation helps solves these challenges. To point out, there are a LOT of automation tools and they come with various pros and cons.

### Tools have Trade-offs 

For example, you can use **automation** to create virtual servers, patch their operating systems, deploy and provision websites and applications to them etc. You may use **CI/CD automation** tooling to deploy your website software. You may use **Infrastructure as Code** tools to automate creation of the servers from scratch. You may use **Configuration Management** tools to provision web server software on the virtual servers. 

So you can see that there are various specialized software within each domain. Some tools are great at **CI/CD**, some are great at **Infrastructure as Code**, others are better at **Configuration Management**. Sometimes you may need all 3, other times you pick 1 and choose the right trade-offs for your company and culture. You may opt-in for 1 tool to do it all. 
There is NO right and NO wrong answer here. 
But the choices you make will either make people's lives easier or more complicated and may cost little or a fortune. (Cost and complexity)

In this module we'll lay the groundwork of the different types of automation, so you can get an idea of where automation is used in different ways and which tools can help for each situation. </br>

### The Onion layers of Automation
In addition to the above, its important to understand that you can automate absolutely everything and anything by using scripts. But it does not always mean you should. You can use scripts to automate your website deployments. However, as we learned in the previous module, there are CI/CD tooling that help in many phases of website deployments. So you would not script everything. We could use Github Actions here instead to perform the "Trigger" on code push, "Checkout" the code and many of the phases. </br>

Scripts are like the core of the automation "onion" and when scripting has its challenge, a tool is developed to overcome that challenge and that tool forms a new layer on the onion. </br>

Similarly , we can use scripts to create our virtual server, and install our web server software (From Chapter 4) as well as configure our website. Or we could use a tool such as Terraform or Ansible. </br>

**Important Note:** </br>
You will not always be able to find a tool that fulfills every need and requirement and sometimes you need a small bash script to fill in the gaps. This is why scripting (bash for Linux and PowerShell for Windows) and coding (Python or Go) can be really important for DevOps engineers as it helps bridge the gap when tools falls short.  


## What is automation

Automation is the process we follow to automate tasks that would otherwise be manual. </br>
This speeds up tasks and reduces human error </br>

### Our automation journey

Let's have a think about everything we learned so far in setting up our Virtual Server in Chapter 2, and installing a Web Server in Chapter 4. 

We learned how to use Virtualization software to create and set up our Virtual Server. </br>
This involved **many manual "clicks"** and UI interactions to set up our Virtual Server. </br>
We set up our web server by running **many manual commands** in the terminal that was required to set up our web server in Linux. </br> 

**Key takeaway: These steps were all manual tasks.** </br>

If we had to ask another person to create a virtual server, even if the steps were documented, it's highly likely that the person would make a mistake with a setting, a click or an input. We'd then end up with two different servers. Slightly different disk sizes, different network settings etc. 

Manual commands are slightly better than using a UI, as there is slightly less room for input mistakes.

**Key takeaway: Cloud Virtual Servers have the SAME problems.** </br>

Automation helps **reduce these human errors** and **speeds up** the tasks of creating our virtual server and setting up our web server. 
I also tend to think that automation helps document a system, better than actual documentation. Documentation becomes out of sync, it requires additional maintenance, whereas a script for example tells us exactly how a system was created. A well written script can be a **living breathing document**. </br>

**Next Steps: (in upcoming modules)**

* We'll take a look at an automation script to provision our Virtual Server
* We'll take a look at a popular tool to make this even easier and solve some scripting challenges which we'll talk about.

## Key Focuses of Automation:

As previously discussed on the Preface, Automation comes in may shapes & forms depending on the challenge we are solving, the complexity involved, the cost and company culture. </br>
Let's start with the foundation. 

![diagram](./automation.drawio.svg)
### Scripting & Coding

As we've discussed up to now you should realise that all tasks can generally be automated by writing some scripts (using a scripting language) or code using a (programming language)
So instead of using the GUI interface to create our virtual server or run manual commands to set up our web server, we could write a script to do it all. 

#### Pros & Cons

The benefits of scripting is that we achieve the outcome of speeding up a task and reduce manual efforts and manual errors by humans. 

**1. Idempotence**

One big problem with scripting is that its not [idempotent](https://en.wikipedia.org/wiki/Idempotence)
This means that when we run a script multiple times, we are not guaranteed to get the same outcome each time. 

![diagram](./automation-idempotence.drawio.svg)

For example, if we have a script that creates a Virtual Server, and it fails at some random line, if we re-run the script, it may fail as the server may already exist. This means our outcome may be left in a broken state. 
The script needs to handle cases where the Virtual server already exists, in order to achieve idempotency. 

This goes beyond creating the Virtual server. </br>
When a script runs to install dependencies for our Web server software, a network call to the dependency packages may fail, and retry logic may be required. 

There would be many many places in the script where potential issues can arise, and we may never be guaranteed true idempotency. As we aim to achieve it, our script will grow and grow in complexity.

In an upcoming module we'll take a look at scripting to provision our Virtual Server & Web server and we'll experience this first hand. Then we'll take a look at specialised tools and how they help solve these challenges

**2. Imperative vs Declarative**

Another problem with scripting is understanding the differences Declarative and Imperative. </br>
Scripting is imperative </br> 

It means that scripts and code are a sequence of instructions that run from top to bottom. </br>
It's explicit instructions that keep changing the state of the system as it runs through control flows like `if/else` and `while` loops which all dictate the order how instructions are executed.
This order can change when a script is rerun. So you can see that firstly we have an idempotency problem and second is we don't always get the outcome we're after. We may or may not reach our desired state.

The declarative paradigm in DevOps is the capability to define "what we want" in a file instead of line by line sequential imperative execution steps (commands).
A Tool will then go and make that happen. It will validate & verify our declarative configuration file and review the actual state of the system and update it to match our desired state. 

### CI/CD: Automated build & deployments

We've covered the topic of CI/CD in the previous module and how automation concepts are used in building, packaging & deploying software. </br> We can achieve CI/CD by writing scripts as we've demonstrated at a high level already, or we can use tooling to assist in many of the CI/CD phases, so we do not re-invent the wheel and we reduce the number of scripts we have to maintain.

In our upcoming modules in this Chapter we'll go through a collection of Automation Lab modules where we will setup CI/CD pipelines which will involve **Tools** and **Scripting**, so you will get the idea of how to stitch together all the phases we learned about in the previous module. 

### Infrastructure as Code

If scripting is "code" and we write a script to build our virtual infrastructure, then is that not Infrastructure as Code ? </br>
We'll, Sort of... </br>
With a script, we will not know if there are any syntax issues until the script is executed. 
Scripts are not compiled and there is generally no build process. 

Many programming languages have compilers that ensure the code is valid. Code frameworks also provide test frameworks where code can be executed, tested and debugged. </br>

Infrastructure as Code takes advantage of this, where we define our virtual server as Code. </br>
Something like Python, Go, JavaScript, C#, etc. So the coding element its benefit.</br>
When using code, you can also write tests. </br>

#### Desired state

IaC also attempts to resolve the Idempotency issues. We don't run our code imperatively on our infrastructure as we would with a regular script. Instead with IaC, we run a tool that takes our code, validates it, and then compares our "desired" state defined in the code, with the "actual" state (our virtual server etc) and then it works out the differences. </br>

Desired state is a modern well-known paradigm in modern infrastructure and automation. </br>
Instead of creating infrastructure through a list of imperative commands, we describe "what we want" </br>
This description can be done through code (Python, Go, JavaScript etc) or it can be done through a declarative file. </br>

Infrastructure as Code generally uses code to achieve desired state. </br>
Orchestration tools (Kubernetes) generally use declarative files (JSON, or YAML) to achieve desired state. </br>

Infrastructure as Code tooling will show us what would change if we were to run the code. </br>
We can then accept it or reject it. </br>
The tool will then go ahead and update our infrastructure so the actual state matches the desired state. 

#### Pros & Cons

With IaC, we get the benefits of a programming language, and better idempotency as well as state management. </br>

There are some Cons. And many of these are my own personal opinions of working with IaC. </br>

* **User Introduced Complexities**
  - With Programming, you get the benefits of modular design, but also it allows humans to make things extremely complex.
  - Nested code, hard to read! 
  - Over-engineering.
* **Managing the state & drift**
  - Engineers tend to change the infrastructure (directly in UI's) without using the code.
    Emergency situations can cause panic where people are inclined to introduce manual changes. This means the actual infrastructure is now different than the code. We call this Infrastructure Drift.
  - Some tools manage state separately, and this can become complicated to manage when a lot of drift occurs

We'll have a separate Chapter that deep dives IaC in more detail with hands-on exercises. </br>

**Popular Configuration Management tools:**
* Terraform
* Pulumi
* Azure Bicep & ARM
* AWS CloudFormation & CDK

### Configuration Management

Where Infrastructure as Code focuses on the entire Infrastructure, Configuration management focuses a little more on what's running on the Infrastructure. </br>
In context with where we are, we can use IaC to spin up our Virtual Servers. Configuration management comes into play where we want to configure our Virtual Server. 

Configuring everything we need on our server, such as:
* `git`
* `ssh` access
* Installing our Web server
* Installing any CI/CD tooling
* Configuring our Web server
* Updating the operating system and ongoing updates

**Popular Configuration Management tools:**
* Ansible
* Puppet
* Chef
* Terraform

### GitOps

GitOps takes many of the Pros of the above and spins a slightly different take on managing infrastructure. </br>

Let's start by thinking of GitOps as Infrastructure as Code, but **where `git` is the source of truth**.
With IaC the source of truth is the actual state and as we noticed, there can be infrastructure drift when people make manual changes and not updating the code. This causes the actual state to be different from the desired state (our code). 
So therefore the source of truth is a bit questionable. 

![diagram](./automation-gitops.drawio.svg)

GitOps tries to solve this by having a tool that does **Automated Reconciliation** of the code and infrastructure.
If someone makes a manual change to the Infrastructure, the GitOps tools will technically revert that change with what is in `git`. </br> Remember, `git` is the source of truth when using a GitOps process. </br>
Some GitOps tools will allow manual changes to the infrastructure to be automatically sync'd back to `git`

If someone makes a change in `git`, the GitOps tool will detect the change, work out the difference between the actual state of the infrastructure and our desired state in the Code, and apply the changes. </br>
GitOps tooling generally runs 24/7 and checks `git` repo for changes and also monitors the actual state (infrastructure) for changes so it can reconcile any changes. 

GitOps has several advantages:

* Increased Reliability and Consistency: 
  - Because Git is the "source of truth" and reconciliation is automated, environments are highly consistent and reproducible.

* Faster and Safer Deployments/Rollbacks: 
  - Changes are made via Git (pull requests, code reviews), providing a full audit trail. If something goes wrong, you can quickly revert to a previous working state in Git, and the system will automatically roll back the environment

* Improved Security 
  - Direct access to production environments can be minimized or eliminated, as all changes go through Git

* Enhanced Collaboration
  - Developers and operations teams can collaborate on infrastructure and application definitions using familiar Git workflows (pull requests, code reviews)

* Auditability
  - Every change to the system's desired state is tracked in Git, providing a complete audit log
