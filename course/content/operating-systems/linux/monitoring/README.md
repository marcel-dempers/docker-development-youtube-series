# 🎬 Introduction to Linux monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../chapters/chapter-3-linux-monitoring/README.md)

## What is Monitoring

Monitoring is the process of collecting, analyzing and using data to track the performance and health of systems. </br>

Monitoring involves the use of tooling to:
* Capture, collect or extract data from systems, services, applications, processes, etc
  * This type of data could be logs, metrics or traces 
* Store this raw data in a storage system where it can be processed
* Process the raw data so it can be analysed
* Visualise the data, which enable teams to track and analyze health and performance of these systems, services, applications and processes. </br>
* Detect and notify engineering teams if any potential issue occurs that needs attention
  * For example, if CPU stays high for a certain amount of time, or disk usage runs over a threshold, or a process crashes with an error </br>

## What is Observability

Observability is a broader concept that refers to the ability to understand the internal state of a system based on the data it produces. </br> 

It goes beyond traditional monitoring by providing deeper insights into the system's behavior and enabling more effective troubleshooting and root cause analysis. </br>

Some examples of Observability includes:
* Logs 
* Metrics
* Traces

Observability is often a more investigative approach to monitoring in order to find bottlenecks in a system or root cause analysis for issues. </br>

## Monitoring examples

The most basic form of monitoring, is to use tooling that the operating 
system provides a way to look at a system's basic resource utilization and analyze its health and performance. </br>

### Average System Load (memory+cpu)

For example, the operating system provides a native command called `top` to analyze and monitor overall system load and some performance metrics </br> `top` is another command line executable that lives at `/bin/top`

If we run `top` on our Linux server, we see system load averages, current memory and CPU usage as well as all the processes and threads running on our system. </br>
That gives us an overview of Memory and CPU usage </br>

Load averages are made up of 3 important numbers. Each number is an average system load for a given timeframe. </br>
The first number is the average system load in the last minute, followed by 5 min, and 15 min for the last number. </br>
This tells us if there is ongoing performance load or just a small recent spike in load average. </br>
In simple terms it tells us if the system was recently busy, or constantly busy </br>

We can also see load averages by printing out the load averages in the location `/proc/loadavg'

Below we see load averages were:

*  `0.62` in the last 1 minute
*  `0.14` in the last 5 minutes
*  `0.05` in the last 15 minutes

```
cat /proc/loadavg
0.62 0.14 0.05 1/568 1807
```

It's good to know that Linux stores a ton of process information in the `/proc` folder. </br>

### Network Utilization

In a previous module, we briefly covered networking as we created and configured a network for our virtual server and we learned about IP addresses. </br>

For servers to communicate with other servers in a network or even over the internet, they need to have an IP address. </br>

In addition, to connect to another server, we need an IP address of that server as well as a port number. </br> All network connections occur over a network port. </br> Ports are a limited resource and a server may only have so many ports available. A server can also only support a certain number of network connections over a given port. </br>

```                                                                                                  
              ---  144.0.1.2 -------- 143.0.1.2:443 -\                                                        
          ---/     public IP          public IP       ---\                                                    
     ----/                                                ---\                                                
  --/                                                         --                                              
+----------------------+                   +----------------------+                                           
|  private IP          |                   |   private IP         |                                           
|  10.0.0.4:1024       |                   |   10.0.0.4:443       |                                           
|            port      |                   |            port      |                                           
|                      |                   |                      |                                           
+----------------------+                   +----------------------+                                                                                                   
```

#### Network resources

There are a number of resources we need to consider when monitoring networks 

* IP addresses
  * Every network has a range of IP addresses, which is limited. A network can run out of IP addresses.
* Ports
  * Every network connection needs a source and destination port number. 
  * Source ports are allocated by the operating system when we make a network connection. The operating system assigns an ephemeral (temporary) port number from a predefined range of ports. This range is typically from 1024 to 65535, but it can vary depending on the operating system configuration.
* Connections
  * A server can only make and receive a limited number of network connections.
  * We may often be tasked to monitor how many connections a server has open, so we know if connections are being exhausted or not.
  * Connections can be dependent on hardware support and operating system settings. 
* Bandwidth
  * I'd like to think of bandwidth as the speed at which our server can operate on the network
  * bandwidth is dependent on network speeds and network hardware that the server uses

#### Network monitoring tools 

If we run commands like `netstat`, or `ss` (Socket stats) we can see network connections on our server which may help us review network connectivity on our server <br> 

`ss -s` gives us a summary </br>

These tools can assist us in troubleshooting if a network port is open. </br>
When we host applications like microservices, web services, databases or applications that accept network connections, these applications usually accept connections by listening on a network port. </br>

`netstat` can also be used to gather networking statistics 

```
netstat -a -l | head -n 10
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State      
tcp        0      0 localhost:ipp           0.0.0.0:*               LISTEN     
tcp        0      0 localhost:34521         0.0.0.0:*               LISTEN     
tcp        0      0 localhost:34789         0.0.0.0:*               LISTEN     
tcp        0      0 localhost:36491         0.0.0.0:*               LISTEN     
tcp        0      0 localhost:44305         0.0.0.0:*               LISTEN     
tcp        0      0 localhost:domain        0.0.0.0:*               LISTEN     
tcp        0      0 Marcel-Laptop:49910     162.159.36.20:https     ESTABLISHED
tcp        0      1 Marcel-Laptop:57518     169.254.169.254:http    SYN_SENT   

```

### Disk space

To monitor disk space usage, we can use `df -h` </br>
If we run `df -h` on our server, we can see file system usage on our server. </br>

```
df -h
Filesystem                         Size  Used Avail Use% Mounted on
tmpfs                              197M  1.1M  196M   1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv   12G  4.4G  6.4G  41% /
tmpfs                              985M     0  985M   0% /dev/shm
tmpfs                              5.0M     0  5.0M   0% /run/lock
GIT                                192G   83G  109G  44% /home/devopsguy/gitrepos
/dev/sda2                          2.0G   95M  1.7G   6% /boot
tmpfs                              197M   12K  197M   1% /run/user/1000

```

This helps give an overview if there are any file systems low on disk space that we need to look into </br>


We can also analyze space in a file system or in specific directories using the `du -h` command </br>
This command takes a directory, in our case we can check from the root directory `/` and dig further down to find large directories or files. </br>

We use `sudo` here as we need it to access certain folders outside of our home directory. </br>
```
sudo du -h -d 1 /
```

## Basic Monitoring Commands

* `top`
* `htop`
* `netstat`
* `ss`
* `df`
* `du`
* `vmstat` (provided by the `sysstat` package)
* `pidstat` (provided by the `sysstat` package)
* `iostat` (provided by the `sysstat` package)
* `mpstat` (provided by the `sysstat` package)

## Logs vs Metrics vs Traces

Logs, metrics, and traces are different data formats produced by systems that help us understand various performance and health aspects.

The processes to produce these data formats differ and require various tools involving both developers and operations.

For example, developers use logging SDKs and configure log verbosity for application logs. These logs are then collected, processed, and stored by tools set up and configured by DevOps engineers for analysis."

### Logs

Logs are generated by applications and programs to provide detailed records of activities and events occurring within software applications. <br/> 
They capture information such as errors, warnings, informational messages, and debugging data, which are essential for monitoring, troubleshooting, and analyzing the behavior and performance of the software. </br>

We already have a little experience in logging in the previous Chapter, when we wrote our first bash script. We used the `echo` command to output events and activities about the execution of our script. </br>

Logs are often written to a file on disk. Applications can generally be configured to write logs to a given file path on disk. </br>
The challenge with writing logs to file is:
* Files can get too large if the application writes to the same file.
* Applications often perform log rotation so only a fixed amount of logs are written to a file before the application will start writing to a new file to prevent a single file from getting too large.
* Ensure logs are cleaned up from the file system to prevent the disk from running out of space.

Operating systems provide output streams for applications to write output to </br>

For example, in previous modules we covered the command line and these programs write output to our terminal. </br>
This output steam is called `stdout` or "standard out" </br>

It's advantageous for applications to write logs to `stdout` rather than to a file, as this avoids the previously mentioned challenges related to writing files on disk. </br>

There are a number of tools that help collect logs:

* <b>Fluentd</b>: An open-source data collector for unified logging layers.
* <b>Logstash</b>: A server-side data processing pipeline that ingests data from multiple sources simultaneously.
* <b>Graylog</b>: A powerful log management and analysis tool.
* <b>Filebeat</b>: A lightweight shipper for forwarding and centralizing log data.
* <b>Promtail</b>: An agent which ships the contents of local logs to a Loki instance.
* <b>Splunk</b>: A platform for searching, monitoring, and analyzing machine-generated big data.
* <b>Elastic Agent</b>: A single, unified way to collect data from your infrastructure and applications.
* <b>Vector</b>: A high-performance, end-to-end observability data pipeline.
  
### Metrics 

Logs are great for monitoring application behaviour, as it reports activities and events, which may include errors. </br>
However, logs can be quite heavy (to store and process) and need to be parsed and stored which can take up a lot of space. </br>
It also takes a lot of compute to process logs into analytical metrics that can be aggregated and used in real time </br>
This is where metrics help. </br>
Think of metrics as "key" + "value" pairs of data. </br>
Metrics are much smaller than logs and faster to process, summarize and perform analytical computations in real time. </br>

For example, CPU, memory and disk usage can be described in metrics format. 
The data is a lot smaller, and we can quickly calculate CPU usage over time to detect high system load. 

There are a number of tools that help collect metrics:

* <b>Prometheus</b>: An open-source systems monitoring and alerting toolkit originally built at SoundCloud. It has a multi-dimensional data model and a powerful query language called PromQL.
* <b>Grafana</b>: While primarily a visualization tool, Grafana can also collect and query metrics from various sources, including Prometheus, InfluxDB, and Graphite.
* <b>InfluxDB</b>: A time-series database designed to handle high write and query loads. It is often used for storing metrics and events.
* <b>Graphite</b>: An enterprise-ready monitoring tool that runs equally well on cheap hardware or Cloud infrastructure. It stores numeric time-series data and renders graphs of this data on demand.
* <b>Telegraf</b>: An agent for collecting, processing, aggregating, and writing metrics. It is part of the TICK stack (Telegraf, InfluxDB, Chronograf, Kapacitor).
* <b>Zabbix</b>: An open-source monitoring software tool for diverse IT components, including networks, servers, virtual machines, and cloud services.
* <b>Datadog</b>: A monitoring and analytics platform for cloud-scale applications. It provides metrics collection, visualization, and alerting.
* <b>New Relic</b>: A comprehensive monitoring tool that provides real-time insights into application performance, infrastructure, and user experience.
### Tracing 

Metrics are mostly designed to give us statistical data about applications, such as CPU, memory, disk IO usage, or even requests per second, or iterations per second of functions etc. </br>
Just like with logs, Developers can add metrics to their applications too.

However, when we have multiple applications and web services, microservices all talking to one another over networks, it can be useful to trace a network request all the way through a system to monitor a full transaction. 

For example, a customer interacts with a website in the browser. That makes a web request to our front end. Our front end makes a few requests to back ends, and some back ends interact with one another and with databases.

This is where Tracing comes in. Tracing is a technology used by applications and some web servers to inject tracking data into requests as it flows through systems. Then we can use visualization tools to see an entire transaction with all its requests

A lot can happen to form a transaction, and sometimes systems can slow down. </br>
Tracing is very useful to detect bottlenecks in a distributed system

There are a number of tools that help collect metrics:

* <b>Jaeger</b>: An open-source, end-to-end distributed tracing tool originally developed by Uber Technologies. It is used for monitoring and troubleshooting microservices-based distributed systems.

* <b>Zipkin</b>: An open-source distributed tracing system that helps gather timing data needed to troubleshoot latency problems in service architectures.

* <b>OpenTelemetry</b>: A collection of tools, APIs, and SDKs that can be used to instrument, generate, collect, and export telemetry data (metrics, logs, and traces) to help you analyze your software’s performance and behavior.

* <b>New Relic APM</b>: Provides distributed tracing capabilities to monitor and troubleshoot application performance issues.

* <b>Datadog APM</b>: Provides end-to-end distributed tracing from frontend devices to backend services, with automatic instrumentation for popular frameworks.

