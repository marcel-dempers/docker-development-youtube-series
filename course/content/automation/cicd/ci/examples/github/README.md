# 🎬 CI Examples: Creating a CI Pipeline with Github Actions

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in DevOps using a variety of CI/CD tooling over the years and observing how the landscape as well as the tooling has evolved. </br>

In the first module on CI/CD, we looked at the software delivery landscape as a whole. </br>
Now we are zooming in on the fist part: CI </br>

At a very high level, CI is all about taking code, and placing it on the shelf ready to be deployed. </br> CI tools fascilitate all this. </br>

As we go through this module, I want you to pay attention to two things: 

* Observe a CI process detached from the technology
  * This will help you see a clear picture of the challenges that software delivery presents without getting attached to the technology used. It will also allow you to avoid biasness when it comes to tools. We are engineers after all. Tools will excite engineers.
  It is important to see a tool for what it is and more importantly, **what it is not**.
* Observe how a CI tool works at solving CI challenges 
  *  This will allow you to see what challenges the CI tool is good at solving and **where the gaps are**. We need to remain critical of software tools because every tool comes with a  **trade-off**

## Continuous Integration overview 

We covered a strong foundation of knowledge in our module on our CI/CD overview. </br> 
Without repeating anything let's quickly take a glance at the overview diagram and see where we will zoom in today. 

![diagram](../../../introduction/cicd.drawio.svg)

When examining the CI steps that we covered in a previous module, we have various important steps like "what triggers the pipeline to start", "getting the code", "packaging the code" and "publishing the code". </br>

To perform these steps we would need a CI server </br>
a CI server is a server (just how a Web server is a server), that performs these abovementioned steps. </br>

## Networking Challenges & Security Concerns

There are a few important points to address when it comes to choosing a CI server that works for you. </br>
The point of this module and course is to inform you about these choices, 
so you pick the right tool for the job and also end up with architecture that makes sense. </br>

There are some **networking** challenges you'll face when setting up CI & CD tools, 
this is because these tools may need access to source control, artifact storage as well as your server environments. </br>

With the access mentioned above, it means there are also **security concerns**.

* Where do you place the CI & CD tools, in which network?
* What do you expose by letting your CI & CD tools access your resources?
* What type of access do these CI & CD tools need in order to achieve their goals? 

### Security: Access and Authentication

For example, servers that help with CI/CD, sometimes come with a UI. This makes it easy for teams to configure CI/CD pipelines and set everything up by using the mouse and keyboard. </br>

Firstly this UI, needs to be accessible by the teams. </br>
We do not want our UI to be publicly accessible even if it requires a password to get in. </br>
Attackers may use brute force to try get their way in, and CI servers are the pot of gold for an attacker. </br>
Keep in mind that CI servers have access to company source code, 
and also CD servers may have access to production infrastructure if its able to deploy code. </br>

Therefore when it comes to CI/CD we cannot take security lightly and need to be very vigilant. </br>

The below diagram highlights what I mean by access and the consequences your choices will have. </br>
The networking and security goes hand in hand. Let's discuss the following:

![diagram](./ci-ci-security.drawio.svg)

The above example demonstrates the tradeoffs that network considerations have on the security. </br>

* Option A:  We dont expose `ssh` publicly, however we need to allow access to the CI server either publicly or perhaps a VPN of some sort
* Option B:  Demonstrates a CI server sitting on its own network and demonstrates the need to open `ssh` port `22` publicly. </br>

Later in the module I will demonstrate the path and tradeoffs we will take. </br>

### Security: Unprivileged Access 

In our Chapter on Servers & Linux, we talked about users. We talked about the fact that users should have least priviledged access to be able to perform only the required task at hand. </br>

#### Security: Source control accounts 

This same logic applies to CI/CD especially CI tools. </br>
We want to ensure that the accounts that are used to run the CI tools only have access to perform the tasks it needs to. </br>
For example, a CI server technically only needs `Read` access to source code in order to detect a code change and perform a "Trigger" phase. </br>
Similarly, the "Get Code" phase would also need `Read` access in order to read a copy of the code and start the build process. </br>

In Chapter 1 we talked about source control. When using something like Github for source control, we may need to consult GitHub's documentation to ensure we only expose our source code with `Read` access for the CI tool we plan to use. </br>
The technical implementation may be different between source control providers. </br>
All the challenges mentioned in this guide remains the same </br>

Therefore we need to take some careful considerations for the source control account that the CI server uses. </br>

#### Security: Linux server accounts

Similarly the CI tool installed on a server may also need to run as an unprivileged user. Keep in mind that CI tools take code, and execute it, so its the perfect attack surface for a malicious actor to run malicious code. Therefore we may create a new unprivileged Linux user when installing CI server tools. 

### Network Considerations 

Let's take a look at our diagram on CI/CD and pay close attention to CI phases. </br>
We only need access to source code and to the location we are going to publish artifacts to. </br>

This is where location and network considerations come in as demontrated earlier. </br>
The following diagram shows the components that require access to one another. </br>
This allows us to discuss where we would like to place the CI and CD tooling:

![diagram](./ci-ci-network.drawio.svg)

* **CI/CD tool server location**
  * This plays a bit part in your network and security architecture.
  * CI/CD tool generally comes as one tool. 
  * We can either run it in the same network as the web server, or in its own.
  * We highlighted a few times that CI/CD needs certain network access to deploy to our web server.
* **Artifact location:**
  * We could secure the artifact location by running it in the same private network.
  * We could securet the artifact storage has a protection layer on top, like IP address whitelisting using a firewall, authentication or a combination of these.
  * CI tool needs `Write` access to push artifacts.
  * CD tool needs `Read` access to download artifacts for deployment.

### Pull vs Push Considerations

Many of the above challenges can be resolved by using either a "pull" or a "push" model. </br>
In the diagrams so far, we observed a "push" model. This is when a deployment server pushes a change out to a production environment. </br
In a "push" model, we can see the tradeoff is that a deployment server needs access to a production server. </br>

In some companies, this is not an option as networks are restricted and CI/CD should not run in Production networks. </br>

This is where a "pull" model comes in. </br>
A "pull" model is where a Production web server has read access to the artifact storage and pulls packages into the private network. </br>
This means deployment server only needs to push artifacts to storage and does not need access to Production servers. </br>
It also means our Production server does not need to expose `ssh` on port `22`. </br>

![diagram](./ci-ci-pull.drawio.svg)

In our lab, we will follow a "pull" model, because our Virtual server is not publicly accessible via `ssh` (only privately from our computer)

**Important Note:** <i> One could make an entire chapter on networking and security considerations for CI/CD tooling, so I would keep notes of all of these and pay attention during our hands-on exercise so you can see how we address some of these challenges. </br>
These will reflect in real production networks too. </br>

If you are fortunate to work for a company that have CI/CD tooling in place, I would highly recommend to ask questions to the people that maintain these systems and use that as another means to learn from a production-ready platform. 

Remember, no question is a bad one, so ask away! </i> </br>

## Choosing & Installing a CI tool

CI/CD tools generally follow a master \ worker architecture. </br> 
This means you have a master or main server that hosts the user interface (website to access the server), storage, the user accounts , CI/CD pipeline definitions etc. </br>

Popular CI/CD tools like Jenkins and TeamCity follow this architecture. </br>

![diagram](./ci-ci-arc.drawio.svg)

Roles of the main server:
* Hosts the UI, This is typically a web server running on a port like `80`, `443`, `8080`, `8000` etc. 
* Stores the data. This data is user accounts, workflow definitions and anything that is "saved".
  * This typically uses a Database in some cases, like SQLLite, SQL, MySQL, Postgres.
* Performs the "Trigger" actions which starts pipeline workflows.
* Place workflow tasks onto a queue for worker servers to perform.

Roles of the worker servers:
* Run the CI/CD pipeline workflow tasks that are queued by the main server. 
* Workers provide scalability so we can run tasks in parallel and support many without queueing them up.
**Note:** 
<i>If you only have 1 server you can run the main and the worker on that same server. </br>
Some CI/CD tools allow you to run make the main server as worker as well. </br>
</i>

### Choices: Hosted vs Self-hosted 

This is where a lot of tradeoffs come in. Having to host a "main" server, you need to take into account all of the security concerns mentioned earlier. </br>
You need to protect it from being publicly exposed, protect it from HTTP attacks, login & brute force attacks etc. </br>
Its also an extra server that you need to maintain, upgrade and patch. </br>

**Hosted option**

Sometimes it's far easier to use a hosted option. </br>
For example, a provider can run the "main" server for you and all you need to do is login and use it. </br>
Github, Gitlab and Bitbuck (Source control providers), AWS CodePipelines and Azure Devops (Cloud providers) take this one step further and they have CI/CD capabilities integrated with their platforms. </br>

To keep this course as simple as possible, and as we are already using Github for our repository, we can rely on Github to host, so we can use Github Actions.
This means that Github runs the "main" server that performs the "Trigger" phases. </br>

**Self-hosted option**

One thing that we will do slightly different is that we will run our own workers to perform the CI and CD pipeline. </br>
This takes into account the security challenges because our servers are private and in a private network. </br>

If we wanted to go full hosted, where Github runs both the main server and the workers, we would have to make some tradeoffs:
* We would have to make `ssh` public and forward that port `22` on our router through to our virtual server.
* That would open port `22` to the public.
* It would also make the port-forwarding more complicated to maintain for this course.
* Everyone following this course have different internet service providers, different internet routers, so having consistency around that becomes difficult. </br>
* You IP address provided by your ISP, may change. This means Github would not be able to contact your server for deployment.

Our architecture:

![diagram](./ci-our-arc.drawio.svg)

The way the Github workers (called "runners" in Github terms) work is they start up and register themselves with Github. This way Github's "main" server
can talk to the workers and orchestrate workflows. </br>

### Installation 

#### Navigating Documentation

This is a little side note for this course. As an engineer you need to start building skills in navigating documentation. </br>
This will set the foundation of how well you implement solutions and form understanding of topics. </br>

**Important Note:**
Remember that websites change and links may break so whatever you see on this video now may be different in the future. </br>
But the workflow we will create should still be the same and I will keep it up to date. </br>

Lets start with the [Github Actions Website](https://github.com/features/actions)

#### Setting up a CI workflow

In this module we'll setup a basic workflow and we'll start to see the relevant CI tasks as well as why our architecture is so important. </br>

My example workflow file we will be using can be found [here](./.test/cicd-workflow.yaml) if you need to reference the same file. </br>

**Steps:** 

1. **Prepare our Repo** To define pipelines, we will be opening our `my-website` GIT repository that we created in Chapter 1. </br>

2. **Review a new Github Action** We will navigate to the `Actions` page on our repository on Github and explore the interface and what it allows us to do. We learn that Github looks for workflow definition files under `.github/workflows` in the repo. 

3. **Prepare our CI/CD Tool** 
    *  Workflows under the `.github/workflows` directory will only run once it's in the `main` branch. 
    *  However, it will not run because we dont have our own CI runner defined yet. 
    *  Let's do that first.

4. **Setup our CI/CD Tool**:  Github Actions Runner.
    * When we follow the navigation -> Github Repo | Settings | Actions | Runners
    * There is a button for `New self-hosted runner`
    * We can select `Linux` as the Runner image
    * The architecture is `x86`.

5. Copy the content of the Linux bash commands provided by the Github Runner Installation page from the `Download` and `Configure` sections 

6. We'll review these commands and notice its limitations in terms of idempotence and scalabbility.
    * One downside is that the TOKEN provided by the Github install page, expires as its short lived. We can use the Github API to get a TOKEN on demand. We'll look at the [GitHub API Documentation](https://docs.github.com/en/rest/actions/self-hosted-runners?apiVersion=2022-11-28#create-a-registration-token-for-a-repository)

7. Let's walk through fixing this and automating the Github Runner install using a bash script which you can find a final complete copy of [here](./.test/install-github-runner.sh)


8. Let's setup the following folder structure in our `my-website` repo with everything we have so far: 

```yaml
/
├── infra/            # or 'infrastructure'
│   └── scripts/      # Bash scripts for automation
│
├── src/              # Website source code
│
├── .github/worflows  # GitHub Actions workflows
│
└── README.md         # Documentation
```

**Make sure we have a terminal open at our my-website GIT repo**:

```shell
mkdir -p infra/scripts/
mkdir -p src/
mkdir -p .github/worflows
```

9. **Copy Infrastructure Scripts**: We can keep a copy of our new [install-github-runner.sh](.test/install-github-runner.sh) in the newly created `infra/scripts` folder.
    *  This is so that we keep a copy of it for futher automation later

10. **Copy Website Source Code**: We can copy our Website Source Code into the newly created `src` folder:
    * [home.html](.test/home.html)
    * [script.js](.test/script.js)
    * [style.css](.test/style.css)
    * I like to keep source code separate from other files

11. **Copy CI/CD Pipeline**: We can now take the file from this course [cicd-workflow.yaml](./.test/cicd-workflow.yaml) and copy it to our `my-website` git repo under `.github/workflows`

12. **Review our Source Code**
    *  Our my-website repo structure should look like the following 

```yaml
/
├── infra/           
│   └── scripts/install-github-runner.sh     
│
├── src/
│   └── home.html
│   └── script.js
│   └── style.css             
│
├── .github/worflows/cicd-workflow.yaml 
│
└── README.md        
```

    * If we run `git status` in the terminal, we should see our files ready to be committed.

13. **Finalise the CI/CD Tool**
    * Lets `ssh` into our Virtual Server and run our script in the `/infra` folder
    * #TODO

```shell
PAT_TOKEN=''
GITHUB_ORG=''
GITHUB_REPO=''
RUNNER_NAME=''
./infra/scripts install-github-runner.sh $PAT_TOKEN $GITHUB_ORG  $GITHUB_REPO $RUNNER_NAME
```

14. **Review our CI/CD Tool installation**
    * We should see our Github Action Runner under the `Actions` page in our repository settings.

15. We can now finally commit our script, as well as our new Github workflow file to our git repo and watch the pipeline run!

```shell
git status
git commit -m "add pipeline files and refactor our structure"
git push origin main
```

# Key Takeaways 

* **CI Phases: Building & Packaging**
* **Architecture**
* **Network Considerations and Security**
  * Servers, Processes connecting to each other - Network Architecture
  * Security & Access required by the processes and servers as well as source control access
* **Reference**
  Creating your first Github Action https://docs.github.com/en/actions/get-started/quickstart 

# Source Code Reference

Our Website full source code can be found on Github. </br>
Keep in mind that Git repositories always show the latest code in `main` branch. </br>

As this course evolves, every chapter is an iteration on top of the previous. To get the source code "as it was" for this Chapter, you can look at the following tag: </br>
https://github.com/marcel-dempers/my-website/tree/chapter-05



