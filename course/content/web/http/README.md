# 🎬 The Basics of HTTP

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 4](../../../chapters/chapter-4-web-servers/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  

HTTP is extremely important in the world of DevOps, Platform engineering and SRE as it is the fundamental and popular way that systems communicate with one another. </br>
All the servers we learned about in Chapter 2, will at some point need to communicate with other servers either on the same private network, or over the internet. </br>

We learned in previous networking modules that the language used by servers is TCP. </br>
This is called a protocol.
In the world of Web servers need to send more information during communication, therefore the protocol has been extended and thus HTTP was created. </br>

As an engineer it's important to understand what HTTP looks like, how to create HTTP requests, interpret responses which will help us troubleshoot any issues that are common in the world of Web. </br>

## What is HTTP

In simple but accurate terms, HTTP is just the "language" that extends TCP. It stands for HyperText Transfer Protocol. </br> 

In our monitoring chapter on networking, we learned more about TCP and that TCP has header fields about the source IP, source port, destination IP and destination port. It has the bare minimum data needed to get a network packet going from A to B. </br>

a TCP packet has a "data" or payload component where a client and server can send arbitrary data to each other </br>

For HTTP, this payload is basically extended with more "fields", because in order to make the web work, we need things like :

* Domain of the website 
* URL path and query string
* Headers (custom fields)
* Status codes (what happened to the request)
* and more!

## Client & Server

We have already covered the definition of client and server in previous modules, but to relate it to the world of HTTP, the client is the process responsible for generating the HTTP request. </br>
The server is responsible for accepting the HTTP request, processing it and generating an HTTP response </br>

## HTTP Requests & Responses

This diagram gives us a breakdown on all the key components and terminology for HTTP requests and responses

It's important to memorize these terms, because it helps to troubleshoot errors in our console as well as in web server logs. </br>

For example, if you see an error stating `404 Not found` you should know it's coming from a server, and it means that the resource we request in our HTTP request is not found by the server.
Or if we see an error stating "Method not allowed" we should know what an HTTP request method is, and what the possible values are.

We'll explore troubleshooting in a future module, but memorizing the terminology is important as it will pop up throughout your day to day work.

![diagram](./http.drawio.svg)

### HTTP Requests

The main overview diagram shows all the terms conceptually. An HTTP request has the following fields:

* Version
* URL
* HTTP Method
* Headers (key-value)
* Body (optional)

#### HTTP Version

The version field tells us which version of HTTP is used. The HTTP protocol has evolved over many years to improve security, performance and speed of communication on the web. </br>
HTTP 1.1 was created in 1997, has evolved and is still used on the web today. </br>
HTTP 2 is newer, faster and more efficient than 1.1 as it addresses how content is loaded on a web page which helps improve page load times in browsers. </br>

#### URL

The URL has a few components to understand. The first part is the protocol, which for web is mostly either `http` (port 80 - plain text) or `https` (port 443 - encrypted)
The `:` is the separator. Then we have the domain, which gets resolved by the client to an IP address using DNS resolution. </br>
Following the domain, we have the URL path, which starts with `/`.
The URL path generally points to a resource the client is looking for or the endpoint the client wants to send a request to. </br>
The URL may end with a `?` character which indicates the start of a query string. </br>
A query string is an optional string a client can send which has keys and values separated by `&` characters </br>
For example `?version=v1&name=John`

#### Body

The HTTP Body field is an optional field containing the message content that a client wishes to send to a server. A client can send a body in an HTTP request and a server can respond with a body in an HTTP response. </br>

Generally for browsing web pages, the HTML content of a web page will be sent in the HTTP response as part of the body field. </br>

When a client is an application that wishes to make a web request or API call, it may want to send data to a web API and the data would be in an HTTP request body field. 

#### Headers

Headers are key value pairs of arbitrary data that a client sends to a server. HTTP headers contain core 
information about the request and response. This information can include things like the 
type of device the client identifies as, what type of information the client is expecting and more </br> 
Headers can be either part of an HTTP request or response or both

Example Request Headers:

```
:authority:  www.google.com
:method:  GET
:path:  /
:scheme: https
accept: text/html
accept-encoding: gzip, deflate, br, zstd
accept-language: en-GB,en-US;q=0.9,en;q=0.8
cache-control: no-cache
cookie: xxxxxxxxxxxxxx
```

### HTTP Response

An HTTP Response is what a client receives from a server in answer to an HTTP Request. </br>
Generally, all HTTP Requests should have an HTTP Response if the server was reached. </br>
If a server was not reachable, a network device in front, like a load balancer or proxy may also answer 
with an HTTP response indicating an error reaching the server. </br>
The most important field of an HTTP Response is the status code. 
The status code indicates what happened to the request. 

The HTTP Response contains the following fields: 

* Status Code 
* Headers (key-value)
* Body (optional)

The Headers and Body fields serve the same purpose as described in the HTTP Request section. </br>
The server can also set HTTP headers in a response. </br>
The Status code field is the most important field that describes what happened to the HTTP request.
This is often where engineers misinterpret or misdiagnose HTTP request issues. </br>
It is very easy to overlook certain issues when the HTTP response status code is misunderstood. </br>

### Status codes

The HTTP status code is a three digit code that tells us what happened to the request. </br>
Status codes start from 100 and go into the 500s </br>
In each increment of 100, the status codes have different meanings. </br>

<b>Important! </b> When we receive an HTTP status code in an HTTP response, this response may not always be coming from the server </br>
This is very important to keep in mind when troubleshooting or interpreting requests, responses or logs. </br>

Pay attention throughout this module on where the status code may come from </br>


![diagram](http-statuscode200.drawio.svg)

* `1xx` Informational
* `2xx` Success
* `3xx` Redirection
* `4xx` Client Error
* `5xx` Server Error

It starts with informational statuses in the 100s range. </br>

#### 2xx Success

HTTP `200` status code means that the request was successful </br>
As a DevOps engineer, there is more to understand here. Generally speaking this means that the client received a "success" response from the server. But it does not mean that the status code came from the server itself. It could come from an intermediate network device too. 
For example, a server can respond with a `200` , but a load balancer, proxy service in front of the server could also respond with a `200`. 
As an example, to improve performance of websites, many websites may sit behind a CDN, which may cache the responses. </br>
Many companies use CDN technology to improve website performance and speed because these services can cache content, preventing every single request from going to your web server. 
CDN's can also run in locations closer to the users (clients), which helps reduce overall latency too. </br>

So the key takeaway here is that status codes do not necessarily come from the server directly! </br>

[More resources](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/200) on 200 status code

#### 3xx Further Action

HTTP `3xx` status codes indicates that a further action is needed to complete the request, typically a redirection to a different URL. </br>
A common example, is when you open a website address with `http://` and the server wants you to access `https://` instead so it will return a `301` to indicate ` Permanent redirect to an `https` URL

Most Common `3xx` status codes:

* `301` Moved Permanently: The requested resource has been permanently moved to a new URL
* `302` Found (Moved Temporarily): The requested resource has been temporarily moved to a new URL

There are other `3xx` status codes too, but the aim of this course is to focus on the more common ones I think are the most important. </br>

![diagram](http-statuscode301.drawio.svg)

[More resources](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/301) on 301 status code

#### 4xx Client Error

HTTP `4xx` status codes indicate that in general, there are some issues on the client side that need to be addressed. </br>
This means the HTTP request could be invalid, malformed where a server cannot understand it, or the HTTP request needs to be authenticated or authorized. </br>
In many cases, it is something that the owner of the client can address, like a developer writing some code that is making an HTTP call. </br> The developer may need to fix the client code to fix the HTTP request.

Most Common `4xx` status codes:

* `400` [Bad Request](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/400): The server cannot understand the request due to malformed syntax. 
* `401` [Unauthorized](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/401): The client needs to authenticate to access the requested resource. 
* `403` [Forbidden](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/403): The client is not authorized to access the resource. 
* `404` [Not Found](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/404): The server cannot find the requested resource. 

![diagram](http-statuscode400.drawio.svg)

#### 5xx Server Error

The HTTP `5xx` status codes generally indicate that something has gone wrong on the server side. </br>
So it's technically the opposite of the `4xx` status codes. The `5xx` status codes can often be misinterpreted just because they may not come directly from the server side and may also be returned by intermediate network devices such as proxies, gateways etc. </br>


Most Common `5xx` status codes:

* `500`  [Internal Server Error](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/500): indicates that the server encountered an unexpected condition that prevented it from fulfilling the request. This is usually an error that occurs because of a bug, code issue or error on the server side. 
* `502`  [Bad Gateway](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/502): indicates that the network device acting as a proxy or gateway, did not receive a valid HTTP response from a server.
* `503`  [Service Unavailable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/503): indicates that the server is not ready to handle the request.
* `504`  [Gateway Timeout](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/504): indicates that the server, while acting as a gateway or proxy, did not get a response in time from the upstream server in order to complete the request.

![diagram](http-statuscode500.drawio.svg)


## Where to from here 

In this module, we primarily focus on deconstructing HTTP to understand what a request and a response looks like. </br>
Using tools like `curl`, `nc` and `wget`, one can generate HTTP requests, send them to servers, receive and interpret HTTP responses. </br>

As a DevOps, SRE or platform engineer, I always keep mentioning how important the command line is. Similarly, one should become not only familiar, but efficient with these command line tools in order to work with HTTP. </br>
Yes - there are UI tools available, but remember that in a production environment you may only have access to a production server via terminal, so it's important to practise.

In an upcoming module, we will focus on troubleshooting HTTP requests and responses using command line and browser tools as well as server logs which can also help paint a picture of what is going on. </br>
