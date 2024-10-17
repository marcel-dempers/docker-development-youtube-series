# 🎬 Introduction to GIT: Source Control

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 1](../../../chapters/chapter-1-source-control-git/README.md)

## What is GIT

[Git](https://git-scm.com/) is a free and open source distributed version control system. </br>

As engineers, we will often work on files, such as code, scripts, configuration as well as documents etc. </br>
Our daily workflow may involve:
* Create files
* Edit files
* Delete files
* Review changes
* Repeat

This will require a mechanism of tracking history of changes in these files and allowing teams of people to work on the same files without overriding each others changes. </br>

## Installing GIT

We can download [Git](https://git-scm.com/) for any operating system and follow the prompts to install. </br>

## Using GIT and Command line

During this course, I will highlight the importance of using command line as a DevOps engineer or SRE. </br>
For many development & IT professions, as well as in DevOps there are GUI tools available that have high user friendliness. </br>
These GUI tools make life easier in many ways and we will come across many such tools in this course. </br>
</hr>

For most IT proffesions, especially for software developers, they start with GUI tools first, such as Integrated Development Environments or IDE's </br>
IDE's are code editors and allows programmers to write, build, run and test code. </br>
Almost like an accountant using Excel spreadsheets or Word documents </br>
Many of these functions stop entirely, when there are no GUI tools available and developers become unproductive. </br>
The accountant stops, when Excel no longer works. </br>

In DevOps, most tooling are built from the commandline up. </br>
So although there are GUI tools available for use, when these tools are not available, DevOps engineers can still 
continue to be productive by continueing to use the commandline instead. </br>

For example, if you are on a Linux server and it ran out of disk space. </br>
There are no GUI tools on the server other than a shell terminal. </br>
In this scenario, you are forced to use the command line, so its very important to be effecient in its use. 

## Installation

To install GIT , head over to [git-scm.com](https://git-scm.com/downloads) download page and follow the installation instructions for your distribution. </br>

Linux:
```
sudo apt-get update 
sudo apt-get  install -y software-properties-common

sudo add-apt-repository ppa:git-core/ppa; sudo apt update; sudo apt -y install git
```

## GIT Command line

Once GIT has been installed, in the command line, type `git` 
```
git

usage: git [--version] [--help] [-C <path>] [-c <name>=<value>]
           [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
           [-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]
           [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
           <command> [<args>]
<and more...>
```

## Configure GIT

We firstly have to tell GIT who we are. </br>
For this we set a username and an email. </br>
This is very important especially when we are planning to host our repository on Github or another source control provider </br>

Remember the command `git help -a` which shows us all the available commands ? </br>
The one we want to start off with is `git config` </br>
note: `git config -a` Also has a `-a` to view all options </br>

We can configure GIT either globally, so which account to use for all repositories by using the `--global` flag, or only for a single repository by using the `--local` flag :

```
git config --global user.name "devops-guy"
git config --global user.email "devops-guy@test.com"
```

After running these commands you should see a `.gitconfig` file created in your user directory. </br>

```
cat $HOME/.gitconfig
```

If we try this command using `--local` we get an error stating that we can only run this inside a GIT repository </br>


## Create your first GIT repository

To turn a folder into a GIT repo its very simple. We simply run command:

```
# create a folder for our GIT repo called my-website
mkdir my-website 

# change directory into our new folder
cd my-website

git init 
```
Now our folder has been turned into a GIT repository

You can also let GIT create the folder and repository for your:

```
git init my-website
```

## ADD & Commit our first file

To add a file to our repository, we can simply create a new file. </br>
Let's create a home page for our new website called `home.html`:

```
echo "" > home.html
```

Now GIT automatically detects changes and this file is actually not part of the repository yet. </br>
To see changes incoming, we can run:

```
git status
```

We'll see one "Untracked file" </br>
To add this changes to GIT, we need to stage the file. </br>
We do this by using the `git add` command:

```
git add <filename>
#or
git add -A
```

Now the file becomes green and is ready to commit. </br>
We can commit it by running the `git commit` command.

```
git commit -m "new homepage for our website"
```

Now our file has been committed, and `git status` now shows no changes to commit. 

## Making changes 

At the moment, our `home.html` file is empty.So lets go ahead and add some content to it, and commit changes. 

Let's paste the following HTML into this file and save it: 

```
<!DOCTYPE html>
<html>
<head>
    <title>My Personal Website</title>
</head>
<body>
    <header>
        <h1>My Portfolio</h1>
    </header>
    <section>
        <h2>About Me</h2>
        <p>Some information about me...</p>
    </section>
    <section>
        <h2>My Skills</h2>
        <ul>
            <li>Skill 1</li>
            <li>Skill 2</li>
            <li>Skill 3</li>
        </ul>
    </section>
    <section>
        <h2>My Projects</h2>
        <ul>
            <li><a href="#">Project 1</a></li>
            <li><a href="#">Project 2</a></li>
            <li><a href="#">Project 3</a></li>
        </ul>
    </section>
    <footer>
        <p>Contact information: <a href="mailto:someone@example.com">someone@example.com</a>.</p>
    </footer>
</body>
</html>
```

To make the changes, again, we have to stage it first:

```
git status
git add home.html
git commit -m "add HTML to our personal website"
```

## Branching 

One big benefit of GIT is branching. </br>
GIT uses a default branch called `master`. </br>
All the work we have done up to now, was done on the `master` branch </br>

Many times, when working on projects, you may not be ready to commit changes to the repository's main branch yet. </br>
Perhaps your changes need further improvement or you are simply not done, but you want to commit your work without impacting others that also may be working on the repository. </br>

In this case , we would create a feature branch, then `git add` our changes to stage it against the feature branch, then `git commit` the changes there until we are ready to merge to `master`


```
         ┌──────────┐        
         │          │        
     ┌──►│  FEATURE ├──┐     
     │   │          │  │     
     │   └──────────┘  │     
     │                 │     
┌────┴─────┐     ┌─────▼────┐
│          │     │          │
│ MASTER   │     │ MASTER   │
│          │     │          │
└──────────┘     └──────────┘
```

To create a new branch, we will use the `git checkout` command, and the `-b` flag to set a branch name. </br>
Let's update our website HTML to set some new skills we are working on by using GIT's branching strategy: 

```
# lets create a branch called "skills"

git checkout -b skills
git status
```

Let's update our skills to set our first skill to source control:

```
 #change
 <li>Skill 1</li>
 #to
 <li>Source Version Control: GIT</li>
```

Now we commit the changes just like we did before. The more we do this, the more comfortable we will become with the comand line, especially as we go further along in this course. </br>

```
git add -A
git status
git commit -m "set our first skill"
```

## Merging Branches

To merge branches, we need to make sure we have both branches. In our case, we do have `master` and the `skills` branch.

Now for interest sake, let's checkout master. Notice when we `cat home.html` to see the content, the content still has our old skill. </br>
This is because our changes are on the `skills` branch </br>

To merge our `skills` branch into `master` we need to be on the `master` branch. 

```
git merge skills
git status
cat home.html
```
Now we can see our master branch has our skill. So technically we can delete the skills branch now </br>

To delete the branch:
```
#see all branches
git branch

#delete an old branch
git branch -d skills
```

Remember that `git branch -h` will give us "help" if we need to look at the syntax </br>

## Cloning repositories

GIT allows us to clone repositories from hosted GIT servers, like Github, Bitbucket, Gitlab and more. </br>

It's important to store your code on a hosted GIT repository as it reduces the changes of us losing our work. </br>

You can host your own GIT server too </br>
In this video, we'll be using Github and I'll show you how to clone the repo that contains all this source code and many other tutorials. </br>

Note that you dont need a Github account to clone repositories as this repo is public.

```
#make sure we are in our respositories folder

cd gitrepos

git clone https://github.com/marcel-dempers/docker-development-youtube-series.git

#show folders and files
ls 

cd docker-development-youtube-series
ls
```


## UI Tools for GIT

So far in this guide we've mostly looked at basic tools to edit folder and files as well as a basic terminal for command line work. </br>
I have done so on purpose so viewers enjoy familiarity and helps understand the course content at play. </br>

During our entire course, we will focus primarily on command line as I've outlined the importance of it. </br>
However for creating folders, editing files and having a terminal all-in-one its very important to start using whats called an IDE. </br>
An IDE is an Integrated Development Environment. </br>
Or in simple terms, its a file editor with a built-in terminal, so we can do everything we have done today, but all in one place. </br>

I will give you some home work to go and install [Visual Studio Code](https://code.visualstudio.com/) </br>

I will demonstrate in this guide, that VSCode has:

* Built-in terminal, so we can run `git` commands without leaving the IDE
* GIT extension, so we can perform `git add` + `git commit`, all without leaving the IDE
* We can create files, manage folders, all without leaving the IDE.
