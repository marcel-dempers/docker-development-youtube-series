# 🎬 Introduction to the Web

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../README.md) for more information </br>
This module is part of [chapter 4](../../chapters/chapter-4-web-and-http/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  

It focuses on my personal approach to understanding the Web as well as how I troubleshoot.

## What is the Web 

The Web is just a bunch of computers and devices that are connected to a network. This is one giant network made up of public and private networks </br>

In our module on servers and virtualization, we created a small private network and we learned that this network is connected to the internet via our gateway (which in my case, was my home router) </br>

In our module on Network monitoring, we covered all of this in a diagram. I will link it here for a recap </br>

The internet is technically private networks that are connected by public networks, similar to our diagram below

![diagram](../operating-systems/linux/monitoring/network/network.drawio.svg)

We'll need to ensure we understand the diagram, because the network components form the core architecture of the internet. </br>

### Understanding the big picture

This entire chapter will break down the following diagram and deep dive into each component. </br> It is critically important as an engineer to understand these fundamentals as it plays a role in building , testing , troubleshooting and working with almost every aspect of DevOps </br>
This diagram puts together all important aspects and components that play a part in Web and HTTP </br>

![diagram](web.drawio.svg)

## What is HTTP

In our previous Chapters on servers and network monitoring, we covered the first diagram showcasing TCP connections between a client process and a server process. </br>
We learned that TCP is the lowest level and smallest form of network connection. </br> 
We learned that a TCP connection goes through various states and we even saw this in a hands-on practical using tools like `netstat` and `ss` to monitor network connections </br>
TCP connections allow clients and servers to transport network packets between each other. This is how "data exchange" happens. </br>

Now in the world of Web, we need a standard of how these network packets need to be structured, so that devices on the web can communicate effectively and securely </br>

Just like human languages have structures, rules and standards which allow us to communicate in ways we can all understand. </br>

When you type a website address into your browser, your browser creates a TCP packet in a "special format" that a web server can accept, understand and it can respond with requested content which your browser can then understand and display to you. </br>

This "special format" of a TCP packet is called "HTTP" </br>
So in a nutshell, an HTTP packet is just a TCP packet, but it has more information that has to conform to certain standards </br>

In order to create an HTTP packet, we can use scripts, tools or programming languages. But before we go there, lets understand more overall components

<i>Note: We will have an entire module dedicated to HTTP</i>

## Client & Server

### What is a Client 

Zooming into our above diagram on the Web, a client is a process that initiates a TCP connection. In the world of Web, we talk about HTTP. </br>

An HTTP client will create an HTTP packet that conforms to the HTTP specification. This HTTP packet is referred to as an HTTP Request.

An HTTP request has information and structure that a "Server" will understand. </br>

### What is a Server 

Looking at our diagram once again, a server is a process that accepts a TCP connection. In the world of Web, we talk about HTTP </br>

An HTTP server will listen on a port and accept a connection. It understands HTTP and will generally only respond to TCP network requests that conforms to HTTP specifications </br>

An HTTP server (or a Web server) will generally respond with an HTTP response

In our diagram above, we can visualize the incoming connection to a server and see the HTTP request as well as the HTTP response. </br>

When it comes to development of HTTP based services like microservices, it is crucial to understand HTTP requests and responses. DevOps engineers spend a lot of time analyzing HTTP traffic to understand performance, health , and security of an overall system. </br>

The HTTP requests and responses make up a large portion of monitoring, so understanding these will help you make sense of web server logs to understand what is happening between clients and servers in your environment. </br>

## HTTP Requests & Responses

### HTTP Requests

In our networking diagram, we can visualize a pure raw TCP connection flow. We can see that in order to satisfy the connection, a TCP packet would need header details such as source and destination IP addresses & ports to know where to connect to. </br>

In order for clients and servers to send each other "data" this data forms part of the TCP payload which is just arbitrary raw binary data.

An HTTP request is simply extra metadata that is sent in a TCP network packet as part of the TCP payload. 

In the world of Web, we require extra fields, for example a web browser needs a URL of the website you need to access. </br>
It needs to know if the website supports HTTPS or HTTP. When a user needs to log into a website, the browser needs to pass this login information to the server. </br>
As users interact with websites, data is passed back and forth between the browser (acting as client) and a web server (acting as server).

These all form part of the "metadata" I described, which needs to be added to a TCP network packet. Remember that under the hood of all this, there are still pure raw TCP connections.

Let's refer back to the diagram, and I will briefly walk through these fields that are required and supported by HTTP </br>

### HTTP Responses

A client is responsible for creating the HTTP request. Once sent and received by a server, the server will process it and generally respond with an HTTP response.
An HTTP response also has certain fields that it both requires and supports. 

Let's refer back to the diagram, and I will briefly walk through these fields that are required and supported by the HTTP response </br>

As this is just an overview module, we will deep dive this protocol further in a later module. </br>

[mozilla.org](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Overview) is a great resource for exploring the world of HTTP and its specifications.

#### Creating an HTTP request

I use `curl` on a daily basis to test and troubleshoot web servers. It's essential for every DevOps, SRE, Platform or Cloud engineer to master. </br>
The `curl` and `wget` tools are available in almost all Linux distributions and may even be available on many of your production servers. </br>
I often use `curl` on production servers to test if web servers are working correctly or if they are responding with the correct HTTP responses. </br>

We can make a simple HTTP request with curl like so:

```
curl https://marceldempers.dev
```
We can pass `-v` to indicate `curl` to give us more verbose data on the HTTP request and response. Doing so allows us to see request and response headers and other fields we talked about in HTTP:

```
curl -v https://marceldempers.dev
```

We can construct a HTTP request ourselves by passing some of the fields we talked about

```
curl -X GET  \
-H "Host: marceldempers.dev" \
 https://marceldempers.dev
```

### Troubleshooting HTTP requests and responses

1. Command Line

We can use `curl` alongside `-v` to troubleshoot HTTP requests responses using the command line. </br>

2. Browser

It's also important to get used to doing a similar thing in the browser. Browsers support developer tools which have network tooling built-in. We can use this to inspect HTTP requests and HTTP responses for websites to analyze web traffic </br> 

When analyzing HTTP traffic , it's important to pay attention to HTTP responses and understand HTTP status codes in depth. In my experience this is the simplest field with the largest oversight when troubleshooting HTTP traffic. </br>
It's important to study and understand the meaning of each HTTP status code. 

3. Logs

We can also gather HTTP details like URLS, Status Codes, IP addresses, Latency details etc from Web server logs. 
All web servers have configuration we can enable that tells the server to output HTTP request and response logs. 

We will use steps 1 - 3 extensively when troubleshooting HTTP requests and responses.
We'll further explore web monitoring in more detail in a future module. </br>

## What is DNS

Looking back at our network diagram, notice that I only have IP addresses for client to server communication. </br>
This is because TCP uses only source and destination IP addresses and fundamentally uses IP's to forward traffic. 

For HTTP, you may have noticed we have an HTTP header in our request that indicates `Host: marceldempers.dev` and there is no IP address there. </br>
It is up to the client to resolve the host name and work out what the IP address is, so that the client can construct a TCP packet. </br>

This means the client takes the URL portion which is the domain, and performs a DNS lookup to get the IP address. </br>

We can do this ourselves with command line tools like `nslookup` or `dig`. </br> 

<i>Note: `dig` needs to be installed with ` sudo apt-get install -y dnsutils`</i>

```
nslookup google.com
```

DNS servers will resolve domains that we purchase, to IP addresses. </br>
Important to know that IP addresses can be dynamically issued by your Internet service provider, so your public IP address of your server could change depending on where you host it. </br>

All cloud providers generally provide static IP addresses which do not change. This helps us point our DNS names to these static IP addresses when running web servers. </br>
Hosting a web server on a public IP address that is dynamic is a little more complicated, but it is possible. </br>

We will further explore the world of Domains, DNS and SSL\TLS in an upcoming module </br>

## What is a web server

A Web server is a type of server that serves HTTP requests and produces HTTP responses. </br>
Web servers listen on ports for incoming TCP connections, just as illustrated in our above network and web diagrams. </br>
Web servers basically serve files, which can be made up of different types of content from text files, to HTML, javascript, CSS stylesheets and more. A web server could host a website or web application such as a microservice. 

### Examples of Web servers 

Popular Web servers that I have used:

* NGINX
* Apache
* Caddy 
* IIS (for Microsoft Windows)

Popular programming languages also allow you to write your own or host web application code on the programming language runtime. <//br>

In this course, we will explore the most popular web servers, but that will require its own module to cover it. For simplicity of this module, I have written a [small standalone web server](.test/server) we can run. 

<i><b>Important Note:</b> Please note that links to these files can change without notice and may be different from what you read on screen in the video at the time of recording. You can always click the hyperlink in the document to get the latest path</i>

We use `cd` to change directory to the above web server and run it: 

```
./server
```

### Web server logs

Web servers have the capability to write logs. Logs are basically just a sequence of events and diagnostic information that helps with troubleshooting and analysis. </br>
Web servers can have two types of logs.

* Error logs: These logs are errors thrown by the web server. This could indicate problems with the web server or with its configuration
* Access logs: These logs contain a combination of important HTTP request and response fields, such as the time the request came in, the IP of the client making the request, the URL and PATH of the request, its response status code and other important data. </br> 
Each line in the access log represent important fields of a request and response

I've mentioned in previous modules that in Linux, our commands write output into `stdout` which is a [Standard Stream](https://en.wikipedia.org/wiki/Standard_streams) in Linux that writes output to a terminal that we can view. </br>

Let's send some HTTP requests to our web server using `curl`: 
```
curl -X GET http://localhost:8080/home.html
curl -X GET http://localhost:8080/home-notexist.html
```

Now in the terminal where we started the web server, we can see our logs being output to `stdout`.
Let's analyze the HTTP logs:

```
./server 
2025/04/01 14:20:50 Starting server on :8080
2025/04/01 14:21:00 GET /home.html 127.0.0.1:37836 HTTP/1.1 200
2025/04/01 14:21:05 GET /home-notexist.html 127.0.0.1:37840 HTTP/1.1 404

```

We can also open our page in the browser and see how the browser requests additional content like the `.js` and `.css` and `.ico` files. </br>
This is because the browser will load the `home.html` page and try to render it. </br>
While rendering, if we take a look at the `home/html` file, we will see `.css` and `.js` script file references. The browser knows to get these and use them to apply styles and rendering to the site. </br>
This is why we see these additional requests: 

```
./server 
2025/04/01 16:19:44 Starting server on :8080
2025/04/01 16:19:54 GET /home.html [::1]:54736 HTTP/1.1 200
2025/04/01 16:19:54 GET /style.css [::1]:54736 HTTP/1.1 200
2025/04/01 16:19:54 GET /script.js [::1]:54736 HTTP/1.1 200
2025/04/01 16:19:54 GET /favicon.ico [::1]:54736 HTTP/1.1 404
```

Notice that when refreshing, we may stop seeing the `.js` and `.css` request. This is because the browser may cache these resources to prevent making repeat requests to the server </br>

We'll learn more about HTML, JS, and CSS in an upcoming module
