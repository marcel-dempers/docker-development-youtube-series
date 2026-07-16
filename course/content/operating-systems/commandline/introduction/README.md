# 🎬 Introduction to Command Line

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 2](../../../../chapters/chapter-2-operating-systems/README.md)

## What is command line ?

[Command line](https://en.wikipedia.org/wiki/Command-line_interface) is the means of interacting with a computer or operating system by inputting lines of text. </br>
Each line being a command, I.E "Command" and "line" </br>

In order to talk to an operating system, we type commands in a Terminal Program. </br>
The terminal program basically runs the operating system shell, which is the command line intepreter </br>
The terminal basically acts as a text-based interface to a Shell program. </br>
The [Shell](https://en.wikipedia.org/wiki/Shell_(computing)) interprets commands, processes it and produces output.

In simple terms communication flows like this:

`User -> Terminal -> Shell -> Kernel`

A User types commands into the Terminal. </br>
The terminal gives us access to the shell, which takes the command, processes it by talking to the Kernal through special API's and produces output and we can see that output in the Terminal. </br> 

## Types of Shells

There are a number of shells out there depending on your operating system. </br>

On Windows we have PowerShell and CMD (Command Prompt), on Linux we have [Bourne Shell](https://en.wikipedia.org/wiki/Bourne_shell) (sh) and [Bourne-Again Shell](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) (bash). Linux also has other shells like Z shell, or zsh. </br>

In this module, I will showcase an example of Windows Shells and Linux Shells. </br> 
So you understand the difference between Terminals and Shells </br>

## Types of Terminals

On Windows as we can see there is a built-in terminal that can run either Powershell or CMD Command Prompt. Windows also launched a newer terminal called [Windows Terminal](https://github.com/microsoft/terminal) which is available on Github. </br>

On Linux, we get different terminals like GNOME Terminal, Terminology, Terminator and more </br>

In this demo, I will show you an example of Windows Terminal and Terminator on Linux </br>

## Why Learn the Command Line

### Troubleshooting

Command line is very important in times when there is no access to Apps \ GUI's or User Interfaces. </br>
Most of the time, when accessing production servers, you may only be able to access it via a terminal through protocols like SSH. </br>

Some examples include:

* When a production server runs out of disk space, you need command line to troubleshoot it
* When services are unable to reach network endpoints, you need command line to test network connectivity

Command line could also be required when certain systems are mostly driven by API's. </br>
For example, some Microsoft Azure or Amazon AWS services have more options to configure via their API's and their GUI portals are often playing catch up. </br>
API's are generally built before UIs, meaning UI's may be behind or may never support certain capabilities. </br>

### Scripting and Automation

Another reason to learn command line, is that it gives you a foundation to start automating from. </br>
When you want to automate something, you gather and stitch together all the commands you need and then write it in a script which allows you to run the script to perform many commands at once </br>

Some examples may include:

* We could stitch together commands that check if folder structures exist  or not and create them if they don't exist. Then automatically create configuration files with default settings in them to setup an application
* We could stitch together commands that make a copy of important files and folders on our system, compress them as a "zip" file and upload it to remote storage to create an automated backup process. 

## Command Line Structure

As a beginner, one of the most important things to learn about is the structure of command lines </br>
Generally, all or most commands conform to some standard. </br>

### Executable 

For example, all command lines start with some form of executable or binary. </br>
A binary executable is a program or script. </br>
For example in our Introduction to `git` , we learned about the `git` command. In that case `git` is the executable. 

So we start off our command with the executable:

```
git
```

### Commands and Subcommands

An executable may have commands, and commands by have subcommands. </br>
However, some executables may have no commands. This is generally true for command line executables where they generally only perform one thing, like `ls`

If you run `ls --help` , you'll notice the usage:

```
Usage: ls [OPTION]... [FILE]...
List information about the FILEs (the current directory by default).
```

`ls` has no commands. </br>

Commands or Subcommands are generally after the executable. For example:

```
<executable> <command> <optional-subcommand>

```
GIT on the other hand, has many commands and some commands also have subcommands. </br>

For example:

```
git checkout
git pull 
git commit
git push
#etc
```

### Arguments, Options, Flags or Parameters

When a command needs inputs, the command may require arguments, options, flags or parameters. </br>
Those terms are often used interchangeably. </br>
Command line options generally start with dashes. </br>
Generally a single `-` character indicates a short abbreviated form of an option and a double `--` is a long form version of the same option. </br>

For example, `-h` and `--help` generally mean the same thing, but one is abbreviated, so shorter to type out and the other is longer, which means it's more descriptive but longer to type out. </br>

If we take a look at `git` that we learned about:

```
git commit -m "commit message here" 
```

Above, `-m` stands for message. 

## Collection of important commands

* `echo` : This command allows us to print text to the screen or print variable contents
    - For example, `echo "hello"`, or `echo $HOME`
* `cat` : This command displays file contents. It can display one or multiple files contents and concatenates them
    - For example `cat <filename>`
* `ls` : Allows to list contents of a directory
* `cd` : Change working directory
* `pwd` : Writes the full path of the current working directory to the screen
* `touch` : Creates a file without content
* `nano` : Starts a simple Linux text editor
* `mkdir` : Create a directory
* `cp` : Copies a file or directory
* `mv` : Renames a file or directory. Useful for renaming files or moving content
* `rm` : Removes a file or directory
* `grep` : Find text in a file or in output
* `df` : Displays the amount of disk space available on the file system
* `du` : Estimates file space usage for a directory
* `ps` : Displays information about active processes. `ps aux` displays all processes
* `top` : Command line utility to monitor system load
* `kill` : Terminates a process by process ID
* `wget` : Downloads files from the web
* `curl` : Create and send a Web request to a server
 
## Command inputs, outputs and redirection

We can capture outputs of a command. </br>

This becomes useful for two main reasons. </br>
Firstly, we can use the output of one command as input for another. Technically this allows us to chain commands together. </br>
This is useful in automation and scripting when taking output of one command execution and passing it as input to another </br>

Secondly, capturing output of a command can be useful to store in a file for troubleshooting purposes or any other reason. </br>
For example, a command's output may be quite large to analyze in a terminal window and may be useful to analyze as a file which we can open in an editor such as Visual Studio Code. </br>

### Command pipe

To capture output of one command and send it as input to another command, we can use the "pipe" character. </br>
This often is referred to "piping" output to another command in command line </br>

For example, we can use `cat` to output file contents and then `|` it to `grep` to search that content for a given string pattern. </br>
This often helps looking for something in a large file. </br>

```
cat <filename> | grep <pattern>
```

### Command redirection

To redirect command line output to a file, we can use the command redirection character `>`

For example when running `ls`, it outputs to the terminal. </br>
We can capture that output by placing the `>` character after the command followed by a filename that we want to pass output to </br>

```
ls -l / > output.txt
```

The single `>` will override the file and replace all its current content with the input provided. </br>
If we want to preserve the current content and simply append input to it, we use the double `>`
For example:

```
echo "add more content" >> output.txt
```
