# 🎬 Linux Network Monitoring

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 3](../../../../../chapters/chapter-3-linux-monitoring/README.md)

This module is based on my long experience looking after servers, performance, monitoring and diagnosing issues. </br>
This is not your usual average Linux network monitoring guide. </br>

Although we'll be covering theory, the objective is not to bombard the viewer with too much detail. </br>
We'll cover the theory conceptually, and then use practical examples to show you real world concepts in action </br>
This guide will feature basic and some deeper advanced topics, tools and techniques for dealing with network observability and monitoring </br> 
Although it's basic, this guide touches on all the components that I still use in day to day modern DevOps & Cloud engineering </br>

It will be important to pay attention as all of the details in this module will form the foundation of monitoring HTTP web & microservices, especially when: </br>

* One service or server cannot talk to another service or server
* Troubleshooting connection errors 
* Understanding latency
* Understanding basic network bottlenecks 

## Network usage - How the Network works

In our chapter on Operating Systems, we covered the basics of system resources that the operating system manages, including the Network which is used by processes to communicate with one another.</br>

Processes can communicate with one another on the same server or across servers, if the network allows it </br>
Unlike CPU, memory & disk, the network has a few components to understand, as each one of the components can cause connection failure, errors, delays and bottlenecks

In the first module, we looked at a high level overview of networking. And in this below diagram, I highlight some of the components we will cover in this module. </br>
As an engineer you want to have an understanding of how servers\processes talk to one another over the network. This happens through a network connection. </br>
There are two types of connection (which are referred to as "network protocols"), called TCP and UDP. </br>
Each protocol varies slightly in the way connections are established. </br>

The box on the left is our source server which runs a process and makes a network connection to the box on the right which is another server with a process on it </br>
This could be a Web browser (left box) opening a web page (Github.com) which connects to a Web server (right box) on the internet somewhere hosted by Github. </br>
It could be Web client to server, Two applications talking to another, Two microservices, a service talking to a database, or technically anything talking to something else over a network or internet </br>

```
connection from server-a to 143.0.1.2:443
or https://143.0.1.2

              ---  144.0.1.2 -------- 143.0.1.2:443 -\                                                        
          ---/     public IP          public IP       ---\                                                    
     ----/                                                ---\                                                
  --/                                                         --                                              
+----------------------+                   +----------------------+                                           
|  private IP          |                   |   private IP         |                                           
|  10.0.0.4:1024       |                   |   10.0.0.4:443       |                                           
|            port      |                   |            port      |                                           
|  server-a            |                   |   server-b           |                                           
+----------------------+                   +----------------------+                                                                                                   
```

## Importand Network components

Let's keep referring to the diagram above, and talk about each network component in this diagram to understand network connectivity    

### IP addresses    

We covered in previous modules that IP addresses are identifiers for servers belonging to a network and a server must have an IP address in order to belong to a network. </br>
An IP address can be either public or private. </br>
Generally speaking, servers always have a private IP address when belonging to a network. A server may or may NOT have a public IP address depending on network setup and configuration. </br>

### Private IP address

For example, in the diagram, `server-a` has a private IP address `10.0.0.4`. In our module on virtualisation we learned that a DHCP service provides us with that private IP when our server joins the network. Virtualization software as well as Cloud providers will generally handle this IP address assignment for you. </br>

It's also important to note that every server will have an IP as `127.0.0.1` which we also call "localhost". IF a server refers to this IP, it is technically referring to itself </br>
For example, opening a Web browser and going to address `http://127.0.0.1` will go to itself, so we would need a Web server running on our server that we can reach via that IP address </br>

### Public IP address

For `server-a` to have a public address, it depends on network configuration. </br>
Generally speaking, the Operating system on a server is configured to have what's called a "gateway" IP address so it knows where to send all outbound network traffic. </br> On home networks, this "gateway" address would usually be your home router address. </br>
Virtualization software and Cloud providers will also generally handle this for you and you don't need to worry about setting up gateways. </br>
In our module on Virtualization and Servers, our software used `10.0.0.0` as the gateway address and software would send that to the Host machine gateway and that way it ends up going out via the router on the home network. </br>

So basically that "gateway" is the gateway to the public internet as the network goes to your router and the router gets a Public IP address from your Internet service provider. </br> That's why when you reboot your home router, your Public IP address may change. </br>

A similar architecture is generally followed in company and office networks. Your computer in the office will route outbound traffic to a network device or router and that will have a Public IP address provided by the company's ISP. All similar to what is shown in the above diagram</br>

In the cloud, servers would generally have a Public IP address you can visibly see in the cloud provider web interface </br>
So each server could have its own Public IP address. Cloud providers also allow you to remove the Public IP address, which renders this server completely private and not accessible  from public networks </br> 

It's important to know that Public IP addresses are used for both inbound and outbound traffic. </br>
So network requests can go from `server-a` to the router, and out via the router's Public IP address </br>
`server-b` or any destination that receives requests from `server-a` will see that it originates from the Public IP address we have for `server-a` <br>
If `server-b` needs to respond to that request, it may just respond to `server-a` over the same connection 

Because the illustration shows a network request from left to right, it's important to know that request can also go from right to left </br>
However to do this, `server-a` needs to listen on a port and have a process running that can accept requests. Also the router device on the left needs to have a "port forwarding" rule to tell the router which Private IP address to send all traffic that is coming over a given port. </br>

Therefore a server and it's router needs to be configured in order to allow network requests to flow all the way through

### Server VS Client

In a previous chapter we talked about "servers" and what a server is. </br> 
This is not to be confused with the concept of servers and clients in networking. </br>
In networking terminology, <u>clients</u> are referred to as the server or process that is <u>initiating the network request.</u> </br>
In networking terminology, <u>servers</u> are referred to as the server or process that is <u>receiving the network request.</u> </br>

### Ports & Connections

In order for a client and server to talk, a network connection must be made. </br>
The client will need a private IP as we've mentioned earlier and it will also need a source port for the connection. This is so that the reply can find its way back to the client. The source port is generally assigned by the client's Operating System. </br>
Source ports are limited and each Operating System can have different limits for the number of source ports it can allow. This means that we could have port exhaustion if a client tries to create too many connections. </br>

Once the client has a source IP and Port, it establishes a connection to the destination IP address. Now there are some technical nuances to establishing network connections and there is more to it, however I'll be keeping this brief and simple. </br>
In my opinion, a simplified understanding is always a better place to start instead of drowning in the depth of theory and details. </br>

When it comes to monitoring we'll have everything we need to know to form a great fundamental understanding in troubleshooting systems. </br>

Now this connection attempt from the client will end up at a destination server which would have a process running on it and listening on a port. This process could be a web server or application. </br>

To accept a connection, a process must "listen" on a port </br>
That port must also be open on the server, meaning no firewall or antivirus should be blocking that port </br>
If there is a network device, proxy, load balancer or router in front as per our diagram, that device needs a port forwarding rule to send traffic to that destination server on a given port. </br>

Please make an important note here, when you see "Connection Refused" It generally means there is no process or application listening on a destination port you are trying to reach. This is a popular error that's often misinterpreted by developers and engineers. 

Another important error is "Connection Timeout". If you see this error or simply a network request hanging, it generally means that the port you are trying to reach is being blocked by something. This could be a cloud security rule , firewall , network device like a router etc. 

### Network Bandwidth

Once a connection is established between client and server, then the client can start sending network requests and the server can respond with network responses. These requests and responses often contain data. </br>
This can be a client web browser getting HTML and web page content from a server, it could be a web browser client calling an API service for data, it could be two microservices communicating with one another. </br>

These network requests and responses may sometimes contain large datasets. These requests and responses generally take up what's called network bandwidth. </br>
Network bandwidth is limited by network speeds which can involve the client and server network interfaces (or network cards), network cables, devices and ISP speeds </br>

Bandwidth is often monitored and measured in bytes per sec, megabytes per sec, gigabytes per sec, etc. </br>

If you are using the cloud, the server will have limits and specifications on its network bandwidth capabilities. </br>

You shouldn't have to worry about this most of the time unless you are working for a company that deals with large data transmission where services that talk to each other exchange large messages. Most servers in the cloud can handle pretty large amounts of bandwidth, and latency will likely be your biggest challenge </br>

### Latency

Latency is the time taken for the client to receive a response from the server. </br>
Latency is generally measured in milliseconds (ms) or seconds (s). </br>

In my experience, the main way I monitor latency is generally through proxy or web server logs. Web logs will generally list each network request and details about the request such as latency so we know how long it took for the server to respond to the client. </br>
This helps us understand timeouts, and long running requests which could impact customers </br>

In our Chapter on Web servers we will take a look at web server monitoring which covers this topic in further detail. </br>
To summarise, generally when a server is slow and not responding to requests from clients in a timely manner, we can get evidence of this by looking at web server logs. </br>
Network devices in front of web servers, like load balancers , proxies or what we call "ingresses" may also provide web traffic logs which can indicate latency of requests as well. </br> This is a key focus point when monitoring HTTP web server traffic performance and stability. </br>

### Network Protocols

There are two major network protocols you will come across in the field </br>
The two are `TCP` and `UDP` </br>

So when we talked about network connection and how the client and server establish this connection I mentioned there are some nuances to how this connection is established. </br>

`TCP` is the main network protocol used by the Web because it's designed to be reliable. </br> Networks are flakey meaning a network packet is never guaranteed to arrive at from client to server </br>
To make the network more reliable, TCP involves a handshake and a few network requests back and forth between client and server to establish connection. </br>
This network connection handshake in TCP is designed to help ensure connections are established when the network can be flakey in nature. </br>

This comes at a performance cost, therefore there is another protocol called `UDP` which is more of a "send and forget" type of network request. Where a client sends a request and waits for a response and will simply retry if it does not get a response. </br>
A client may throw an error after a few retries when a UDP request fails. </br>

## Making a Server

We have learned that a server is the receiver of network traffic </br>
A server has to run a process that listens on a port that is ready to receive traffic </br>

If this server needs to be accessible on a private network only, we just need to ensure we can access it from another server on the same network and no firewall is blocking the port. Sometimes firewall software on a server can block ports. Linux has a basic network firewall and Windows has a firewall too. </br>

Remember what I said - If we see "Connection refused" when trying to connect to our server, it means the port we are trying to connect to is not listening. </br>
If we see "Connection timeout" it's generally a firewall blocking the request, so it hangs and times out </br>

It's important to learn how to validate if a server port is open. </br>
To do this lets use a tool called netcat or `nc`:

```
nc -zv localhost 12345
nc: connect to localhost (127.0.0.1) port 12345 (tcp) failed: Connection refused
```

- The `-z` flag can be used to tell nc to report open ports, rather than initiate a connection
- The `-v` flag tells `nc` to turn on verbose output so we can get more helpful information about the port being open or not. 

Let's use `nc` to start a server. </br>
We can create a script to start our server if we wanted:

```
echo "Starting TCP server on port 12345..."
nc -lk 12345
```
- The `-l` flag tells netcat to listen on a give port for incoming connections
- The `-k` flag tells netcat to listen for another connection once it has received it. Without `-k`, netcat will close once it receives one connection

We can leave our server running in a terminal and open another to test the port </br>
This helps us during monitoring to test whether a server is listening on a port or not:

```
nc -zv localhost 12345
Connection to localhost (127.0.0.1) 12345 port [tcp/*] succeeded!
```

## Making a client

Now that we have a server , we can make a client that sends requests to the server. </br>
The whole point of this exercise is to learn about how client and server connectivity works, so that we can learn how to monitor these requests and connections </br>
We can use the `nc` command to troubleshoot and test connectivity between clients and servers. </br>
We will also learn about a couple of other important and popular network commands. </br>

We can use the `nc` command to create network requests. </br>
Send one request to our server but hold the connection: 

```
echo "Sending one message and holding connection!" | nc localhost 12345
```

The above is a basic TCP client. </br>
Another network protocol which is built on top of TCP, is `HTTP`. </br>
HTTP is used in Web communications and we'll learn more about the Web in a future Chapter. This is how we create an HTTP connection and send an HTTP message:

```
curl -X POST http://localhost:12345 -d "test"`
```

We'll get into more variations of what the client can do. But before we do this, we need to learn a couple of popular network tools and we can monitor the network at the same time. 

## Network Connection Monitoring Tools

To monitor and further understand TCP connections, we can use two popular tools:

- `ss`
  - Another utility to investigate sockets. <br/>
    `ss` is used to dump socket statistics. It allows showing information similar to netstat.  It can display more TCP and state information than other tools.
- `netstat`
  - Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships

We can monitor current TCP connections and see our server listening and also see the current in-progress open and established TCP connections between client and server

ss:
```
ss -a | grep 12345
```
netstat:
```
 netstat -a | grep 12345
```

This will help us understand the connection lifecycle as well. We can view the connection state, in this case it's `LISTEN` and `ESTABLISHED` </br>
We've learned now that `LISTEN` is for servers generally listening on a port for inbound connections. `ESTABLISHED` is the state when a connection is open and the client can send and receive messages to each other. </br>
In our example client and server, we are keeping the connection open so we can observe it in our command line tools above. </br>

TCP socket states represent the various stages a TCP connection goes through during its lifecycle. </br>

Interestingly if we close our client by pressing `CTRL+C`, we quickly use `ss` or `netstat` to check the connection and we will notice the `ESTABLISHED` has gone and we now see a `TIME_WAIT` state. </br>

This is when a connection has been used by client and server and is about to be closed.The connection is now in a "recycling" state where the operating system will get to reuse that port. Generally a connection will spend roughly 60 sec in `TIMED_WAIT` before the operating system will be able to re-use that connection </br>
The time wait timing can be adjusted in Linux, if we require a lot more connections quickly, we may reduce the time wait time </br>

We can run the following loop that keeps sending messages every 5 seconds and will close the connection after sending each message. This allows us to observe `TIME_WAIT` sockets

```
echo "Sending messages one after another!"
while true; do
  echo "Sending message and closing connection!" | nc -q 5 localhost 12345
done
```

Here are the TCP socket states in order:

* `CLOSED`: The initial state. No connection exists.
* `LISTEN`: The server is waiting for incoming connection requests.
* `SYN_SENT`: The client has sent a connection request (SYN) and is waiting for a matching connection request (SYN-ACK) from the server.
* `SYN_RECEIVED`: The server has received the client's connection request (SYN) and sent a connection acknowledgement (SYN-ACK), waiting for the final acknowledgement (ACK) from the client.
* `ESTABLISHED`: The connection is open, and data can be sent and received between the client and server.
* `FIN_WAIT_1`: The client or server has initiated the connection termination and is waiting for the other side to acknowledge (FIN).
* `FIN_WAIT_2`: The side that initiated the termination has received the acknowledgement (ACK) of its FIN and is waiting for the other side to send its FIN.
* `CLOSE_WAIT`: The side that received the first FIN is waiting to send its own FIN.
* `CLOSING`: Both sides have sent FINs, but neither has received the final acknowledgement (ACK).
* `LAST_ACK`: The side that sent the first FIN is waiting for the final acknowledgement (ACK) of its FIN.
* `TIME_WAIT`: The side that sent the final acknowledgement (ACK) is waiting for a period of time to ensure the other side received it.
* `CLOSED`: The connection is fully terminated, and no further communication is possible.

## Network Traffic Monitoring

For all the above, we monitored network connections. </br>
But once a connection is established the client and server can go back and forth and send and receive messages to one another. </br>

To monitor this network traffic, I often rely on logging. </br>
Traffic will flow from client to server, and the applications we run on servers generally have capability to log traffic requests. </br>
All popular web servers have this feature that you can configure. </br>

Once configured, the server process will write incoming traffic logs to a file, where each line in the file represents a request. </br>
It is generally in a similar format like this, either delimited by spaces or commas separating each field: 

```
<date> <client-IP> <request-info> <status> <time-taken-in-milliseconds> <etc>
```

With these logs, we can monitor where traffic is coming from, how long the request has taken to complete and also whether or not it was a success or failure. </br>
This will be of great help when we start troubleshooting slow requests which can result in systems slowing down </br>
We'll be covering web server logs and log monitoring in more detail in our chapter on Web servers </br>