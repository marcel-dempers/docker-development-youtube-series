# 🎬 Linux Memory Monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../../chapters/chapter-3-linux-monitoring/README.md)

This module draws from my extensive experience in server management, performance monitoring, and issue diagnosis. Unlike typical Linux Memory monitoring guides, which delve into the technical history and intricate workings of memory, this guide is rooted in practical, hands-on experience. It focuses on my approach to understanding memory functionality and usage, the key aspects I prioritize when monitoring Memory usage and exhaustion, and the strategies I employ to address related challenges.

## Memory usage - Understand how memory is used

Memory is an important resource for the health and performance of systems and applications. </br>
As a DevOps, SRE or platform engineer, we need to form a basic understanding of how memory is used by applications that developers write. </br>

Since developers will write applications that use data, this data would ultimately be stored and processed which will consume memory on the server. </br>
Software developers may forget that memory is not an unlimited resource and may also not know what the memory limits on servers are. </br>

In my experience most applications use anywhere between roughly `50mb` to `2gb` of memory depending on the role of the applications.

The size of the memory is generally dependent on what data the application uses. </br>
In a future chapter, we will cover HTTP & Web servers. </br>
Some applications that developers write, will be hosted by web servers and these applications will often accept HTTP requests and accept some data. 

These applications may also rely on external systems like databases and other microservices to retrieve data like inventory records etc. </br>

The size of this data can cause memory usage to increase. </br>
Think about it. If an application accepts an HTTP request and loads a customer record from a database, let's say that record is `500kb` in size. If we deal with a large number of HTTP requests per second, this data could add up quickly.
That means memory usage could be pretty high for this one application. </br>

As requests by this application increase or decrease, memory usage can either increase or decrease over time. </br>

If this application runs on a server that is shared by other applications, they would contend for the memory resource. </br>

<b>General Memory usage: </b>
* Memory is used when an application starts as it is loaded into memory from disk
* Memory usage is dependant on the type of application and how it uses data
  * Database application may hold data in memory (caching)
  * Microservice may load data over the network (from another service or database)
  * Memory used in an ongoing manner by application and needs cleanup
  * Job application that processes a large amount of data 

<b>Things we have to consider: </b>
* Is the memory usage of each application justified and normal ?
* What does the memory usage of this application look like over time ?
* How much memory is available on the server where these applications run ?
* Do we have enough capacity on our server for future requests ?

<i>Side note:
Now this does not always mean that we need to prevent applications from sharing the same server.
Sharing resources mean we save costs, and reduce the number of servers we have to manage. There are always management vs cost trade-offs to consider </i>

### Memory leaks and garbage collection

As I've explained above, memory will be consumed by applications that developers write. </br>
As we monitor the ongoing memory usage of these applications, we would expect the memory usage to be justified by the logic the application is performing, as well as be consistent over time. </br>

Application framework and web servers generally rely on a process called [garbage collection](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) </br>
Garbage collection is the automatic process that releases memory that was used by the application which is no longer needed. </br>

Application frameworks often attempt to be smart by holding on to memory that it may need. </br>
In the previously mentioned example of an application that deals with customer records, if the ongoing memory usage is `1gb`, an application framework may attempt to hold onto this allocation to improve performance and may release memory that is no longer needed when demands drop. </br>

Garbage collector may run as a separate thread in the application that runs periodically and releases memory back to the operating system. </br>
There are more trade-offs here. </br>
The more frequent the GC runs, the more CPU it may consume to iterate memory objects, but the less frequent it runs, the longer it will hold onto memory. </br>
The larger the applications memory demand, the more CPU the GC may consume to iterate over large quantities of memory. </br>

If this process is not working as expected, or there is a bug in the application code preventing release of unused memory, we would call that a memory leak </br>

We detect this by monitoring memory over time for an application to see it rise above the justified norm and not reducing when demand for the application decreases </br>

We will take a look at memory monitoring tools and how we can detect memory leaks in this module </br>

## Out of Memory (OOM) 

When a Linux system runs critically low on memory, the kernel's Out of Memory (OOM) killer mechanism is triggered to free up memory by terminating one or more processes. This is a last-resort measure to prevent the system from crashing due to memory exhaustion. </br>

The OOM killer selects processes to terminate based on several factors, including:
* `OOM Score`: Each process has an OOM score, which is calculated based on its memory usage and other criteria. Processes with higher OOM scores are more likely to be killed.
* `OOM Adjust`: Processes can have their OOM score adjusted using certain parameters. This allows certain processes to be protected or made more likely to be killed.

It's important to know that if you see "oom" or "oom killed" or a combination of these terms in logs, you may be facing an "Out of Memory" issue. </br> Logs may not indicate an exact message that states "Out of Memory", but may be abbreviated as "oom" or "oomkilled" etc. </br>

So as engineers we have to be on the lookout for these terms </br>

I have often troubleshooted "Out of memory" issues and look for clues in the Kernel ring buffer logs. </br>

The `dmesg` command prints the kernel ring buffer messages, which include OOM kill events.

You can use `grep` to find "oom" messages </br>
This command filters the dmesg output to show only lines containing "oom", which will include OOM kill events :

```
dmesg | grep -i 'oom'
```

## SWAP Memory

Memory (RAM) is physical hardware in a computer, just like the CPU is a physical processing unit. </br> 
We learned that servers are physical computers, but servers can also be Virtual. </br>
Virtual servers are "software-based" servers. </br>
The operating system is also capable of creating Virtual memory. </br>
Virtual memory is also known as SWAP memory. 

SWAP memory is a portion of the hard drive designated to be used as virtual memory when the physical memory is fully utilized. It acts as an overflow area for the system's memory, allowing the operating system to continue running applications even when the physical memory is exhausted.


## Tools to test monitoring tools

Throughout my journey, I have used many monitoring tools in the field when production servers are having issues. There is one positive thing from learning about monitoring tools as we will do in this module, but theres another positive which is very important, which is to have tooling available to TEST your monitoring tools to help validate your understanding of said monitoring tool. </br>


Sometimes our interpretation of what a monitoring tool tells us can be incorrect or slightly off. 
For example, when using tools that provide insights to memory usage, and you see a process is using 100mb of memory, you can validate the tool by writing a small script that utilises 100mb of memory and validate whether or not you can see that same number in the monitoring tool. Now you know you are using 100mb of memory because you wrote that in the script and you also see the 100mb memory usage in the monitoring tool. </br>
The above is a very simplistic example, but when it comes to monitoring things like networks, it can be extremely helpful to validate your findings by using simulation scripts. </br>

In this guide we have a [simulation script](./.test/memoryleak-memory.sh) to generate memory usage, simulating a memory leak. </br>
<i><b>Important Note:</b> Please note that links to these files can change without notice and may be different from what you read on screen in the video at the time of recording.</i>


## Memory monitoring tools for Linux

The simplest way to check memory usage on a server is to use the native `top` and extended `htop` tools. </br>
[top](https://man7.org/linux/man-pages/man1/top.1.html) is a native linux monitoring tool that provides an overview of the current system load. </br>
This will give us overall Memory usage for processes as well as other performance metrics.

[htop](https://www.man7.org/linux/man-pages/man1/htop.1.html) is a very popular alternative to `top` as it allows us to be able to scroll horizontally and vertically as well as having other features such as searching and filtering processes and more. </br>

To install `htop` on Ubuntu Linux
```
sudo apt-get install -y htop
```

### free
[free](https://man7.org/linux/man-pages/man1/free.1.html) is a Linux tool that displays the amount of free and used memory in the system. The `free` command in Linux provides a summary of the system's memory usage, including total, used, free, shared, buffer/cache, and available memory </br>

#### Understanding the output

* `total`: The total amount of physical RAM in the system.
* `used`: The amount of RAM currently used by processes and the operating system.
* `free`: The amount of RAM that is completely unused.
* `shared`: The amount of memory used by the tmpfs filesystem, typically used for shared memory.
* `buff`/cache: The amount of memory used for buffers and cache. This memory is available for use by processes if needed.
* `available`: An estimate of the amount of memory available for starting new applications, without swapping. This value is calculated by the kernel and is more accurate than the free column for determining how much memory is truly available.

The above tools are great when you need to jump into a server to see what's going on in relation to system load, cpu and memory usage. </br>

Linux stores memory statistics in `/proc/meminfo`

Check it out:
```
cat /proc/meminfo
```

### sysstat tools 

[sysstat](https://github.com/sysstat/sysstat) is a collection of performance tools for the Linux operating system. </br>

The `sysstat` package provides us with many performance monitoring tools such as `iostat`, `mpstat`, `pidstat` and more. </br>
Some of these provide insights to memory usage by the system as well as individual breakdowns of memory usage by applications.  </br>

<b>Important Note:</b> <br/> 
`sysstat` also contains tools which you can schedule to collect and historize performance and activity data. </br>
We'll learn throughout this chapter that Linux writes performance data to file, but only writes the current statistics to file. So in order to monitor statistics over time, we need to collect this data from these files and collect it over a period of time we'd like to monitor. </br>

To install sysstat on Ubuntu Linux: 
```
sudo apt-get install -y sysstat
```

### pidstat

`pidstat` provides detailed statistics on memory usage for individual processes. This is useful for identifying which processes are consuming memory and reporting on memory usage for a specific process over time . </br>

According to documentation: 

```
pidstat - Report statistics for Linux tasks.
The pidstat command is used for monitoring individual tasks currently being managed by the Linux kernel.
```

We can run `pidstat --help` to see high level usage

```
pidstat --help
Usage: pidstat [ options ] [ <interval> [ <count> ] ] [ -e <program> <args> ]
```

We can have `pidstat` also monitor a given program using the `-e` option as shown above. `-e` allows us to:

```
Execute program with given arguments and monitor it with pidstat.  pidstat stops when the program terminates.
```

Examples:

The `-r` option allows to `Report page faults and memory utilization.` specifically
```
pidstat -r 1 5
```

Understanding the output of `pidstat` 

* `Time`: The timestamp of the sample
* `UID`: The user ID of the process owner, with 0 being `root` and non 0 being a non `root` user 
* `minflt/s`: The number of minor faults per second.
  
Minor faults occur when a process accesses a memory page that is not in its working set but is in memory. These faults do not require loading the page from disk.
* `majflt/s`: The number of major faults per second.

Major faults occur when a process accesses a memory page that is not in memory, requiring the page to be loaded from disk. In this output, all major faults are 0.00, indicating no major page faults occurred during the sampling period.

* `VSZ`: Virtual memory size in kilobytes.

The total amount of virtual memory used by the process, including all code, data, and shared libraries plus pages that have been swapped out.
* `RSS`: Resident Set Size in kilobytes.

The portion of the process's memory that is held in RAM. This includes all code, data, and shared libraries that are currently in physical memory.
* `%MEM`: The percentage of physical memory used by the process.

The proportion of the total physical memory that the process's RSS represents.
* `Command`: The command name of the process.


## sar - report statistics over time

The sysstat tools covered so far all report statistics that Linux provides in realtime and have no historical data. In order to 
get historical data over time, we would need to enable the `sysstat` service and it will collect statistics in the `/var/log/sysstat` folder. </br>

To enable and start `sysstat`, we run the following commands: 

```
sudo systemctl enable sysstat
sudo systemctl start sysstat
```

Now the `sysstat` service will start collecting metrics and store it so we can use another command called `sar` to view historical data. 
View the available logs by running `sar` against the files. Each file represents the day of the month

```
ls /var/log/sysstat/
sa07

sar -r -f /var/log/sysstat/sa07
```

This is good to understand, but keep in mind that in modern distributed systems, we generally export this statistical data to external storage.

## Where to from here 

Remember that I always go back to talking about how Technology is a giant onion with layers. </br>
For the topic we discussed today there are several layers from 1) being the deepest

1) Linux Kernel writes process statistics in the `/proc` folder
2) top, htop, sysstat and other tools read and parse this data and displays it
  * Historical data is not stored, this is up to engineers to use as they please
  * We can keep historical data with the `sysstat` service and use `sar` to report on it
3) Tools like Prometheus, DataDog, NewRelic has "collectors" or "agents" that read that data so it can be sent to a central storage
  * We do not want to store data on our servers 
  * Send data off to highly available storage
4) Prometheus has a time-series database that stores large amount of metrics
5) Visualization tools like Grafana helps us visualise the data in Prometheus by building queries and dashboards