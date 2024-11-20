<span style="color:yellow">
</span>

🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦🎦 
  

<span style="color:yellow">
[seg-1]  - [desk]
  - Welcome to chapter `2`, and a new module that coveres the introduction to operating systems
  - In the chapter overview I mentioned that under every system that you will work on, there is an operating system.

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
  - Now in this module we'll be covering mostly theory, but i've hand picked the most important  components of operating systems that I deal with day to day.
  - You can do an entire year or two studying operating system fundamentals at university, so it has a lot of depth, but hopefully with this module you will get an idea of the important parts of the operating system so after completing this course, you can make informed descisions about where to put your time and energy to learn more. 
  - Also a lot of this theory will become important, especially in modules where we learn about how to use command line, scripting, and also I have an entire chapter on monitoring, so how to monitor a system effectively, and to do so you will need an understanding of operating system resources.

</span>


[seg-1]  - [desk] 
  - Welcome to another video & new chapter in the ultimate course for DevOps
  - In the previous modules and chapters we covered GIT, source control and IDEs.
  - We learnt how to work with source code and why its important for devops, SRE and platform engineers
  - This this new chapter we'll be covering operating systems
    - As a DevOps, SRE or platform engineer you will be dealing with operating systems 
      - when doing automation, scripting, monitoring and all the tooling we use
      - interfaces with the operating system
    - It's probably the most boring part of IT
    - because Operating systems are often huge portions of university IT semesters
    - and so much time is spent on out of date irrelavant CPU chipset architectures of the 1980s
    - that a lot of the important things for modern devops is lost
    
    - today will be focussing on everything you need to know regarding operating systems in a modern devops field
    - so without further adeau, lets go  
  
[intro]

[seg-2]  - [desk]
  - first of all , what is an operating system
    - An operating system is software that allows us to run programs
    - All the automation , the scripting and tools we will run, run directly on the operating system
    -[side] It allows user programs to [show desktop pc box] interface with the hardware
  - [motion] 
        - in computer systems, users interact with applications 
        - applications are made up of code and scripts that developers write 
        - these applications interact with the operating system
        - and the operating system fascilitates the interaction with the hardware 
        - the operating system ensures we have access to memory, network, disk and CPU cycles to do processing
  - [desk] 
    - Why is the operating system so important for devops?
    - *read ouf the section to camera*
  - [pc] 
    - now the operating system has some key functions.
    - talk through user and permission section
    
[seg-3]  - [desk]
  - Now When we use a computer, we dont work directly with the hardware.
  - Even when we write automation scripts or code, we dont talk to the hardware directly
  - All this is faciliated through the operating system
  - *read out hardware management about mouse and keyboard explanation* [camera]
  - [side]
    - So how does the operating system achieve this ? we'll it uses whats called device drivers       

[seg-4]  - [pc]
  - *read device drivers section* [can show USB drivers in device management tab]
  
  - [desk]
    - now with the boring stuff out the way, the next important and most relevant part of operating systems for devops is resource management
    - there are three most important resources, cpu , memory and disk
    - for devops this is very important because we will monitor how these resources are used in our systems
    - how to view these resources, how to monitor and how to tune our system to effectively use these resources
    - we'll touch on the basics on this in this module and then deep dive futher throughout this course
  - [pc] 
    - *read CPU section* , combine [pc] and [side] views

<span style="color: yellow"> *when discussing cycles:* why is this important ? when we start looking at CPU performance from a devops perspective, we'll start seeing interesting challenges when analyzing CPU usage as a percentange. When we see CPU is at 100% utlization, what does that really mean ? what could application be doing when its stuck at high CPU usage, what could cause that and how can we solve this<span/>

  - [desk] 
    - *read memory section*  

<span style="color: yellow"> *discuss persistance in more depth* The word persistance is very important in software engineering<span/>

  - [pc] 
    - *read Disk section* , combine [pc] and [side] views

    <span style="color: yellow"> Explain database transactions and tradeoff with performance<span/>

[seg-5]  - [desk]
  - So far we've been talking about what the Operating system does, like device and resource management.
    These operations are performed by part of the Operating system called the "Kernel" or core.  
  - Now applications and code or scripts that we write as engineers dont directly talk to hardware or in fact it does not interface directly with memory, CPU or disk.
  - When we write code to say, "create a file on disk with this content"
  - we send that instruction to the operating system and it will talk to disk for us
  - The operating system piece that handles all of these "system calls" is called the kernel

[seg-6]  - [pc]
  - **walk through kernel section in readme**
  - **walk through types section in readme**
  - **walk through cmd section in readme**
  <span style="color:yellow">walk through why command line is important, I.E running out of disk etc</span>

 [outtro]

- Hope this video helps you with xxxxxxx.
- In the next one we'll take a look at ...
- Remember to like & subscribe, hit the bell
- You can also join the community
- If you want to support the channel futher, hit the join button - become a member.
- Thanks for watching , until next time , peace!  
