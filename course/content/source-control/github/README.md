# Introduction to Github: Hosted Source Control

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 1](../../../chapters/chapter-1-source-control-git/README.md)

<b> 📢 Important note: I would highly recommend to checkout module 1 and 2 for an introduction to GIT and VSCode before starting this module </b>

## What is Github

Think of Github as a "GIT" in the cloud. </br>
A remote git server that runs on a server managed by a provider </br>
Just like you may have the ability to edit documents directly in the cloud instead of running the applications directly on your computer. </br>

Github is an online hosted source control service </br>

So we "stage" and "commit" files to our local GIT repository, we have been using commands like `git add` and `git commit`, we will now look at an additional command like `git push` and concepts of `origin` and `upstreams` and the origin is the GIT server we are pushing to. </br>

Benefits:

* You project is not hosted on your machine which could fail
* Multiple people working on the same projects
  * You can work from multiple devices
* Security features
  * access
  * auditing
* Review process for changes

## Create Github Account

Let's head over to [github.com](https://github.com/) and Sign up. </br>

## Create a Personal Access token

### Important Notes about Authentication 

When we authenticate with GIT for Windows, GIT may popup a browser window that asks you to sign in to Github. </br>
When you sign in this way, Github sets up authentication tokens for your account and GIT for Windows uses the Windows Credential manager
to store the credential. </br>

When you authenticate with GIT on Linux, you will get a prompt for username and password. </br>
Important to know that the password here is not your Github password, but rather a personal access token. </br>

Knowing how to create and manage personal access tokens (or PAT) is important because:
* These tokens have a shorter lifespan than your password making them more secure.
* They may have limited access, for example we may only use them to read repositories, or write, or both.
* We may use these tokens to login via scripts and automation pipelines

### GIT Credential Store

Another important note is that when we sign in using GIT for Windows, Windows may cache the credential so we don't have to login every time we access Github. </br>

For Linux, this is not done by default, and we may need to run a command once to set up credential caching. </br>

We can either cache our token for default 15 minutes, by running the below command:

```
git config --global credential.helper cache

```
Or we can specify a time in seconds if we wanted a shorter or longer time than 15 minutes

```
git config --global credential.helper 'cache --timeout=3600'

```

Or we can store it in the credential store for Linux and store it forever

```
git config --global credential.helper store
```

You can choose the appropriate command above for you so you don't have to keep signing in to GitHub when using GIT. 

## Create our first repository

On the [github.com](https://github.com/), top menu, you will find a `+` button where you can create a new repository </br>
Create a new repository called `my-website` </br>
Be sure to mark it as `Private` so only you can view it. </br>
Once our repository is created, we can navigate it in the Github UI, and we can also make changes using the browser. </br>

## Using GIT with Github

On the repository page, you will find a `<> Code` button, where we can grab the URL of the repo </br>
Note that this URL has a `.git` extension </br>

Now back to the command line, but this time instead of using a command line terminal, we can use VSCode, as we have done in the previous module. </br> 
Let's fire up VSCode, enter the terminal and use the command line to make sure we are in our GIT folder </br>


Then we run `git clone <url>` 
In my case its `https://github.com/marcel-dempers/my-website.git`

## Git configuration

In our first module on GIT, we learnt about the `git config` command and have taken a look at the `.git` hidden folder that is part of every repository. </br>
Inside this `.git` folder, is a `config` file which contains configuration of the repository. </br>
You have a global config that covers your entire account, as well as the config files per repositories </br>

In this module, we'll configure this repository to use our new Github account specifically with the `git config` command and the `--local` flag

```
git config --local user.name "your-username"
git config --local user.email "your-email"
```

## Making changes

### commits 
I want to show you that you can perform changes using the Github UI. </br>
For example, we have a readme file we can change and commit directly using the browser </br>

### history

We can use the browser to see history of our repository, who made the changes as well as what changes were made </br>

### using GIT and Github

Although we can make direct changes in the UI, most of the time its better to make changes on your local machine </br>
In our VSCode, let's go ahead and make some changes and follow the GIT process as we have done before, but this time we will push the changes up to Github. </br>

* `git add` to stage changes
* `git commit` to commit our staged change
* `git push origin master` to push our changes up to the origin

<i> note:  In Git, "origin" is a shorthand name for the remote repository that a project was originally cloned from. </i>

Notice we get an error. This is because our local repository main branch is out of sync, because we do not have the latest changes we made to the readme file </br>

* `git pull` to get the latest changes

Now we should be able to push again. </br>
And in Github, we can see the change history 

### branching 

Now in GIT, we covered branching, and we created a branch in the first module, using the `git checkout -b` command. </br>
Let's do that again and revise our newly learnt skills

```
git checkout -b github-skills
```

We can make our changes and follow the `git add` , `git commit` & `git push origin github-skills` to push these changes

### Pull requests

Now in Github, we can see the new branch, we can check the history on that branch and we can raise what's called a `Pull request` to propose to merge these changes into our main branch </br>

Pull requests gives us the opportunity to review the changes </br>
We can raise a pull request using the Github UI </br>
The Github UI provides a way for people to review, leave feedback, maybe request additional changes, and to approve the pull request </br>
Once merged, the feature branch can be deleted. </br>

Now back to our terminal we can switch back to our main branch using `git checkout master`.
And pull the latest changes `git pull origin master`

<i>By now, you should start to get comfortable with the GIT commands, and I would encourage you to build muscle memory with these commands and try to make additional changes to your website code. You can also use the README file as practise</i>

<i>Learn to repeat:
* `git status`
* `git add`
* `git commit`
* `git push origin master`
* `git pull origin master` 
</i>

### Issues

In Github, we can track tasks or work items by creating "Issues" </br>
The issues page can be used for various types of work items. They could be actual issues like problems with software, bugs etc. Or improvements and suggestions, or just a to-do list of tasks you would like to complete for this repository. </br>

* We can assign an issue to people.
* We can label issues.
* We can track issues by adding them to Github Projects.
* We can also link issues to a pull request that will fix or complete the issue if it's merged.

### Actions

Github actions is all about automation </br>
Creating a Github action will create a YAML file in your repository that allows you to setup some automation task </br>

Actions will run based on triggers </br>

This is very useful in DevOps because as you may know, DevOps is all about automation. Automation activities generally involve trigger points. </br>
Like "when this happens, run this pipeline" </br> 
For example: "when code is pushed to a repository, then:"

* trigger a build pipeline
* trigger tests
* trigger a deployment

For DevOps, automation will service a wide variety of technical aspects, such as patching and updating a server, deploying software automatically, send monitoring alerts and more </br>

### Projects

Github Projects is a page where we can manage Github issues we saw earlier and we can manage project timelines, roadmaps and kanban style boards to manage work item tasks for our repositories </br>

## VSCode integration

Now it's worth showing you that VSCode also has integration to GIT. 
You can create a branch, checkout different branches, stage and commit changes as well as pull and push </br>

## Settings

Repository settings allow us to :

* Manage collaborators, by adding people or teams to our repository.
* Branch protection, allowing us to configure rules when people push to branches, like enforcing review processes
* Github Actions, for running pipeline workflows on our repository
  - we will take a look at this in our automation guides
