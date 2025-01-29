# 🎬 Linux CPU Monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../../chapters/chapter-3-linux-monitoring/README.md)

This module is based on my long experience looking after servers, performance, monitoring and diagnosing issues. </br>
This is not your usual average Linux CPU monitoring guide. </br>

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

In Chapter 2, I suggested thinking of a CPU Core this as a spinning wheel, each "wheelspin" is a CPU cycle that processes 1 task. </br>
When the wheel is spinning, it cannot process another task and has to finish its cycle so another task can hop on. </br>

This means that tasks may queue to wait their turn to get onto CPU. </br>

### CPU Multitask execution

Because a single CPU core is extremely fast as mentioned above, it may process tasks extremely fast and appear as if its completing them all at once. </br>

However, we now know that a single CPU core can only perform one task at a time. </br>
To speed this up, the CPU has multiple cores for true paralel processing. </br>

## Why is this important for DevOps, Platform, Cloud & SRE Engineers ?

Its important to understand how CPU works for the following reasons:

* CPU monitoring and usage visualization
  * Because of the way CPU works, How does the operating system determine CPU usage ?
  * If the operating system tells us, we have 65% CPU usage, How is that determined ?
  * Because of the way CPU works, the operating system has to deal with complexities in translating CPU usage to human understandable form.
  * Overall CPU as a `%` gives us an indication how busy a server is, but does not always tell the full story

* In addition to the above, we may need to understand the `%` of CPU usage <b>on all cores</b> of CPU. </br>

###  Understanding Single vs Multithreaded applications </br>
 
Since we know that a CPU core can only execute one task at a time, it's important for engineers to know how to take advantage of this and also avoid the pitfalls.  

Example 1:  If you have a task which may contain poorly written code, it could keep CPU cores busy unnessasarily, causing other tasks to queue up for execution. This can slow the entire system down. </br>

Example 2: If you have code that may be poorly written, you could end up in situations where you are only utilizing 1 core during task execution. This is common when engineers write loops in their programs. </br>
This means your application is not running optimally and not utilizing all available CPU cores. </br>

Example 3: Another example of poorly written code where one tasks is awaiting on another task, may end up in whats called a CPU "Deadlock". This occurs when all executing tasks are waiting on each other in a circular reference. </br>

To solve the above issues, programming, scripting and runtime frameworks allows us to write our code in such a way that we can create whats called "threads", "worker threads", or "tasks". </br>

Web servers are a good example of this. When an HTTP request comes to a web server, the web server can create a "worker thread" to handle that request. Then that request gets executed on a CPU core, the web server may issue another "worker thread" to handle other incoming requests. </br>

This way, a web server can handle mutlitple incoming requests and makes use of all available CPU cores on the system </br>

We may call the HTTP request handler of the web server a mutlithreaded application. </br>
The web server code for handling each HTTP request may be viewed as "single threaded" code </br>

We can take a look at two bash scripts that I have written that demonstrates single vs multithreaded code </br>

If we execute `./singlethread-cpu.sh`, we will notice that only one core of CPU is busy at a time. </br>
Now because we execute a loop, each iteration of that loop will run on a different core. </br>
Bash itself is single threaded, so this script needs to be optimized if we wanted to make use of all available CPU cores. </br>

If we execute `./multithread-cpu.sh`, we will notice all CPU cores get busy. </br>
This is because in this script, we read the number of available cores with the `nproc` command. </br>
Then we loop the number of available cores and execute our `simulate_cpu_usage()` function. </br>
At this point it would technically still be single threaded, because its still a loop and does not create more threads or processes. </br> To get around this we use a special character in bash called `&` at the end of our function. </br>

In Bash, the `&` character is used to run a command or a script in the background. When you append `&` to the end of a command, it tells the shell to execute the command asynchronously, allowing the shell to continue processing subsequent commands without waiting for the background command to complete. </br>

So if we have 8 CPU cores, our script will spawn 8 instances of our function, and each will run on a different CPU core. </br>

This means our application a is mutlithreaded application! </br>

### CPU Busy vs. System Busy: Understanding the Difference

Common scenarios where CPU usage interpretations may lead to misdiagnoses include:

When CPU overall CPU usage is low, applications or systems can still be slow if there are other bottlenecks </br>

<u>For example:</u>

* Poorly written code that causes single-threaded execution can result in low CPU usage and slow application performance, which may be difficult to detect when examining overall CPU metrics.

* CPU usage may be low, but an application can still slow down if it is experiencing network or disk I/O bottlenecks.

It's important to be aware that having overall low CPU usage, does not always mean your systems are performing optimally, and there could be other bottlenecks and symptoms of slowness. </br>

## How is CPU usage measured 

From the above explanation, we now have a high-level understanding of CPU usage. Measuring CPU usage can be challenging because many processes alternate between active execution and waiting to be executed. </br>

Most operating systems display CPU usage as a percentage, indicating the proportion of CPU resources used across all cores. Calculating this percentage is complex, and operating systems strive to provide the most accurate measurement possible. </br>

The percentage display mostly makes it easier for humans to interpret if CPU usage is high or not. </br>

## CPU monitoring tools for Linux 

### top
[top](https://man7.org/linux/man-pages/man1/top.1.html) is a native linux monitoring tool that provides an overview of the current system load. </br>
This will give us overall CPU usage for processes as well as other performance metrics

### htop

[htop](https://www.man7.org/linux/man-pages/man1/htop.1.html) is a very popular alternative to `top` as it allows us to be able to scroll horizontally and vertically as well as having other features such as searching and filtering processes and more. </br>

### sysstat tools 

[sysstat](https://github.com/sysstat/sysstat) is a Linux system performance tool for the Linux operating system. </br>

The `sysstat` package provides us with may performance monitoring tools such as `iostat`, `mpstat`, `pidstat` and more. </br>
All of these providing insights to CPU usage and load on the system </br>

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
This means that if the system is only recently busy, the average CPU usage may be very low because its an average from the time the system started up </br>


To get the current stats, we can provide an interval and we can also provide a count of snapshots we want of the stats. For example, we can get stats for every second, with a total of 5 snapshots. 

We can also provide `-c` option, which states that we are only interested in CPU stats and not I/O stats. </br>
```
iostat -c 1 5
```

Let's understand the output. This output is pretty common across `sysstat` tools

* `%user` 
  * Percentage of CPU utilization that occurred while executing at the user level (application).
* `%nice` 
  * Percentage of CPU utilization that occurred while executing at the user level with nice priority.

The Niceness value is a number assigned to processes in Linux. It helps the kernel decide how much CPU time each process gets by determining its priority.

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
Execute program with given arguments args and monitor it with pidstat.  pidstat stops when program terminates.
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

## Troubleshooting High CPU usage

The first challenge to resolve high CPU usage is to locate the overall cause of CPU usage, by using the above tools such as `top`, `htop` or the `sysstat` tools </br>

We can use filtering, searching, and sorting techniques to find and display the culprit process. </br>
All of the above tools are great at identifying the culprid process and what you need is the process ID or also known as "PID" number. </br>

The `PID` is the process identifier of a process running on the operating system. 

Once we have the process we can see the command that is causing it. </br>
Most of the time, you can identify the culprit by looking at the `COMMAND` column of `top` or `htop`. </br>

If its an executable, a script or program running somewhere we can located it using the `ps` command. </br>
This command line helps us display information about a selection of the active processes </br>
You can checkout `man ps` to get more details about all the available options. We are after the process using the `-p` option to pass the process ID of the culprit. This process ID can change, so if you follow along be aware of that. 

```
ps -p 6268 -o pid,cmd
PID CMD
6268 /bin/bash ./singlethread-cpu.sh
```

Above, we use `-p` to pass the process ID, we use `-o` to display output about the process and its command and that helps us locate the executable, in our case a bash script called `singlethread-cpu.sh` which is executed by bash, under the `/bin` folder. </br>

Now we learned in our command line episode, that the `./` allows to execute scripts and executables in the current working directory. </br>
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
3) Tools like Prometheus, DataDog, NewRelic has "collectors" or "agents" that read that data so it can be sent to a central storage
  * We do not want to store data on our servers 
  * Send data off to highly available storage
4) Prometheus has a time-series database that stores large amount of metrics
5) Visualization tools like Grafana helps us visualise the data in Prometheus by building queries and dashboards

