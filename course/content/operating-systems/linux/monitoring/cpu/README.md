# 🎬 Linux CPU Monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../../chapters/chapter-3-linux-monitoring/README.md)

This module draws from my extensive experience in server management, performance monitoring, and issue diagnosis. Unlike typical Linux CPU monitoring guides, which delve into the technical history and intricate workings of CPUs, this guide is rooted in practical, hands-on experience. It focuses on my approach to CPU functionality, the key aspects I prioritize when monitoring CPU performance, and the strategies I employ to address related challenges.

## CPU usage - How the CPU works

In our chapter on Operating Systems, we covered the basics of system resources that the operating system manages, including the CPU, which is the Central Processing Unit.</br>
The CPU performs all processing operations of a computer. 

Applications and software consist of code that translates into 'tasks' for the CPU to process when the application runs. <br/>

### CPU Single task execution

The CPU is composed of multiple cores, and each core can process one task at a time. Because CPU processing is extremely fast, it appears as if tasks are executed simultaneously and in parallel. However, each CPU core can only execute one task at a time.

So in simple terms, a CPU core will take a task and perform one or more "CPU Cycles" to process that task. </br>

To revisit Chapter 2 on CPU - A CPU with a clock speed of 3.2 GHz executes 3.2 billion cycles per second.
The cycle of a CPU involves: 

1) Fetching the instruction to execute
2) Decode the instruction
3) Read the address from memory
4) Execute the instruction

In Chapter 2, I suggested thinking of a CPU Core as a spinning wheel, each "wheelspin" is a CPU cycle that processes 1 task. </br>
When the wheel is spinning, it cannot process another task and has to finish it's cycle so another task can hop on. </br> 

This means that tasks may queue to wait their turn to get onto the CPU. </br>

![cpu-single](cpu-single.drawio.svg)

### CPU Multi Task execution

Because a single CPU core is extremely fast as mentioned above, it may process tasks extremely fast and appear as if it's completing them all at once. </br>

However, we now know that a single CPU core can only perform one task at a time. </br>
To speed this up, the CPU has multiple cores for true parallel processing. </br>

So CPU is a shared resource, similar to memory and disk and other resources. </br>
If an application needs disk space, it gets disk space allocation. </br>
If an application needs memory, it gets memory allocated. </br>
However, CPU is a much more intensively shared resource as it does not get allocated to an application . <br>
Applications all queue up their tasks and CPU executes it </br>

This makes it difficult for operating systems to display exactly how much CPU each application
is using, but it does a pretty good job in doing so. </br>

![cpu-single](cpu-multi.drawio.svg)

### Understanding CPU as a Shared Resource

The CPU, like memory and other resources, is a critical resource that is shared among all running applications on a system.

The CPU is not allocated to applications in fixed chunks. Instead, the CPU time is shared among all running applications through a process called scheduling.

The operating system's scheduler rapidly switches the CPU's focus between different applications, giving each one a small time slice to execute it's tasks. This switching happens so quickly that it appears as though applications are running simultaneously.

Because the CPU is a highly contended resource, the operating system must manage this sharing efficiently to ensure that all applications get a fair amount of CPU time and that high-priority tasks are executed promptly.

## Why is this important for DevOps, Platform, Cloud & SRE Engineers ?

Understanding how the CPU is shared and utilized is crucial for several reasons:

* Reporting CPU usage 

When we monitor a system's CPU usage using some tool, we need to understand what the values mean. </br>
If you work for an organization, they would likely use external 3rd party systems like [New Relic](https://newrelic.com/) or [DataDog](https://www.datadoghq.com/) to monitor servers and applications. </br>

As engineers we'll want to understand :
* How these metrics are generated 
* Where these metrics come from
* How they are measured and collected
* What it all means for performance and health of the system.

The operating system provides metrics on CPU usage, but interpreting these metrics requires an understanding of how CPU time is shared among processes.
For example, a reported CPU usage of 65% indicates that 65% of the CPU's time is being used to execute tasks, but it doesn't specify which applications are using that time or how it is distributed among them.

Overall CPU as a `%` gives us an indication how busy a server is, but does not always tell the full story. </br>
And that's what we'll be learning here today. </br>

In addition to the above, we may need to understand the `%` of CPU usage <b>on all cores</b> of CPU. </br>

###  Understanding Single vs Multithreaded applications </br>

Since we know that a CPU core can only execute one task at a time, it's important for engineers to know how to take advantage of this and also avoid the pitfalls.  

Example 1:  If you have a task which may contain poorly written code, it could keep CPU cores busy unnecessarily, causing other tasks to queue up for execution. This can slow the entire system down. </br>

Example 2: If you have code that may be poorly written, you could end up in situations where you are only utilizing 1 core during task execution. This is common when engineers write loops in their programs. </br>
This means your application is not running optimally and not utilizing all available CPU cores. </br>

Example 3: Another example of poorly written code where one task is waiting on another task, may end up in what's called a CPU "Deadlock". This occurs when all executing tasks are waiting on each other in a circular reference. </br>

![](cpu-cpu-threads.drawio.svg)

### Worker threads and tasks

To solve some of the above issues, programming, scripting and runtime frameworks allows us to write our code in such a way that we can create what's called "threads", "worker threads", or "tasks". </br>

Web servers are a good example of this. When an HTTP request comes to a web server, the web server can create a "worker thread" to handle that request. Then that request gets executed on a CPU core, the web server may issue another "worker thread" to handle other incoming requests. </br>

This way, a web server can handle multiple incoming requests and makes use of all available CPU cores on the system </br>

We may call the HTTP request handler of the web server a multithreaded application. </br>
The web server code for handling each HTTP request may be viewed as "single threaded" code </br>

We can take a look at two bash scripts that I have written that demonstrates single vs multithreaded code </br>

#### example: single threaded code

If we execute [./singlethread-cpu.sh](.test/singlethread-cpu.sh), we will notice that only one core of the CPU is busy at a time. </br>
Now because we execute a loop, each iteration of that loop will run on a different core. </br>
Bash itself is single threaded, so this script needs to be optimized if we want to make use of all available CPU cores. </br>

#### example: multi threaded code 

If we execute [./multithread-cpu.sh](.test/multithread-cpu.sh), we will notice all CPU cores get busy. </br>
This is because in this script, we read the number of available cores with the `nproc` command. </br>
Then we loop the number of available cores and execute our `simulate_cpu_usage()` function. </br>
At this point it would technically still be single threaded, because it's still a loop and does not create more threads or processes. </br> To get around this we use a special character in bash called `&` at the end of our function. </br>

In Bash, the `&` character is used to run a command or a script in the background. When you append `&` to the end of a command, it tells the shell to execute the command asynchronously, allowing the shell to continue processing subsequent commands without waiting for the background command to complete. </br>

So if we have 8 CPU cores, our script will spawn 8 instances of our function, and each will run on a different CPU core. </br>

This means our application is a multithreaded application! </br>

### CPU Busy vs. System Busy: Understanding the Difference

Common scenarios where CPU usage interpretations may lead to misdiagnosis include:

When overall CPU usage is low, applications or systems can still be slow if there are other bottlenecks </br>

<u>For example:</u>

* Poorly written code that causes single-threaded execution can result in low CPU usage and slow application performance, which may be difficult to detect when examining overall CPU metrics.

* CPU usage may be low, but an application can still slow down if it is experiencing network or disk I/O bottlenecks.

It's important to be aware that having overall low CPU usage, does not always mean your systems are performing optimally, and there could be other bottlenecks and symptoms of slowness. </br>

## How is CPU usage measured 

From the above explanation, we now have a high-level understanding of CPU usage. Measuring CPU usage can be challenging because many processes alternate between active execution and waiting to be executed. </br>

Most operating systems display CPU usage as a percentage, indicating the proportion of CPU resources used across all cores. Calculating this percentage is complex, and operating systems strive to provide the most accurate measurement possible. </br>

The percentage display mostly makes it easier for humans to interpret if CPU usage is high or not. </br>

We've also learned that CPU reporting is generally based on time spent "on-cpu". So if the Operating System reports `65%` CPU usage, it means 65% of the CPU's time is being used to execute tasks. </br>

Most monitoring tools we'll be looking at will show us `%` CPU. </br>
Some advanced monitoring tools may show us CPU by process in `seconds` instead. </br>
So that tells us how much CPU time the process had and we can calculate the `%` ourselves. </br>

## CPU monitoring tools for Linux 

### top
[top](https://man7.org/linux/man-pages/man1/top.1.html) is a native linux monitoring tool that provides an overview of the current system load. </br>
This will give us overall CPU usage for processes as well as other performance metrics

### htop

[htop](https://www.man7.org/linux/man-pages/man1/htop.1.html) is a very popular alternative to `top` as it allows us to be able to scroll horizontally and vertically as well as having other features such as searching and filtering processes and more. </br>

### sysstat tools 

[sysstat](https://github.com/sysstat/sysstat) is a collection of performance tools for the Linux operating system. </br>

The `sysstat` package provides us with many performance monitoring tools such as `iostat`, `mpstat`, `pidstat` and more. </br>
All of these providing insights to CPU usage and load on the system </br>

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

`iostat` Gives us average CPU load and IO since system startup </br>
This means that if the system is only recently busy, the average CPU usage may be very low because it's an average from the time the system started up </br>


To get the current stats, we can provide an interval and we can also provide a count of snapshots we want of the stats. For example, we can get stats for every second, with a total of 5 snapshots. 

We can also provide the `-c` option, which states that we are only interested in CPU stats and not I/O stats. </br>
```
iostat -c 1 5
```

Let's understand the output. This output is pretty common across `sysstat` tools

* `%user` 
  * Percentage of CPU utilization that occurred while executing at the user level (application).
* `%nice` 
  * Percentage of CPU utilization that occurred while executing at the user level with nice priority.

The Niceness value is a number assigned to processes in Linux. It helps the kernel decide how much CPU time each process gets by determining it's priority.

* `%system`
  * Percentage of CPU utilization that occurred while executing at the system level (kernel).
* `%iowait`
  * Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.
* `%steal`
  * Show the percentage of time spent in involuntary wait by the virtual CPU or CPUs while the hypervisor was servicing another virtual processor.
* `%idle` 
  * Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.

### mpstat

`mpstat` provides detailed statistics on CPU usage for each individual core. This is particularly useful for identifying imbalances in CPU usage across cores, which can help in optimizing application performance. </br>

According to documentation: 

```
mpstat - Report processors related statistics.
The  mpstat  command  writes  to  standard output activities for each available processor, processor 0 being the first one.  Global average activities among all processors are also reported.
```

We can run `mpstat --help` to see high level usage

```
mpstat --help
Usage: mpstat [ options ] [ <interval> [ <count> ] ]
```

Examples:

```
mpstat -P ALL 1 5
```

### pidstat

`pidstat` provides detailed statistics on CPU usage for individual processes. This is useful for identifying which processes are consuming the most CPU resources. </br>

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

The `-u` option allows to `Report CPU utilization` specifically
```
pidstat -u 1 5
```

As per our discussion on single vs multithreaded applications, `pidstat` allows us to report on thread-level activity too using the `-t` option. 

```
pidstat -t 1 5
```

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

sar -u -f /var/log/sysstat/sa07
```

This is good to understand, but keep in mind that in modern distributed systems, we generally export this statistical data to external storage

## Troubleshooting High CPU usage

The first challenge to resolve high CPU usage is to locate the overall cause of CPU usage, by using the above tools such as `top`, `htop` or the `sysstat` tools </br>

We can use filtering, searching, and sorting techniques to find and display the culprit process. </br>
All of the above tools are great at identifying the culprid process and what you need is the process ID or also known as "PID" number. </br>

The `PID` is the process identifier of a process running on the operating system. 

Once we have the process we can see the command that is causing it. </br>
Most of the time, you can identify the culprit by looking at the `COMMAND` column of `top` or `htop`. </br>

If it's an executable, a script or program running somewhere we can locate it using the `ps` command. </br>
This command line helps us display information about a selection of the active processes </br>
You can check out `man ps` to get more details about all the available options. We are after the process using the `-p` option to pass the process ID of the culprit. This process ID can change, so if you follow along be aware of that. 

```
ps -p 6268 -o pid,cmd
PID CMD
6268 /bin/bash ./singlethread-cpu.sh
```

Above, we use `-p` to pass the process ID, we use `-o` to display output about the process and it's command and that helps us locate the executable, in our case a bash script called `singlethread-cpu.sh` which is executed by bash, under the `/bin` folder. </br>

Now we learned in our command line episode, that the `./` allows us to execute scripts and executables in the current working directory. </br>
We cannot see what the current directory is, so if we need to locate that culprit, we can try to find it using the `find` command </br>

`find` takes a file path, where we can specify to search from the root of the filesystem with `/` </br>
We specify `-name` with the name of what we want to search for. </br>

The `find` program will return messages for every match to `stdout` but will also return all the non-matches as error messages to `stderr`.

We can use special redirection to redirect all `stderr` messages to a special file that discards all data written to it. 

This may take a long time to run on systems with large about of files because we will be trawling recursively through the entire filesystem from `/` :
```
find / -name singlethread-cpu.sh 2>/dev/null
```

Explanation:

* `2>` : This specifies that the redirection applies to the standard error stream (file descriptor 2).
* `/dev/null` : This is a special file that discards all data written to it, effectively making the error messages disappear.

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

