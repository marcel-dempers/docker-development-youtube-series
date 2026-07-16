# 🎬 Linux Disk Monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../../chapters/chapter-3-linux-monitoring/README.md)

This module draws from my extensive experience in server management, performance monitoring, and issue diagnosis. Unlike typical Linux Disk usage and monitoring guides, which dive into the technical history and architectures of different types of disk storage, this guide is rooted in practical, hands-on experience.
As DevOps engineers, we don't necessarily  need to understand exactly how disk writes and reads are performed and how disk sectors are structured and designed. </br>
Much of these low level details that I have learned in University has not helped me at all in today's modern DevOps and Cloud Architecture roles. </br>

This module focuses on my approach to understanding disk functionality and usage, the key aspects I prioritize when monitoring disk usage and performance bottlenecks, and the strategies I employ to address related challenges.


## Disk usage - Understand how disk is used

Disk is an important resource for the health and performance of systems and applications. </br>
As a DevOps, SRE or platform engineer, we need to form a basic understanding of how disk is utilised by applications that developers write as well as by the operating system. </br>

As we have learned in our introduction to monitoring, disk is the only long term persistent storage for a computer system </br>
When a server shuts down, all important data, files, applications and configuration files are all retained on disk. </br>

![diagram](disk.drawio.svg)

### Applications are just files

Applications that developers write are compiled into binaries or executables which are all but files on disk. </br>
Those files are then combined alongside configuration files and are "packaged" - which could just be a `zip` file. </br>
The package is then deployed (which is a fancy word for "copy") to a server (just like our server), extracted into a directory and the application can then be started. </br>

### Linux is a bunch of files 

In Linux, everything is a file. Files are stored on disk </br>
We've seen this by looking at CPU and Memory statistics and we also learned that Linux writes all the statistics into `/proc` where other tools can read the statistics from </br>
Therefore, even real time statistical data is kept in files by the operating system. </br>
In our chapter on operating systems and servers, users, permissions, the operating system, all the statistics are all just files. </br>

Reason I am stating this is because files take up space. </br>
This makes disk space a critical resource for the operating system and the applications running on it, to remain stable and healthy. </br>

### Disk space

In Windows, we have C:\ drive as well as any other disk drives added. </br>
In Linux, instead of viewing a single disk like C:\ , we view File systems instead. </br>
When we installed Linux in chapter 2, we had to select partitions and the Linux installation process would divide our single disk drive into multiple partitions which we now see as separate file systems. 
This includes things like the boot partition, 

Linux file systems can run out of space, we need to monitor overall disk space
* Avoid storing data on the operating system disk. Instead use attached disk or other volumes or storage

When space in Linux runs out, many things come to a halt. </br>
In my experience, catastrophic things can happen with the operating system disk space runs out. </br>
Unlike Windows, Linux does not pop up with prompts and alerts stating that the available disk space is low. </br>
We need to have monitoring in place to detect this before it's too late. </br>

### Disk Reads & Writes (I/O)

Disk space is not the only resource to monitor. a Process may read a file to read the content or it may write to the file if it needs to write content. </br> These read and write operations are technically a resource too. The Disk has a certain bandwidth of reads and writes per second that it can handle. </br>

This metric is generally referred to as Input/Output speeds or disk I/O. </br>

When deploying virtual servers in the cloud, the cloud provider will often publish the maximum disk speeds. </br> This is important during cost analysis and choosing the right server for your needs. </br>
Some applications need to write to disk and may require better performing disks. Most web applications may not need fast disk. </br>

This is where it becomes important for engineers to understand the differences between stateful and stateless applications. </br>

### Stateful vs Stateless applications

When talking about disk, it's important to understand that applications or processes in software engineering can have state. </br>
Many applications may simply run, receive some data, process it and respond with output. </br>
When these types of applications are terminated, they technically do not lose any data, because they rely purely on inputs and outputs. </br>

A Web server is a prime example of this. Requests come in from a customer's browser to load a page and the web server simply responds with the content. If the web server is restarted, shut down, or terminated unexpectedly, it can recover and continue to serve content. The web server process does not have any internal state. </br>

However, let's say a customer is logged in on our website and their session data is stored on our web server, this makes our web server stateful. It also means all web requests for that logged in customer needs to be routed to that same web server where the session data is. </br>
If our web server shuts down or terminates unexpectedly, the customer may need to be routed to another healthy web server. However their logged in state is on our faulty server which has terminated. 

A fix for the above scenario is that a web server should not store session state and this state can be handled by something like a database. So our web servers can be restarted and the customers all remain logged in. 

This means web applications should strive to be stateless. </br>

An example of a stateful application is something like a database, media storage, caching server or a content system. </br>
Stateful applications have to strive to ensure when they are terminated, that their data is always written to disk and potentially even replicated across a few servers to ensure data integrity. </br>

#### Why is this important for DevOps SRE and Platform engineers ?

It's very important when it comes to disk monitoring, to understand what will be running on our servers in a production environment. </br>
When largely managing stateless applications, you should not expect high disk usage or high number or read and write operations. These types of applications may need some disk access to load any libraries, files or configuration needed to run, but should not be dependent on storing state. </br>

When managing stateful applications, we need to be aware and expect higher disk consumption and utilization.

Things to keep in mind:
* Stateful vs Stateless applications and impacts to disk usage
  * Stateless applications are less reliant on disk space and I/O
  * Stateful applications are more dependant on disk space and I/O
* Build servers (CI/CD) generally use higher disk space 
  * Build servers download large codebases, which will consume disk space
  * Build servers compile code which may involve a lot of disk read\write operations
* Antivirus and security software scan files on disk, resulting in high disk read\write operations 

## Disk monitoring tools for Linux

For disk monitoring tools, we should be able to monitor disk space for the overall system as well as be able to investigate directory sizes and where disk space consumption is coming from. </br>

In addition to disk space, we also need to monitor disk I/O usage to identify any bottlenecks when processes are intensively writing to disk. </br>

In this guide we have a [simulation script](./.test/app_logger_disk.sh) to generate disk usage and high I/O. </br>
<i><b>Important Note:</b> Please note that links to these files can change without notice and may be different from what you read on screen in the video at the time of recording.</i>

### df

[df](https://man7.org/linux/man-pages/man1/df.1.html) is a native linux monitoring tool that allows us to review disk space of each file system. </br>

The `df` tool allows us to view the file systems in our operating system. </br>
It's a great start to see overall used and available space for our server </br>
The `df` command takes a `-h` flag that tells `df` to print the output in human readable form.

### du

[du](https://man7.org/linux/man-pages/man1/du.1.html) is a native linux monitoring tool that allows us to review files and directory sizes. </br>

It allows us to summarize disk usage of the set of files, recursively for directories.

`df` also supports a `-h` flag to output in human readable format. </br> 
In my experience, `du` is great at finding large directories and files manually whilst navigating the file system.

`du` and `df` are great because they are available in most Linux distributions. 

These commands may help in production environments where there may not be any tools installed, so I think it's important to practise on how to use these two tools efficiently in combination</br>

You can run `du` on different depth levels which tells the command to how far to traverse within directories. </br> This helps us visualise directories that might have high disk usage and then focus on those directories to locate large files or large sets of files. <br>

### tree

[tree](https://linux.die.net/man/1/tree) is a pretty cool utility in Linux that provides us similar capabilities as `du` but with a more human-friendly output. It allows us to list contents of directories in a tree-like format. </br>

It can print sizes of files and folders as well and we can also control the depth at which to list directories. </br>

There are a large number of tools when it comes to Linux and monitoring. </br>
In this course, I generally focus on the main most popular ones, but also keep in mind that in modern distributed systems, you may use an external monitoring system and may not log into production servers directly. </br>

### sysstat tools 

[sysstat](https://github.com/sysstat/sysstat) is a collection of performance tools for the Linux operating system. </br>

The `sysstat` package provides us with may performance monitoring tools such as `iostat`, `mpstat`, `pidstat` and more. </br>
Many of these providing insights to disk utilization </br>

<b>Important Note:</b> <br/> 
`sysstat` also contains tools which you can schedule to collect and historize performance and activity data. </br>
We'll learn throughout this chapter that Linux writes performance data to file, but only has the current statistics written to file. So in order to monitor statistics over time, we need to collect this data from these files and collect it over a period of time we'd like to monitor. </br>

To install sysstat on Ubuntu Linux: 
```
sudo apt-get install -y sysstat
```

### iostat

According to the documentation:

```
iostat -
Report Central Processing Unit (CPU) statistics and input/output statistics for devices and partitions.
```

We can run `iostat --help` to see high level usage

```
Usage: iostat [ options ] [ <interval> [ <count> ] ]
```

`iostat` Gives us average disk utilization and IO since system startup </br>
This means that if the system has only recently become busy, the average utilization may appear low because it is averaged over the entire uptime of the system.  </br>

To get the current stats, we can provide an interval and we can also provide a count of snapshots we want of the stats. For example, we can get stats for every second, with a total of 5 snapshots. 

We provide the `-d` option, which states that we are interested in device utilization stats. We also use `-x` which indicates we would like extended statistics </br>

```
iostat -dx 1 5
```

Let's understand the output for the I/O report:

* `Device`: The name of the device (e.g., loop0, sda).
* `r/s`: The number of read requests per second.
* `w/s`: The number of write requests per second.
* `rkB/s`: The number of kilobytes read per second.
* `wkB/s`: The number of kilobytes written per second.
* `rrqm/s`: The number of read requests merged per second.
* `wrqm`/s`: The number of write requests merged per second.
* `r_await`: The average time (in milliseconds) for read requests to be served.
* `w_await`: The average time (in milliseconds) for write requests to be served.
* `svctm`: The average service time (in milliseconds) for I/O requests.
* `%util`: The percentage of CPU time during which I/O requests were issued to the device (bandwidth utilization for the device).

### pidstat

`pidstat` provides detailed statistics for individual processes. This is useful for identifying which processes are consuming the most resources. </br>

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

The `-d` option allows to `Report I/O statistics` specifically
```
pidstat -d 1 5
```

## Troubleshooting Disk usage

When it comes to troubleshooting disk usage, we are either looking at disk space monitoring to understand disk consumption or disk I/O monitoring to troubleshoot bottlenecks when applications are performing high reads or write operations on disk. </br> 

I would always start with `df` and then `du` to dive deeper into working out which file system is affected and where space consumption is taking place. </br> So we try to locate the folder that takes up all the space. </br>

Since we learned that Linux is just a bunch of files, that is a good thing, because these files are all documented. </br>
This means when you see space is consumed by a directory that you don't know much about, you can simply search that directory online and you would be able to locate a root cause fairly quickly in most cases. </br>
We can then also quickly find out if it's safe to remove those directories or not. 
And also work out if we can prevent that space from being consumed in the future. </br>

Other than disk space - Another situation is working out if there are any disk bottlenecks. </br>
Generally speaking web services and applications don't do much disk I/O, so it's quite rare in my experience to come across disk trouble. </br>
One place where you may find it come up as a problem is in the cloud, especially in Build servers. 
Build servers are servers that are dedicated to compile applications, run tests and package applications into artifacts that can be deployed. </br>
We'll cover build and deployments in our CI/CD chapter. But it's important to know that build operations involve high disk interactions. </br>
This means you'll want an SSD type disk which is fast for build pipelines. </br>
The cloud may give you certain virtual servers that you can choose from, but keep in mind, that disk is important so you will need to choose a server with an SSD type disk and not a general HDD. Otherwise your build pipelines may run a lot slower than expected. </br>

For disk I/O, we use `iostat` as mentioned above to work out if the overall server is having any I/O trouble. </br>
We can then also use `pidstat` to workout which process is the culprit and then see where that process is coming from by running the following and replacing `<pid>` with the process ID that we discovered in `pidstat`

```
ps -p <pid> -o pid,cmd
```