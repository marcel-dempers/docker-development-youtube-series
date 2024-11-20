<span style="color:yellow">
</span>

🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦 C:\Users\aimve\Videos\video-projects\videos\footage\

<span style="color:yellow">
[seg-1]  - [desk]
  - Welcome to chapter `2`, and a new module that coveres the introduction to Linux
  - In the chapter overview I talked about servers and the importance of servers in DevOps.
  - We'll most of the internet are run by Linux servers
  - Linux is extremely important in DevOps, SRE, platform & cloud engineering.
  - [pc]
  - let's start of with the course material
  - in the course notes you will find the link to my Github repo
  - when you open this link it will take you to github and you will see many files and folders
  - what you are looking for is the course folder.
  - i want you to get comfortable with navigating and familiarise yourself with the course content
  - in the course folder is the course landing page which has table of contents 
  - and that takes you to each chapter. we are interested in chapter 1 so click that
  - each chapter lists out a number of modules so in this video we cover ... so click that 
  - and now you are at the course content for this lesson.
  - [side]
  - Now throughout this course we will be following practical steps, so i wont just be 
  - explaining theories and concepts , but I also urge you to follow along
  - the benefit of this course is that we will be building you
  - a personal website profile using DevOps tools, principals and practises
  - to help you learn devops concepts.
  - so at any time, pause this video if you need to and try to follow along
</span>

[seg-1-yt]  - [desk] 
  - Linux is a very important operating system to learn about in DevOps, SRE, cloud and platform engineering.
  - Linux makes up most of the internet
  - Over 96% of the top one million web servers use Linux
  - the Cloud is made up of around 90% Linux servers
  - it runs all of the worlds 500 fastest super computeres and runs the world stock exchange.
  - I think i've given many good reasons why Linux is important,
  - But the reason its very important for you when learning DevOps is because much of the servers you may  deal with in your career, may be running linux.
  - We need to know how to work with linux
  - How to setup
  - How to configure
  - How to monitor
  - How to secure
  - and effectively manage linux servers 
  - So today we take a look at Linux for beginners 
  - we'll create a linux server
  - learn about the terminal, shell & command line and some of the linux basics
  - we've got a lot to cover to
  - without further adeau, lets go  

<span style="color:yellow">
Importance for Linux - learning docker containers and kubernetes
</span>

[intro]
<span style="color:yellow">play course placement</span>

[seg-2]  - [pc]
  - show course material 
  - [side] 
    - so be sure to checkout the link down below to the readme file so you can follow along
  - [pc]
  - "so what is linux" - read (what is linux first paragraph)

<span style="color:yellow">walkthrough distributions chart</span>
  - [side] 
    - Linux comes in many distros, or distributions
    - Think of a distribution as a flavour
    - You have the core of linux which is the burger, we all know what a burger is made of 
    - then you have many flavors of burgers, like mcdonalds, burger king, five guys, etc
    - read (distro section)

[seg-3]  - [desk]
  - now to learn about Linux, you will need a linux environment
  - It's very important to have a linux environment on hand
  - which you can use as a playground
  - if you mess it up, you can delete it and start again
  - without affecting your own computer
  
  - In a previous video we talked about servers and virtualization
  - And I showed you how to create a linux server using virtualization software.
  - If you dont have a server, check the preface in the course material
  - which will take you to an earlier video on how to create a linux server

  - [pc]
  - show preface
  - show a server quickly
  - show SSH from VSCode 

<span style="color:yellow">
importance of having a linux server you can create and destroy
  - this is why the module on servers and virt is so important
  - use it as a baseline to :
- automation - infrastructure as code 
- we'll use it as practise
- you will likely create your server over and over
</span>


[seg-4]  - [desk]
  - [pc]<span style="color:yellow"> discuss (Terminal, Shell & command line section)</span>
  - When working with linux we are immediately faced with a window
  - where we have to type commands.
  - This window is called a terminal
  - It accepts inputs and outputs

  - The terminal is responsible for running whats called a Shell
  - a Shell is a program that exposes the operating system to users
  - we dont type commands directly to the operating system
  - we type it into the shell
  - the shell is a command line intepreter which inteprets our command, processes it
  - and outputs the results
  
  - [pc]<span style="color:yellow"> discuss (Users & Security)</span>
  - Now when running commands
  - You have to run commands as a user
  - When you installed Linux, you would have been prompted to create a user account and set a password
  - This is the account you are encoured to use
  - Linux also comes with a root account
  - We are discouraged from using the root account 
  - Because its a highly privileged administrative acount
  - that can change anything about the operating system
  - This can lead to mistakes, disastrous outcomes or have catastrophic effects
  - Firstly lets take a look at how to setup a user account and manage users

  - [pc]
  - *walkthrough (SUDO & Administrative tasks) section *


[seg-5]  - [desk]
  - This brings us to the file system 
  - Linux has an intimidating and overwhelming filesystem
  - This is because all the files and directories are sort acronyms
  - that we may not understand the meaning of
  - this is ok.
  - i'd highly recommend to ignore these when you start
  - and you'll slowly get introduced to directories you need to care about
  - otherwise you will be overwhelmed with information
  - and to be honest I still dont know what most of the directories in linux are used for 
  - I've been using linux for years now and some directories 
  - i got really involved in, and then never touched them for years, and forgot what they are used for 
  - The key with Linux is just stay in your lane
  - You dont have to know everything, and you'll be fine
  - Let's focus on learning how to navigate the file system

  - [pc]
  - cover (File system & Navigation section)

[seg-6]  - [desk]
  - This may all seem like a lot of information
  - Which is ok, The key with Linux is practise
  - And all of these commands need to become muscle memory
  - So using the command line as often as possible 
  - will make things a lot easier for you in the long run
  - It's all about building that experience and being consistent 
  - with learning
  - Navigating the filesystem is important
  - But you will very often work with files
  - So we need to learn about working with files in Linux
  - Sometimes you need to make a change or few a config file on a production server on Linux, or you may need to reconfigure something.
  - Let's take a look at that 

  - [pc]
  - *walkthrough (File management) section* 

[seg-7]  - [desk]
  - To install things on Linux, we need to learn about package managers
  - Every Linux distrobution has a package manager
  - used to install packages
  - in our first modules in this course, we installed GIT using a package
  - we installed VSCode using a package.
  - There are different type of package manager
  - You have RPM, YUM, APK
  - For Ubuntu Linux, and Debian based Linux you have apt or apt-get

<span style="color:yellow"> 
when i used linux, i never deep dived package management too much
when ever i needed to install something, i googled it.
and then ran the commands i found to install anything i needed.
over time i learned what the commands mean
but i never really understood the concept of apt sources and package sources in the beginning. So don't be overwhelmed by it
</span> 

  - [pc]
  - *walkthrough (Package managers) section* 

<span style="color:yellow">
</span>

 [outtro]

- Hope this video helps you with xxxxxxx.
- In the next one we'll take a look at ...
- Remember to like & subscribe, hit the bell
- You can also join the community
- If you want to support the channel futher, hit the join button - become a member.
- Thanks for watching , until next time , peace!  


https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04
https://www.digitalocean.com/community/tutorials/an-introduction-to-linux-basics

