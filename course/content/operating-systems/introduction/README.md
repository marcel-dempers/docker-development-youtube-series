# 🎬 Introduction to Operating Systems for DevOps

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 2](../../../chapters/chapter-2-operating-systems/README.md)

## What is an Operating System

An operating system is software that allows us to run programs and interface with the hardware. </br>
Users of a computer, and applications that are installed, do not have to interact with hardware directly. The operating system does this for us so it sits as a layer above the hardware </br>

All our programs, code, scripts, automation and "things" we do in DevOps will run on an Operating System</br>
```
┌──────────────────────────────────┐
│             USERS                │
│                                  │
└───────────────┬──────────────────┘
┌───────────────▼──────────────────┐
│        APPLICATIONS              │
│                                  │
└───────────────┬──────────────────┘
┌───────────────▼──────────────────┐
│      OPERATING SYSTEM            │
│                                  │
└───────────────┬──────────────────┘
┌───────────────▼──────────────────┐
│         HARDWARE                 │
│                                  │
└──────────────────────────────────┘

```
For example, your Laptop, Mac, PC, iPhone or Android device all has an operating system. </br>

## Why is the Operating System important for DevOps

In DevOps, SRE and even cloud engineering, Operating systems will play a fundamentally important role in our day to day work. </br>
Your workstation or laptop, may be Microsoft Windows, or MacOS, or you may have Linux installed </br>
Therefore all the tools you may use will be dependant or compatible with either one for different operating systems </br>

In addition to that, production systems you may interact with will have an Operating system installed </br>
Therefore, scripts, automation and deployment pipelines, open-source tools, infrastructure as code and all the DevOps technologies we will use in this course, needs to cater for operating systems </br>

To support and monitor production systems in real world scenarios, you need to have an understanding of how to use different operating systems </br>

For example, a production web server could be using Linux and its currently malfunctioning because it has no disk space left </br>
We'll need to know how to :
* Access the production server
* Navigate and troubleshoot issues
* Solve the issue to return the server to normal operation
* Add preventative measures and improvement to prevent future outages

For the above, you may need basic understanding of Linux, command line tools like SSH to access the server, `df` and `du` which are tools to troubleshoot disk space and other problem solving skills we'll learn throughout this course </br>

## Operating System Key Functions 

The operating system has many important key functions:

### User Management 

You can have multiple users sign into the same device </br>
For example, you may be the administrator of your home computer. So you setup the printer, the internet access, WIFI etc </br>

Another family member has a different user account and password, when they sign in, they have permission to set their own profile, their own desktop wallpaper and manage their own files. </br>
However they cannot alter the internet and printer settings </br>

This is important in DevOps because we can run systems using unprivileged accounts </br>
We will discuss security throughout this course </br>

### Permission management 

The operating system as shown above, allows us to set different permissions for different users or groups. </br>
Users like Grandma and Dad, may be part of the "Family" group and can manage their own profile, files, browse the web and have their own applications that they use. </br>
Members of the "Family" user group cannot change operating system settings like internet access, printers and network settings etc. </br>
You may be a member of the "Administrators" group, which allows you to manage the operating system. </br>

The above user and permission concepts are important, as we will face these throughout the remainder of this course. </br>

When we build our own applications or use applications for automation and scripting, we will need to think about security and what permission these automation processes have </br>

### Hardware management 

When we use a computer, we don't work directly with the hardware. </br>
Our applications also do not directly communicate with the hardware. </br>
For example, our keyboard and mouse may be plugged in via USB cable into a USB port. </br>
The Operating System interfaces with the hardware and communicates with the USB device, so when you move the mouse, signals are sent to the USB port and the Operating system knows how to interpret these signals and makes the mouse pointer move on the screen </br>

Similarly, you have a screen or monitor plugged into your computer via a display port or HDMI port. Similarly the monitor sends display signals to that port and the Operating System knows how to interpret signals on the display port and allows you to see things on the screen. </br>


#### Device drivers

Operating Systems use device drivers to talk to the hardware. </br>
Device drivers are software that is specialised and specifically coded to be able to interpret and understand the hardware signals and functionality. </br>
In the examples of your mouse, keyboard and computer screens, your operating system has USB drivers as well as display drivers, so it knows how to use a mouse and keyboard as well as a screen. </br>

If the Operating System had to know about every device that exists in the world, it would become too large as it would need trillions of different device driver software installed </br>

To help with this , many devices use generic USB drivers, and technology like USB helps to keep things "generic" </br>
So many devices would prefer to use USB as its generic and most computers would understand it and it helps developers of devices not have to build custom driver software for devices. </br>

### Resource Management

For DevOps, resources is a very important topic, as it will impact almost everything you do in a real world production scenario </br>
The Operating System manages access to resources. For example, it will make disk space available to an application that wants to use it </br>
It also controls security to these resources, like which folders an application can access </br>
Some resources may be reserved for system use </br>

#### CPU

The CPU, or central processing unit is the heart of a computer. </br> It is the brain power that is responsible for all compute operations </br>
Any computation that occurs on a system involves the CPU. </br>
When you run a script, a program, some code, or a popular app that you like to use like the browser, the operating system will initiate the CPU to process it. </br> 

CPU operates in "cycles". A CPU with a clock speed of 3.2 GHz executes 3.2 billion cycles per second.
The cycle of a CPU involves: 

1) Fetching the instruction to execute
2) Decode the instruction
3) Read the address from memory
4) Execute the instruction

Without over complicating it, I'd suggest thinking of a processor as a bicycle.  </br>
The operating system will throw "things to process" on the bicycle and the bicycle will "cycle" and execute these instructions </br>
One bicycle with one wheel can only execute 1 cycle at a time, so if there is an instruction that takes a long time to execute, other instructions may have to wait to get their turn on the bicycle. </br>
This execution time is often referred to CPU time </br>

A processor can have multiple cores which can be seen as multiple wheels on the bicycle. </br> 
This allows the processor to perform multiple executions concurrently </br>

In this course, we will be taking a look at CPU monitoring when we start building applications and how to effectively monitor their CPU usage and overall system performance </br>


#### Memory

There are two high level storages in an Operating system. I would like to think of it as short term memory (RAM) and long term memory (Disk). </br>
Now Memory or RAM (Random access memory) is extremely fast storage non-persistent storage. </br>
Which means it does not persist, so when the computer is turned off, or looses power, the data in memory is lost </br>
It's like turning off your mobile phone suddenly and when you turn it back on, everything you had open is gone. </br>
Memory is usually limited in capacity, you may have 8GB , 16GB, 32GB or 64GB of memory, while your hard disk may be 500GB or even a Terabyte in size </br>

<b>Why would you have such small and volatile non-persistent storage and not just use a large disk ? </b>

This is because memory is extremely fast and operates at a much higher frequency speed than a disk would. </br>
The speed at which data is read and written to memory is typically in nanoseconds or faster, where as with Disk, it's in milliseconds </br>

From a DevOps perspective, it's important to know that whenever a CPU has to run something, it has to be loaded into memory. </br>
For example, if you double click a program to start it, or click a document to open it on your computer, the program and document has to be loaded into memory in order for the CPU to process and run the application. </br>

Therefore memory is the most active and important resource in junction with the CPU processor. </br>

Throughout this course, we will take a look at monitoring memory and what impacts memory has on software in the DevOps world </br>

#### Disk

Disk is the main persistent storage for the Operating system. It allows us to store data long term and the data is stored even when the computer is powered off </br>

You get different type of storage types:
* HDD = Hard disk drives
* SSD = Solid state drives
* Flash drives = USB type disk drives

Disk is storage that applications will use to write configuration and program data in order to persist it in times the application is no longer running. </br>
Think of a Database as an example. </br>
When a database process starts, CPU cycles start executing instructions, and first the database application is loaded into memory from disk, then configuration files are loaded into memory. </br>
The database program starts performing database operations which are CPU cycles, reading instructions from memory and loading database data from disk. </br> The data is loaded from disk and the database is then ready to be accessed. </br>
Every database is different, but hopefully you get the idea of the application lifecycle and why disk is important. </br>

When the database program is shut down unexpectedly, all data in memory is lost so there is a challenge to ensure all important data is committed to disk as often as possible to prevent data loss

## The Kernel

So far we've been talking about what the Operating system does, like device and resource management.
These operations are performed by part of the Operating system called the "Kernel" or core. </br>

The kernel is a piece of software that is part of the Operating system's core and controls everything in the system. </br>
The kernel is the interface that allows applications to access hardware

```
                ┌───────────────┐          
                │               │          
                │ APPLICATIONS  │          
                │               │          
                └──────┬────────┘          
                       │                   
                ┌──────▼────────┐          
                │               │          
                │   KERNEL      │          
                │               │          
                └───┬─────┬─────┴───┐      
                    │     │         │      
                    │     │         │      
                ┌───▼──┐┌─▼────┐ ┌──▼─────┐
                │ CPU  ││MEMORY│ │DEVICES │
                │      ││      │ │        │
                └──────┘└──────┘ └────────┘
```
Applications interface with the kernel through system calls. </br>

## Types of Operating Systems

There are two main operating systems that you will be dealing with as a DevOps, SRE or Platform engineer. </br>
<b>Windows</b> and <b>Linux</b> </br>
It is important to be proficient in both of these, however, experience is the key and the only way to get experience is through hands-on practical use and exposure. </br>
In this course, we will become exposed to both to some degree, but our main focus will be Linux </br>

## Terminal and Command line

Another way for us to interact with the Operating system is through the Terminal </br>
This forms another important part of this course as it will be one of the primary ways we will interact with systems in upcoming chapters </br>

In upcoming modules, we will learn about the command line and how to use a terminal. </br>
This will form the basis and foundation of our automation and scripting skills </br>