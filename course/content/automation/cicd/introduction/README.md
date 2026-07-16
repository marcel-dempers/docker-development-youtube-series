# 🎬 Introduction to the CI/CD

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in DevOps in a  time before the concepts around CI/CD existed. To understand CI/CD, we need to understand the fundamental requirements of distributing software.

In our previous chapters we looked at GIT, source control, then servers, then we got a web server configured on our virtual server and configured a small website. 

In the real world, you would have many developers working on such a site or web service (microservice) or API. 

Technically - every time a developer makes a change in GIT, we'd like that change to reflect in our web server. 

**The process of delivering this software involves:**
* **Trigger**: Detecing the changes
* **Get Source**: Taking a copy of the code we want to deliver
* **Packaging** the software (Often requires a build process)
* **Testing** the copy of the software
* **Publish**: Copying this package to some storage for later deployment 
* **Deploy**: Copying this package to a target server
* **Configure**: Copying any configuration files for this application to the target server. This can also involves sensitive configuration called "secrets"
* **Live**: Make the new application live.   

Let's discuss the following diagram in more detail:
![diagram](./cicd.drawio.svg)

That's literally it! 

This was the world before CI/CD & DevOps concepts existed, and it remains part of the world therafter all the way to our current time. </br>

Dealing with the above steps can be a complex engineering endevour.

This is why CI/CD exists. 
CI/CD takes these steps and places them into two separate areas of concerns 


## What is Continuous Integration (CI)

CI, or Continuous Integration takes the first half of the steps which concerns integration of changes into source control

CI involves the start of the workflow, which involve detecting changes in our source control system, triggering the packaging (build) process, running any tests (optional), any security scans (optional) and pushing the packages to some storage for later deployment. 

CI involves the following steps:

* **Trigger**: Detecing the changes
* **Get Source**: Taking a copy of the code we want to deliver
* **Packaging** the software (Often requires a build process)
* **Testing** the copy of the software
* **Publish**: Copying this package to some storage for later deployment 

To understand the above  steps, let's break them down further and in relation to everything we've learned to far in Chapters 1 - 4. </br>

### CI: Triggers

For the CI and CD process to start, we need a trigger. In modern software engineering, this generally occurs when a `git` commit is pushed to the mainline branch. 

Let's discuss the following diagram in more detail:

![diagram](./cicd-trigger.drawio.svg)

In Chapter 1, we made file changes in VSCode (our IDE editor) and we used the `git` commandline tool to `git add -A` , `git commit -m "message"` and `git push origin main` to stage, commit and push changes to our Github repo. </br>
When this push occurs, we would like to trigger a CI pipeline or workflow. </br>

It's important to know that a trigger is simply an event and could technically be anything. Most common trigger is when source code changes happen.

#### Triggers: Writing a CI trigger

In a world before modern DevOps, we could write a trigger using a Bash script similar to Chapter 2. 

When can write a `while`loop, that `sleep` for 60 seconds and within each loop, it can go a `git pull` to pull the latest source code to a folder. We can store the deployed commit hash in a file and use a set of commands such as `if` statements to compare the current commit hash to the one that we deployed. 
Therefore we are able to "detect" change and can run another Bash script to trigger our CI steps. 

Given what we learned in Chapter 2 with Bash scripting, it sounds like it would not be too difficult. </br>

Imagine the following pseudo code for a potential bash script:

```shell
while true
  sleep 60s
  CURRENT_COMMIT=$(read from a file)

  echo "getting latest source code..."

  git pull origin main
  LATEST_COMMIT=$(git log -1 --pretty=format:"%H")

  if CURRENT_COMMIT != LATEST_COMMIT
    ./trigger.sh
    echo $LATEST_COMMIT > to a file
  else
    echo "no changes detected... waiting 60 seconds"
done
```

However there can be a lot more complexity with the following **painpoints**. </br>
* If our server reboots, we'd need a mechanism to re-run our script and keep it healthy. 
* We'd need a way to detect if the script has any problem or is NOT running
* If our script detects two or more changes it needs to handle parralel execution
* Need for a separate server to host these trigger scripts
* And there may be more complexities such as maintaining these scripts.

#### Triggers: CI Tools

Instead of writing our own, we could use a CI tool for change detection and triggers. 
CI tools have the capability to either poll git repositories to check for changes at an interval, or can be setup so the source control system (Like Github) sends an event to the CI tool. 
Many hosted source control solutions like Github, Gitlab &  Bitbucket all provide these as a service.
You can host your own by using CI tools like Jenkins 
There are many pros and cons to different CI tools with these features. 
I'll discuss a few regarding hosted and self hosted options in the course video.
These Pro's and Cons are mainly around:

* Cost
* Scalabilty
* Complexity
* Maintainability
* ** Biasness, Opinions & Experiences (<i>**to be discussed</i>)

### CI: Get Source

The "Get Source" phase is pretty simple and self explanatory. It's receiving the latest desired copy of the code
that we need to package. 
In our example above, this phase would be completed when we run our trigger, as we have to grab the latest commit in order to detect if the commit hash is different to the previously retreived commit hash. </br>

All CI tools that we will talk about, all generally check out the latest of the code in the first stage when a trigger occurs. </br>

It's essentially just a `git pull origin <branch-name>` 

However, CI tools usually perform a light weight `git` query to detect change during the trigger, then 
proceed with a shallow clone into a temporary folder where all the rest of the CI pipeline can run on. </br>

**So few things to discuss here:**
* CI Tools manage temporary checkout directories for `git` branches
* CI Tools do light weight git clones to save space and spend less time cloning large codebases
* CI Tools have to cleanup old directories to prevent the server from running out of space 


### CI: Package & Testing

The packaging step involves taking all the files we want to deploy and creating a package, also known as an artifact. </br>
In the case of our website files we've been using throughout all the chapters, this would be a simple `zip` command and we'll produce an artifact as a compressed `zip` file. 

A package is generally just a compressed format file, like `zip` or `tar` etc. 
There are different types of packages for different systems:

* `zip` and `tar` is just a compression format for files, and we can call that compressed result a package
* Programming Languages create packages too:
  * C# .NET creates `NuGet` packages and compiles into `.dll` files
  * NodeJS, JavaScript creates `npm` packages
  * Python creates `pip` packages
  * Go compiles into a `binary` file 
* Infrastructure packages:
  * Docker or Container images
  * Kubernetes manifests or `helm` charts
  * Azure ARM templates or Bicep template packages
  * AWS CDK packages

These are important as you will run into them when working with developers. 

**Important: To create packages and compile code, your CI server will often require dependencies.** </br>

### CI: Publish

The publish step involves taking the package or artifacts produced by the previous Package step, and sending that off to some storage. (Where it can later be deployed) </br>

There are a number of benefits in doing so:

* Decoupling CI from CD - The publish step does not deploy any software. It simply places it on the shelf for a deployment pipeline to pick it up
* Allows intermediate tools to run:
  * Security scanners can scan packages or artifacts for any malicious files. 
  * Artifacts can be signed - This helps deployment tools confirm the package was not tampered with.
* Not all packages or artifacts need to be deployed
* Staging \ UAT or Local environments can pull certain packages for testing and evaluation

Types of storage can also differ depending on technologies used, see Packaging step above. </br>

## What is Continuous Delivery \ Deployments (CD)

CD, or Continuous Deployment (or Delivery) takes the remaining half of the steps which concerns deployment of an artifact to a target environment.

CD can trigger automatically or based on another trigger which depends on the requirement. </br>
For example, we can automatically trigger a CD pipeline when we detect a package or artifact arriving at our storage </br>
Or the last step in the CI workflow can start the first step in the CD workflow. </br>

CD can also be manual when there are manual human approvals required for deployment. 
Or if you have a Quality assurance team that needs to sign off the packages. </br>

These intermediate phases between CI and CD and what that trigger looks like, is highly dependant on company culture and how mature 
teams are. </br>
We'll discuss this further in this module video </br>


CD involves the following steps:

* **Deploy**: Copying this package to a target server
* **Configure**: Copying any configuration files for this application to the target server. This can also involves sensitive configuration called "secrets"
* **Live**: Make the new application live.

### CD: Deploy

The main objective of the Deploy phase is literally just the act of copying files to a server. 
During the CI process we "prepped" the files into a package. <br>

The package is pulled by the CD tool.

Let's discuss the following diagram in more detail:

![diagram](./cicd-deploy.drawio.svg)

The CD tool then pushes the package to the target environment. </br>
This could be done via simple `ssh` copy with a new commandline tool called `scp`
The new package would usually be copied to a temporary location on the server. </br> 
We do this so we prevent impacts to the current website. </br>

We need to ensure our entire CI/CD process does not disturb the current environment, hence we will learn the concepts of "graceful" deployments throughout this Chapter. 

**Graceful Deployments**

This is the act of doing a deployment with 0 disruption to current HTTP traffic. 
* We copy our website files (unzip our package) to a temporary directory
* We create a new directory where the webserver will serve from
* The webserver will continue to serve requests from the current directory as configured
* Once we have the files copied to a new location, we update the webserver to look at the new directory. This is the **LIVE** phase we will discuss soon. Basically the act of making the new changes "live".
* The Webserver will automatically serve the new files for new connections, while existing connections use the current (old) directory. The webserver will stop serving the old content for any new connections. Depending on how slow the site is, it may take time to "drain" old connections this way. 

Let's discuss the following diagram in more detail:
![diagram](./cicd-deploy-2.drawio.svg)

### CD: Configure

As we've discussed **Graceful Deployments** above, we've jumped the gun a little. We've gone from Deploy to Live.
In more realistic production systems, there is an intermediate phase before going live and enabling HTTP traffic on the new webserver. </br>

Many websites, web APIs or microservices have configuration files. 
Most of these systems have other dependencies like databases or storage or other web services.
For example, our website in a DEV environment may connect to a database in the DEV environment.

This means that the DEV website, needs a connection configuration to tell it where the DEV database is. </br>
This also may require credentials like a user account name and password. These sensitive configuration options are called "Secrets" </br>
The PROD website also needs this in order to connect to the PROD database.
This means we need configuration files & Secrets for DEV and PROD websites. 

This whole concept is referred to as Configuration Management and Secret Management and there are all sorts of tooling available to manage these depending on business requirements, security standards and company culture. 

**Use the same package for each environment!**

**Deployment** = Package + Environment Configuration

Let's discuss the following diagram in more detail:
![diagram](./cicd-config.drawio.svg)

### CD: Live

The "Live" phase is a final phase of CD where we make our deployment changes and configuration changes live. </br>
This is the only phase that technically affects our customers. </br>
As we let our webserver know (using configuration files) to serve traffic from the new folder location

Let's discuss the following diagram in more detail:
![diagram](./cicd-live.drawio.svg)

## Popular CI/CD Tools

* Jenkins - open source, highly extensible, large plugin-ecosystem, free
* CircleCI - highly configurable CI/CD platform known for its speed, performance, and intuitive interface.
* Hosted CI/CD : Github Actions, GitLab, BitBucket - Built-in to hosted source control
* Cloud Based: Azure DevOps, AWS CodePipelines - Cloud Provider offering
