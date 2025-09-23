# 🎬 CD Examples: Creating a CD Pipeline with Github Actions

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in DevOps using a variety of CI/CD tooling over the years and observing how the landscape as well as the tooling has evolved. </br>

In the first module on CI/CD, we looked at the software delivery landscape as a whole. </br>
We went pretty in-depth with CI in the previous module. </br>

This allowed us to setup the foundation such as the architecture, security, network layout as well as installing our CI/CD tool that we will be using (Github Action Runner) </br>

We also have the pipeline workflow file ready to go, so all we really need to do is add the deployment steps </br>
With the foundation set, we can now zoom in on the second part: CD </br>

At a very high level, CI was all about taking code, and placing it on the shelf ready to be deployed. </br>
Now CD is all about taking the package from the shelf and deploying it. </br>
In our case this is going to be really easy. 

Therefore, as we go through this module, I want you to pay attention to two things: 

* Observe a CD process detached from the technology
  * This will help you see a clear picture of the challenges that software delivery presents without getting attached to the technology used. It will also allow you to avoid biasness when it comes to tools. We are engineers after all. Tools will excite engineers.
  It is important to see a tool for what it is and more importantly, **what it is not**.
* Observe some of the challenges of a deployment pipeline
  *  When we deploy files to a running web server, we can impact the current web requests that may be serving the old version. This is where rolling deployments come in. 
  
## Continuous Delivery Overview 

In the previous modules we covered the basics, and overview and an in-depth first look at Continous integration. Let's take a moment to recap the CI/CD phases so we get a sense of where we are at:

![diagram](../../../introduction/cicd.drawio.svg)

## The Deployment Phase

In our CI/CD overview module, we talked briefly about that a Deployment means. </br>

Conceptually its simply the act of moving files from an artifact storage, to the live environment where the files are required for deployment. </br>

The phrase "Deployment" often refers to the entire task of deploying new software which can include pulling the artifact to the location of the server, configuring it, injecting sensitive configuration (like secrets), and then making it "live". 

![diagram](../../../introduction/cicd-deploy.drawio.svg)

Now I've been in the industry long enough to know its not this simple. </br>
In my course, I will define the act of deployment as a small part of the entire task. </br>
It's simply to "copy" of files to a target location </br>

The diagram above, illustrates concepts, but does not mean it fits our architecture. </br>
There are certain challenges, some we observed before:

* Can the servers access the artifact storage to get the files for deployment?
  *  Are we using a "pull" vs "push" model for deployments?
* Are there any Network or Security restrictions that we need to consider?
* Will our actions of "Deployment" disrupt current network request or server operations ?
* Understand **Graceful deployments**

## Architecture Review

Let's review our architecture to get a vision of what the Deployment phase of CD would look like:

![diagram](../../../ci/examples/github/ci-our-arc.drawio.svg)

**Pull vs Push** </br>
To secure our architecture, we are using a "pull" model so our CI/CD tool runs on the same network as our server. </br>
We're also using a "pull" model because it greatly simplifies a course setup. </br>
For a "push" model to work, Github would need access to `ssh` to our Virtual server te perform the deplyment phase. </br>

**Multiple Servers VS Single Server** </br>

Now for simplification in this course, we have the CI/CD tool (The Github Runner) installed on the SAME Virtual Server as the Web server. This is done purely for simplicity as it would be a large overhead for learners to run multiple Virtual Machines on their laptop. 

The diagram illustrates a more realistic environment where you would physically seperate these. You would have the CI/CD tool on a different server than your Production Web server. This is for security reasons we discussed and also allows these servers to be decoupled and be individually scalable </br>

**Perform Graceful Deployment** </br>

In order to perform a graceful deployment without disturbing any existing traffic to our web server, we will deploy the code into a new directory that the web server is not currently using. 

After we deploy (copy) our files to server, we would tell our web server about the new files in the "Configuration" phase. 

This concept can also be referred to as a **Rolling Deployment** </br>
In our case a rolling deployment is us rolling a new deployment as a new directory and updating the web server to use it. That way the old deployment (directory) is still in place and we can roll back. 

**Rolling Deployments** can also mean to deploy entirely new servers on each Deployment. So instead of a deployment being "new files" on a fixed amount of servers, it can be new Virtual Servers entirely. 

![diagram](./cd.drawio.svg)

## Setting up a CD workflow phases

In this module we'll continue our workflow and we'll start to see the relevant CI and CD tasks come together. </br>
My example workflow file we will be using can be found [here](./.test/cicd-workflow.yaml) if you need to reference the same file. </br>

**Steps:** 

1. **Overview** Let us review the following block of Github Actions

```
- name: deploy package
  run: | 
    . infra/scripts/deploy.sh $PACKAGE_NUMBER

- name: configure
  run: | 
    . infra/scripts/configure.sh $PACKAGE_NUMBER

- name: make live
  run: | 
    . infra/scripts/live.sh $PACKAGE_NUMBER
```

2. **Copy our script** We can use VSCode to copy the following scripts into the `infra/scripts` folder in our `my-website` GIT repository. Let's do that:

* [deploy.sh](.test/deploy.sh)
* [configure.sh](.test/configure.sh)
* [live.sh](.test/live.sh)

3. Commit the chages and push the changes to Github and watch our full CI/CD pipeline run

```
git status
git commit -m "add continous delivery to our pipeline"
git push origin main
```

# Key Takeaways 

* **CD Phases: Deploy, Configure & Pubishing**
* **Architecture**
* **Graceful Deployments**
* **Rolling Deployments**
* **Network Considerations and Security**

# Source Code Reference

Our Website full source code can be found on Github. </br>
Keep in mind that Git repositories always show the latest code in `main` branch. </br>

As this course evolves, every chapter is an iteration on top of the previous. To get the source code "as it was" for this Chapter, you can look at the following tag: </br>
https://github.com/marcel-dempers/my-website/tree/chapter-05

