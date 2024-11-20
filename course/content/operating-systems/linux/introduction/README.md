# 🎬 Introduction to Linux for beginners

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 2](../../../../chapters/chapter-2-operating-systems/README.md)

## What is Linux

Linux is a family of operating systems based on the Linux kernel. It's the most popular operating system on the planet that powers most of the internet today as well as Android which is powered by the Linux operating system. </br>

It's important to note that Linux comes in many [distributions](https://en.wikipedia.org/wiki/List_of_Linux_distributions) or Distros. a Distro is the Linux kernel, packaged alongside many supporting operating system software and libraries.
There are many popular Linux Distros, including Ubuntu, Debian, Arch Linux, Fedora and more. </br>

## Setup a Linux Server

For this module, you will need access to a Linux server. </br>
See the Preface section above for a link to the chapter overview page where you will find the module on Servers & Virtualization for creating a Virtual Server. </br>

Checkout [chapter 2: module: Introduction to Servers & Virtualization](../../../chapters/chapter-2-operating-systems/README.md) if you need to create a Linux server </br>

## Terminal, Shell & Command Line

The first thing we see when we start this server is that our software gives us a window where we can type commands. </br>

The window we see is called the Terminal. The terminal is just a window that accepts inputs and outputs </br>

The Terminal is responsible for running what's called a Shell. </br>
A Shell is a program that exposes the operating system to a user or a program. So we talk to the operating system by typing commands into the Shell. The Shell is the command intepreter which processes our commands and outputs the results. </br>

## Users & Security

When we created our server, we had to provide details of a username and password for our default user. </br>
This is a standard practice for setting up Linux. </br>
Linux by default is shipped with a `root` user. The `root` user is the administrative account of the operating system which have very high privileges. </br>
You are discouraged from using the `root` user because it can make almost any change to the operating system which can be disastrous or have catastrophic effects. </br>

### SUDO & Administrative tasks

In order to do administrative tasks on a server, instead of using the `root` account, we can setup whats called `sudo` </br>
`sudo` is a program that enables users to run programs with privileges of other users, by default the super user. </br>

SUDO originally stood for "superuser do" as it allowed us to do something as a super user. </br>
Then it became "su do", as in "Substitute user, do", because `sudo` can run commands as other users as well. </br>

To see `sudo`, type `sudo --help` </br>
We can run `sudo -i` to switch to `root` and `exit` to go back to our account

### Add a user

To add user: (We see only root account can do this)

```
adduser bob
Fatal: Only root may add a user or group to the system.
```

We elevate our privilege by using `sudo` in the front (and follow the prompts)

```
sudo adduser bob 
```

Bob is now an unprivileged user. </br>

We can make Bob privileged by adding him to the SUDOers group then Bob can use `sudo` to elevate privileges just like our user. 

```
# see the help text
usermod --help

# we are going to use -a which means append
# we are going to use -G which means we want to append to a group
usermod -aG sudo bob

# we get permission denied! remember to use sudo
sudo usermod -aG sudo bob
```

### Delete a user
To delete a user, we simply run `userdel`

```
userdel --help

#remove bob
sudo userdel -r bob
```

### List users

User information in Linux is stored in the `passwd` file which exists in the `/etc` directly

View the content of the file with the `cat` command

```
cat /etc/passwd
```

Another useful command to see who we are logged in as, is the `whoami` command

## The File system & Navigation

In order to navigate around, we will learn a couple of commands to navigate, but we also need to know what we are navigating </br>

* `ls` is a command we use to list contents in a directory. `/` id the root directory in Linux, which is the equivalent of `C:\` in Windows. Type `ls /` in the terminal to list out contents in the `/` root directory.

* <b>Home directories:</b> Under the `/` directory you will see a directory called `/root` which is the `$HOME` directory of the root user. </br> All other user folders are under `/home` </br>
The `~` also represents the full path to the home directory for your current user

* The `.` is a short-hand for the current directory. Example: `ls .`

* `echo` prints stuff to the terminal. One example is to see the content of the $HOME environment variable, we can type `echo $HOME` to see it </br>

* Environment variables are something we will dive into an our command line and scripting modules, but some variables are important as we have already seen with `$HOME`. Another one is `$PWD` which contains the current working directory. </br>

* <b>Create a directory:</b> we can create a directory with `mkdir` </br>
  Create a directory in your HOME directory: `mkdir $HOME/testdir`
* <b>Change directory:</b> We can change the current working directory by using the `cd` command. </br> 
  I.E `cd $HOME`
* <b>Delete a directory:</b> we can create a directory with `rm` </br>
  Let's delete our test directory: `rm -rf testdir`

### File management

Let's change directory to our $HOME folder

```
cd $HOME
```

We can see where we are

```
echo $PWD
```

Create a new folder where we can keep our scripts

```
# create directory
mkdir scripts

# view directory
ls -l

# change directory
cd scripts
```

Create a new script file with the `touch` command
We can use the `touch` command to create any type of file

```
touch my-first-script.sh
```

We can see the file is empty
```
ls
cat my-first-script.sh
```

We can edit the file using `nano` which is a simple text editor for Linux

```
nano my-first-script.sh
```

We can now write content in the file </br>
Let's add the following content to the script file:

```
echo "This is our first script running from directory $PWD !"
```

To exit `nano` and save the file, we press CTRL+x and type `Y` to save the file.

In our scripting module, we will look at running scripts in more details , but quickly:

```
# try run the script
./my-first-script.sh
-bash: ./my-first-script.sh: Permission denied

# needs exectute permissions
chmod +x my-first-script.sh

# try run the script again
./my-first-script.sh
This is our first script running from directory /home/devopsguy/scripts !
```

To copy a file, we can use the `cp` command. Lets make a copy of our script

```
cp --help
cp my-first-script.sh my-second-script.sh
ls -l 
```

We can delete the second file with the `rm` command
```
rm my-second-script.sh
``` 

We can move or rename a file with the `mv` command
```
cp my-first-script.sh my-final-script.sh
```

## Package managers 

To install things on Linux we need to learn about the package managers. </br>
Ubuntu Linux maintains sources list for where it downloads packages by default </br>
The sources list contains URLs or Links to Mirror sites or Repository sites where packages are looked for when trying to install things </br>

To install a package :

```
sudo apt-get install -y <packagename>
```

To remove a package : 

```
sudo apt-get remove <packagename>
```
