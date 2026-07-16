# 🎬 Introduction to IDEs: Visual Studio Code

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 1](../../../../chapters/chapter-1-source-control-git/README.md)

## What is an IDE

An IDE stands for an Integrated Development Environment, basically a super rich code editor that is packed with features to help us write scripts, code, configuration and documentation. </br>

## What is VSCode

[Visual Studio Code](https://code.visualstudio.com/) is a streamlined code editor with support for development operations like debugging, task running, and version control. <br>

As engineers we will often work on scripts, code, configuration files and documentation and we will store these files securely in a GIT repository. </br>

VSCode is an editor that helps us work on these files as well as folder structures in an efficient manner. </br>

## Installation

We can head over to the [downloads](https://code.visualstudio.com/download) page to download it and follow the prompts for installation

## Basic Usage

Code editors usually have a lot of panels, and lots going on, which can be overwhelming for a beginner and first time user. </br>
Let's take a look at the UI. </br>

* `File` panel at the top is very similar to other programs like Office, Word, Excel, where you have to create a file, save and open a file and exit the application. Edit section allows copy paste and find \ replace functionality
  * `Selection` allows us to manipulate files like adding new lines and adding cursors to files to do some pretty cool things.
  * `View` section allows us to change the appearance of VSCode
    * We can turn certain panels on and off
    * View allows us to access some of the panels on the left 
  * `Go` is more for programming where you are looking at code - not really relevant to this course
  * `Run` is more for programming as well, as it helps run and debug programs - not really relevant to this course
  * `Terminal` is important it allows us to open multiple terminals
    - We can even split the terminal
    - We can also achieve this at the bottom right of the terminal panel

* On the left you have the directory & file `Explorer`
  * You can manage folders, files, move them around, create new ones etc
  * You can open files in the file explorer directly from here
  * Open folders directly in the terminal
  * Workspaces to visualise different projects you may be working on

* `Search` feature is really handy if you want to search for something very specific in a file
  - for example, we can search for a `docker run` command if we want to copy it
  - you can even search and replace across all files and folders

* `Debugger` - this is to debug programs - not really relevant to this course
* `Extensions` This allows us to install various add-ons for VSCode like
  - custom themes to make the editor look like we want
  - addons for different programming and scripting languages 
  - Copilot which is an AI assistant
  - There are some cool addons like Draw IO which allows us to do drawings in VSCode, just like our course roadmap

All the rest of the items on the panel was added by extensions that I have installed

## GIT features

In our previous guide we've taken a look at GIT and how to use it mostly with the command line. </br>
You can use GIT in VSCode by accessing the `Source Control` menu on the right. </br>

We need to open our GIT repo from the previous video in order to interact with GIT. </br>

In VSCode, I will demonstrate:
* How to add new file
* Make file changes
* View the diff of our changeset
* Commit the change
* Revert the change

## Settings

We also have a really in-depth settings section where we can further customise VSCode </br>
These settings are defined in JSON files which can be backed up so you can export your settings to other computers </br>


