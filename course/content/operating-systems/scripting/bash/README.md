# 🎬 Introduction to Bash scripting for beginners

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 2](../../../../chapters/chapter-2-operating-systems/README.md)

## What is Scripting

[Scripting](https://en.wikipedia.org/wiki/Scripting_language) is a process of creating a script file that contains a simple set of  instructions that automate a process that otherwise would be manual. </br>

For example, in previous modules we learned about commands. </br> As an engineer, we may have tasks that require us to run commands on servers. </br>

Some examples of tasks:
* Backup files on a server and copy it to remote long-term storage
* Update servers with latest operating system security patches and updates
* Deploy a website to a target server

To perform these tasks, you may automate access to the server via SSH or through the use of DevOps tooling. </br>
You may need to use a combination of commands like `cp`, `mkdir`, `tar` and potentially commands to send backup files to remote storage. 

You may spend time creating a list of commands that you need to run in order to accomplish this task. </br>

Scripting is the process of taking these commands and placing it in a script that can be executed. Which means we dont have to manually run all the commands one by one, and we simple execute the script instead. </br>

## Benefit of Scripting

There are a few benefits of scripting. </br>
The first one being the automation of tasks I highlighted previously. The second being that scripts can be portable, so we can take a script and run it on many servers. </br>

This benefits a team of engineers because there is a one-time time investment to create the script and then allow a team of engineers to use the script and save time. </br>
Therefore it allows teams to be more agile by automating boring repetitive tasks and letting engineers focus on things that matter </br>

## Cons of Scripting 

There are two main negative points or "cons" of scripting that I had dealt with over the years and still continues to be a challenge when writing scripts. </br>

### Imperative programming

Imperative programming is a paradigm where a programmer explicitly tells the computer how to achieve a desired outcome. 

### Idempotency

An Idempotent operation is an operation that can be performed multiple times without changing the result beyond the initial application


## Setup a Linux Server

For this module, you will need access to a Linux server. </br>
See the Preface section above for a link to the chapter overview page where you will find the module on Servers & Virtualization for creating a Virtual Server. </br>

Checkout [chapter 2: module: Introduction to Servers & Virtualization](../../../chapters/chapter-2-operating-systems/README.md) if you need to create a Linux server </br>

Now before firing up our server, I will setup a shared folder between our local GIT repository we created in [Chapter 1](../../../../chapters/chapter-1-source-control-git/README.md)

To do this, we will use the `vboxmanage` command we used in earlier modules to create our networking for our virtual servers. </br>

Let's setup a shared folder for our server so we can play around with our scripts locally and run it directly in a Linux server. This will make it convenient for us to test our script and we don't need to copy it to the server all the time. </br>

<b>important: take note of the host path which is the path to your GIT directory we created in an earlier chapter</b>

```
# Windows:
VBoxManage sharedfolder add "my-website-1" --name GIT --hostpath "C:\gitrepos\"

# Linux:
VBoxManage sharedfolder add "my-website-1" --name GIT --hostpath "~/gitrepos"
```

## Access our Server

In our Introduction to Servers as well as Linux, we learned how to SSH into our server to manage it remotely. </br>
Let's go ahead and do that again:

```
ssh devopsguy@127.0.0.1 -p 2222
```

Now that we're in, we can setup a mount point to our shared folder, so our Linux server can access the shared folder as a directory on this server. Let's run the following command to do so:

```
mkdir ~/gitrepos
sudo mount -t vboxsf -o uid=1000,gid=1000 GIT ~/gitrepos
echo "GIT $HOME/gitrepos vboxsf defaults,uid=1000,gid=1000,rw  0  0" | sudo tee -a /etc/fstab
```

Above, the `mkdir` creates our directory we want to use as a mount point </br>
The files we have on our machine will be mounted into this folder in the virtual server </br>
The `mount` command performs the mount operation that will mount the `vboxsf` file system managed by VirtualBox called `GIT` (That we created with `vboxmanage`). </br> 
It will mount it to the newly created `~/gitrepos` directory. </br>
Now we should be able to see our GIT repositories on our virtual server when running `ls ~/gitrepos` </br>

However when the server reboots, mount points are not persisted, so we need to create an entry in a special Linux file that defines mount points in Linux, called `etc/fstab` </br>


Let's change folder to our GIT repo:

```
cd ~/gitrepos/my-website
```

## Creating our first Bash Script

To complete this tutorial, we will need to fire up our Virtual Server that we created in a previous module with Virtual Box. </br> This will give us the ability test out our script in a Linux environment. </br>

I would highly recommend to follow along as per our course video so you get the experience of writing your own script and experimenting with tweaks along the way </br>

### A shell script

Let's create our first Bash script. </br>
In a previous chapter we created our Github repository called `my-website`. </br>
Let's open that up in Visual Studio Code just like we have done before in previous modules. </br>

In our repo let's create a `scripts` folder.
In our new  `scripts` folder, we can create a new script called `setup.sh`
We can use this script to automate the setup of our server and learn about scripting in the process. </br>
Shell or Bash scripts generally have a `.sh` as a file extension to indicate a "script file".

### Understanding the Shebang in the script

Let's add this first line to our script:

```
#!/bin/bash
```

a Shell script or Bash script starts with the [Shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)) in the script file, right on the first line. </br>
This is technically just an indicator for when a text file (like our script) were used as if it were an executable. When it gets executed, a few things happen:

* The operating system program loader loads up our script and determines intepreter to run
* The whole `#!` line gets ignored by the intepreter because its technically a comment because it starts with `#`
* The program loader knows which interpreter to load indicated by the path in the Shebang, for example: `#!/bin/bash` will load the `/bin/bash` program as interpreter and our file will be passed to it.

### Executing our script

So to learn how to execute the script, lets just add a simple command to the script that outputs a message:

```
echo "hello from our script!"
```

To execute our script we can use the `./` characters followed by the path to the script file. </br>
The `./` character allows us to run a script in a new subshell. </br>

So in our terminal where we are connected to our server, we can run a script by saying `./<script-name>.sh`. </br>
So our command to run the script will be: 

```
# change directory to where the script is
 cd ~/gitrepos/my-website/scripts

 # run the script
 ./setup.sh

 #notice the error
 -bash: ./setup.sh: Permission denied
```

We should get an error in doing so because our script does not have execute permission. </br>
This is a security permission so arbitrary malicious processes cannot execute any scripts on your machine without our consent<br/>

Therefore we need `sudo` to set execute permission on our script and we do so by using the `chmod` command. `chmod` takes a permission. </br>

If we do `ls -l` in the directory of our script we can see its missing execute permission. </br>

#### First character

Entries that start with `d` are directories or folders. </br>
Entries that start with `-` are files </br>

#### Next 9 characters (Permission string)

* `r`: Read permissions. File content can be viewed
* `w`: Write permissions. File content can be edited
* `x`: Execute permissions. The file can be executed if its a script or binary program
* `-`: No permission

Permissions are shown in three groups, each group has 3 characters, making up a total of 9 characters for the permission string. </br>
* The first set of 3 characters represent permission of the owner
* The second set of 3 characters represent permissions of the group
* The last set of 3 characters represent others permission


To give our script execute permission `x` , we run `chmod +x setup.sh`

```
sudo chmod +x setup.sh
```

Now we can execute our script 

```
./setup.sh
hello from our script!
```

## Sripting 

Now we can finally start scripting. </br>
In this module, we'll extend our `setup.sh` script which we will use to setup our website content and basically deploy our website files to a destination folder where a web server could be running from. </br>
This is a common traditional method used in deploying website files to a server, so I will walk through some of these techniques and we'll learn about Bash scripting along the way. </br>

This script may prove useful in future modules, so once we are done, we will store it in our GIT repo </br>

### Variables and strings

Variables are simply pointers to information </br>
We give a variable a name, and then we set a value that the variable will hold. </br>

This way we can store valuable information that our script needs. 
We can pass variables to a script as parameters or we can set variables at the top of our script. </br>

In previous modules we used `$PWD` and `$HOME` as variables </br>

An example for our `setup.sh` script might set some variables at the top of our script and use that as settings that we may want to control. </br>

Let's create some variables in our `setup.sh` script file

The first varialbe points to our Gthub source code for our website we would like to deploy. </br>
The second and third variable indicate where we will download our source code to and the destination where we will be copying the files to. </br>
This destination folder may be where our web server could be serving the files from </br>

```
# settings
GITHUB_REPO_URL="https://github.com/marcel-dempers/my-website.git"
DEPLOYMENT_SOURCE_DIR="$HOME/gitrepos/my-website"
DEPLOYMENT_DEST_DIR="/webites/my-website"
```

Printing logs our output is very useful in scripting. </br>
This helps us know if our script is working successfully and gives us feedback. </br>
For this, lets use the `echo` command to print things to the output. A useful start is to print out that the script has started and what task its performing </br>

```
echo "starting our deployment script..."
```

### Making descisions & Conditional statements

Scripting and automation is all about following logic. When writing logic we mostly make descisions. </br>
For example, we often make descisions like:

* If a file does not exist, create it
* If a variable is empty, throw an error
* If a server is out of date, update it
* If "something is x" , "perform y" , else "perform z"

Descision trees make up most scripts in traditional automtion concepts </br>

Let's take a look at `if` \ `else` statements

#### IF/ELSE Descision trees 

```
if [ expression ]; then
    # commands go here
fi

###############################

if [ expression ]; then
    # commands go here
else
    # other commands
fi

###############################

if [ expression ]; then
    # commands go here
elif
    # other commands
else
    # last commands
fi
```

Let's use an `if` block to validate if our settings have been set correctly </br>
We call this validation. We will often check various scenarios and criteria to ensure our script don't fail. </br>
In our case we can check if our environment variables are empty and stop the script with an appropriate error message if so. </br>

Let's start with a message what we are about to do:

```
echo "checking website source directory: $DEPLOYMENT_SOURCE_DIR and destination directory: $DEPLOYMENT_DEST_DIR..."
```

Throw an error if our variables are empty (null)

```
# Validate source and destination directories
if [ -z "$DEPLOYMENT_SOURCE_DIR" ]; then
  echo "ERROR: source directory is not set!"
  exit 1
fi

if [ -z "$DEPLOYMENT_DEST_DIR" ]; then
  echo "ERROR: destination directory is not set!"
  exit 1
fi

if [ -z "$GITHUB_REPO_URL" ]; then
  echo "ERROR: github repository url is not set!"
  exit 1
fi
```

#### Conditional statements

The conditional statement is the expression in the IF/ELSE Descision tree.
For example: "If a directory does not exist, create it" , "If a file does not exist, create it" , or "If a file exists, throw an error" 

The simplest form of a conditional expression is testing if a variable is equal to some value

```
if [ "$TEST" = "" ]; then
  echo "TEST is correct"
else
  echo "TEST is not correct"
fi
```

One benefit of writing scripts is that we can simply paste the above bash into a terminal and run it without needing a script. This is very useful when developing a script as it helps test our commands</br>


#### Conditional Operators

File Test Operators

* `-e FILE`: True if the file exists.
* `-f FILE`: True if the file exists and is a regular file.
* `-d FILE`: True if the file exists and is a directory.
* `-h FILE`: True if the file exists and is a symbolic link.
* `-L FILE`: True if the file exists and is a symbolic link.
* `-r FILE`: True if the file exists and is readable.
* `-w FILE`: True if the file exists and is writable.
* `-x FILE`: True if the file exists and is executable.
* `-s FILE`: True if the file exists and is not empty.

Example:

```
# check if a directory exists
if [ -d "$DIRECTORY" ]; then
  echo "$DIRECTORY exist."
fi

# check if a file exists
if [ -f "$FILE" ]; then
  echo "$FILE exist."
fi
```

String Operators can be used when dealing with strings and variables

* `-z STRING`: True if the string is empty.
* `-n STRING`: True if the string is not empty.
* `STRING1 = STRING2`: True if the strings are equal.
* `STRING1 != STRING2`: True if the strings are not equal.

Numeric Comparison Operators

* `NUM1 -eq NUM2`: True if the numbers are equal.
* `NUM1 -ne NUM2`: True if the numbers are not equal.
* `NUM1 -lt NUM2`: True if NUM1 is less than NUM2.
* `NUM1 -le NUM2`: True if NUM1 is less than or equal to NUM2.
* `NUM1 -gt NUM2`: True if NUM1 is greater than NUM2.
* `NUM1 -ge NUM2`: True if NUM1 is greater than or equal to NUM2.

We can use a combination of the above in our deployment script to do a few things:

* Check to ensure the source and destination directories have been set
* Check to ensure the source and destination directories exists

We can throw an error with a message if the above checks fail. 

Let's setup a soure folder where we will sync the latest changes for our website </br>

```
# setup our source for our website

if [ ! -d "$DEPLOYMENT_SOURCE_DIR" ]; then
  echo "source directory does not exist, creating... "
else 
  echo "source directory exists, pulling latest changes..."
fi

```
Now our script checks if our source folder exists or not and creates it if it doesnt, else it does not create it. </br>
Let's use `mkdir` to create our folder

```
mkdir -p $DEPLOYMENT_SOURCE_DIR
```

If the source directory does not exist, we create it as above, and then we can use `git` to clone our website repo, for the first time:

```
git clone --depth 1 $GITHUB_REPO_URL $DEPLOYMENT_SOURCE_DIR
```

And a message to say we have done this successfully:

```
echo "source directory created."  
```

If the folder already exists we dont need to create it and we can assume our code is there, as we would have cloned it in a first run of our script. Therefore we can use `git` to pull the latest changes.

```
git pull origin master
```

And a message to display that we updated our code

```
echo "source directory updated."
```

### Deploying our website code

<p>
Now that we have our website source synced to our server, we want to deploy it. </br>
Usually a website runs from a specific folder that a web server is configured to serve. </br>
In a future chapter, we will cover Web Servers in detail, but its important understand at a  high level as it helps us write a very realistic automation script that we can actually apply to real world concepts. </br>
</p>

<p>
A web server has a configuration file that tells the server where to find the website files to serve. </br>
Traditionally, we would create a folder with our website code and then update the configuration file to point to that folder. </br>
When we deploy changes, we create a new folder with the latest website code and update the configuration file to point to the newly created folder </br>

This way we are not updating the current files being served directly as it may impact and disrupt our web traffic. </br> 
If the web server is reading that file while we update it, thats not good </br>
Also, if we deploy a new folder every time its very easy to rollback if we notice an issue with the new code. We simply update the web server configuration file to point it back to the previous release. </br>
</p>

Let's make our script take our source that we synchronized, and copy it to a new folder under the web server location and update our web server configuration file with the new release </br>

We start with a message:

```
echo "deploying latest changes..."
```

We want to create a unique folder per release. It's often a popular practise to use timestamps to create unique folders. In Linux there is a `date` command that gives us a date in a format we prefer. 

Also, we can run bash commands and capture its output into a variable by using `$()` 
```
TIMESTAMP=$(date +"%Y%m%d%H%M%S")
```

Then we use this to build up a string path to our destination directory where we want to copy our source and create the directory

```
NEW_DEPLOYMENT_DIR="$DEPLOYMENT_DEST_DIR/$TIMESTAMP"

mkdir -p "$NEW_DEPLOYMENT_DIR"
```

Now that we have our target folder set, we can use `cp -r` to recursively copy our website source to the web server destination

```
# Copy source files to the new directory
cp -r "$DEPLOYMENT_SOURCE_DIR"/* "$NEW_DEPLOYMENT_DIR"
echo "deployed website to $NEW_DEPLOYMENT_DIR"

```

### File manipulation

Now that we have our website files deployed in a unique folder location, we want to tell our web server to serve those new files and not the current ones. </br>

To do this, we need to automatically update the web server configuration file. Generally these configuration files have a directory path indicating where to serve files from. <br>

For example, a popular Web server called NGINX uses a configuration file that looks like this:

```
server {
  listen 80;
  server_name my-website.com;
  root /path/to/website/;
}
```
Now this is not a web server tutorial so we won't get into details about that just yet, but basically our script just needs to create this configuration if it does not exit, and update it if it does exist. </br>

Starting with a message:

```
echo "updating website configuration..."
```

We may introduce a filename variable that tells our script where the configuration file is:

```
CONFIG_FILE="$DEPLOYMENT_DEST_DIR/nginx.conf"
```

Now we will need another `if` and `else` statement. </br>
If the config file exists, update it, else create it. That should be simple for us now:

```
if [ -f "$CONFIG_FILE" ]; then
  echo "updating website configuration..."
else
  echo "creating website configuration..."
fi
```

To update our config, we will use a command that allows us to search for TEXT in a file and substitute it with the variable that contains our new PATH. <br>

In linux, this command is called `sed`. Run `sed --help`

```
sed -i "s|root $DEPLOYMENT_DEST_DIR/.*;|root $NEW_DEPLOYMENT_DIR;|" "$CONFIG_FILE"

echo "website configuration updated."
```

If the file does not exist, we simply need to create it. We can use `echo` and direct the `echo` output to a file as we learned in our previous module about output and redirection

```
echo "server {
    listen 80;
    server_name my-website.com;
    root $NEW_DEPLOYMENT_DIR;
  }" > $CONFIG_FILE

echo "website configuration created."
```

And finally a message to say its all done:

```
echo "deployment complete"
```

### Loops

There are two main loops in bash that I use the most. </br>
For loops & While loops. </br>

For loops follow a syntax of "For every value in this list , do this"
While loops follow a syntax of saying "While this thing is true, do this" 

You may notice a distinct difference here. With For loops we are very distinct about a list of things we want to loop through. Because we said "For every value in this list" </br>
This type of loop is useful when looping through files or directories or when you have some data in a file or a distinct list of things to iterate on. </br>

With While loops, we loop forever while a condition is true. </br>
For example, While all existing connections to my webserver is not zero, continue to wait </br>
Or While something is not ready, sleep for 1 second and try again. </br>
Or While a value in a list of file is not found, keep looping until the end of the list </br>

Basic For loop syntax

```
for variable in "list"
do
  #do some work with $variable
done

```

The list could be any type of list. For example:

```
for number in 1 2 3 4 5
do 
  echo "number: $number"
done

for file in $(ls)
do 
 echo "file or folder: $file"
done
```

As you see above `$()` is an expression. It allows us to execute commands and use its output in the loop.  </br>
This is very handy and I will show you why
This results no result: `$(echo "test")` </br>
This is because the command runs inside `$()` and that result is not printed in the terminal
Instead we can grab the result into a variable: `result=$(echo "test")` and the `echo $result` to see the output.

This is very useful because you can run useful commands that generate useful output and assign those to variables which we can loop over or process in different ways </br>

Let's take a look at an example `while` loop 

```
counter=1
while [ $counter -le 10 ]
do
  echo $counter
  ((counter++))
done
```

### Inputs and Outputs

#### Inputs
In Bash we can read inputs from a user by using the `read` command

```
read variablehere
echo "You typed: $variablehere"
```

We can use the `read` command to allow user to enter values for a script. </br>
If you have a process you need to kick off manually that may require secrets like passwords or authentication, you may use the `read` command to ask the user to type this information in. </br>
This means we don't store sensitive information in our scripts. 

<i>Important note: DevOps engineers should always avoid storing sensitive information such as API keys, tokens, passwords and authentication details in scripts.
Auotmation tools for CI/CD & automation pipelines generally allow for a way to pass sensitive details to scripts</i>

#### Outputs

The pipe operator or `|` is used to pass the output of one command as input to another command.
For example, we can grab the content of our website config

`cat /webites/my-website/nginx.conf`
Let's say this is a large config file and we only want the `root` section we can pipe it to another command called `grep`.

[grep](https://en.wikipedia.org/wiki/Grep) is a command-line utility for searching plaintext datasets for lines that match a regular expression

`cat /webites/my-website/nginx.conf | grep 'root'`

This command and pipe is useful if you have to search a large file for some specific text. <br/>

### Functions

Functions allow us to group certain commands together to perform a specific task or function </br>
For example, if you want to restart a service, or update a file or perform any specific isolated task, you can wrap that logic into whats called a `function()` 

#### Why Functions

As you write scripts, you will quickly find out that the script grows in size and as it grows, it also grows in complexity, and becomes difficult to read. </br>
As this complexity increases, your script becomes more prone to failure and errors and becomes harder to test </br>

You can split code up into smaller logical chunks by using functions. </br>
It not only splits up complex long scripts, but also makes it easier for humans to read and understand as they only need to focus on relevent overall functionality of the script when reading it </br>

Another benefit of writing functions is each function can be tested in isolation, or should be testable in isolation. This makes the author think about tests, how to pass inputs and what the desired output of a function should be so that it can run and be tested without having to run the whole script </br>

#### Function definition

Let's take a look at the structure of a function:

```
functionName() {
    echo "this runs within a function"
}
```

You can execute this function by just running command: `functionName`
The cool thing about functions is you can simply paste them into your terminal and then run them to test them out. </br>

#### Function variables

When working with scripts and functions its important to learn about variable scopes </br>
Generally when defining a variable in a script, its scope is the entire script, meaning its available everywhere in the script. </br>

You will generally see me define variables for my script at the top, which allows readers to see them all in once place </br> Just like our script we have been working on, we can see the source and destination directories for our script and allows the reader to easily change it without having to read through the entire script and find the variables. </br>

I would generally try to avoid tieing up these variables with functions. </br>
As soon as a function starts using variables that are defined elsewhere in your script then that function can no longer be tested in isolation because it relies on variables to be defined </br>

We can pass many inputs to a function by positional parameters after the function name, for example:

```
functionName firstParameter secondParameter thirdParameter
```

For our function to use these inputs we can simply access variables by their numerical position. </br>
For example, `$1` , `$2` or `$3` where the number represents the position of the parameter.

Let's try that:

```
functionName() {
    echo "this runs within a function and prints 1: $1 , 2: $2 , 3: $3"
}
```

We can also set `local` variables for our function. This is great as it keeps these variables available to the function only. 
For example:

```
functionName() {
  local first=$1
  local second=$2
  local third=$3
  
  echo "our first variable is: $first"
  echo "our second variable is: $second"
  echo "our third variable is: $third"
}
```
