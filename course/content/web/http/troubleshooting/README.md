# 🎬 HTTP Troubleshooting

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../README.md) for more information </br>
This module is part of [chapter 4](../../../../chapters/chapter-4-web-and-http/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  

This module is where all the content from the previous modules of this chapter amalgamates and comes together. In earlier modules, we talk about the web overview, the big picture and how the web is constructed out of clients and servers that talk to each other over HTTP protocol that is built on top of TCP networking. </br>
So all these pieces come together and it's great when all things are functioning as expected, but in many cases there are things that go wrong. </br>

For DevOps engineers, this can be during setting up an environment that contains web servers. The web server may throw errors and not start up correctly.
The web server may not respond to HTTP requests as expected. The HTTP requests may be timing out because of another issue.
It can also be a web server environment that is working perfectly, but then starts throwing errors to clients, or it starts becoming slow in replying to HTTP requests. 
In this case it will be up to DevOps, SRE and Platform engineers to understand why, and go through various troubleshooting steps to work out the root cause.

There is sometimes no perfect playbook for what to do when things go wrong, and that is why these earlier modules are so important because many times it takes intuition to understand something instinctively. For a simple example, when seeing a `404` HTTP status code, one would instinctively know that there was an HTTP request to a resource that was not found on the server that is throwing the HTTP response with the `404` status code. </br>

Understanding the first part may be simple, however depending on how complex the environment is, people often make assumptions about this `404`. </br>
If a client `A` talks to server `B` asking for a resources, and server `B` gets the resource from server `C` it takes further intuition to ask yourself, "Is this `404` coming from server `B` or server `C` ?" </br>
Sometimes in complex environments its easy to make assumptions about errors, but we need to also understand the setup. Reading the correct logs of server `B` and server `C` and working out that the client here is actually server `B` (and not assuming client is always `A`) and then seeing through logs that server `C` is actually returning the `404` - is the type of intuition and skill you will develop over time. Complex environments create a lot of "fog of war" which makes it quick for people to make assumptions and start trying to fix server `B` when the resource is actually coming from server `C`.

This module will give you a basic overview of where to start troubleshooting, where to look, what to look for, and the rest will come with experience and building the intuition and understanding of all the above mentioned. </br>

# Troubleshooting client and servers

## 1. Where to start

<b>Client & Server!</b> </br>

When it comes to troubleshooting the web, always remember that there is a client and a server. </br>
I always start with the following image in my mind. This is because the problem is either at the client end, the server end, or possibly in rare cases it could be in between client and server.

![diagram](../../html/html-basic-server-client.drawio.svg)

### Be Proactive - before a problem starts

<b>Don't wait for a problem to occur, get good now!</b> </br>

Building the intuition of understanding where to start, where to look etc, will mostly come from experience. This experience can be gained from hands-on work. Like using `curl` in your daily life, to make web requests. Use command line & tools as often as you can. 

<b>Be curious and Validate!</b> </br>

When working with developers, ask them "How do I make a request to this service you are building?" </br> 
Try making an HTTP request using `curl`, validate the response. Ask questions. </br>
You might need the URL, perhaps any query string parameters, HTTP Headers and the HTTP Body. </br>
Once you get these details, make a request and validate the HTTP response details. 

This type of practice will not only make you better, but it will make it much faster for you to act and get to a root cause quicker. </br>
Experience is built on the observations you make, and when you are used to seeing things working, you will understand what a "working" system looks like and make 
better sense of a system that is NOT working. 

## 2. Observations

When a problem occurs, the first action for us is reviewing our observations. Since you are now aware of client and server, and you have experience of making HTTP requests using `curl`, you will know that the problem can occur either on the client, the server or somewhere in between. </br>

Being curious and proactively learning on working systems, you automatically know what a working system looks like, so when a problem occurs you are more likely to sense where this problem may be occuring. </br>

<b>Things to Observe: </b> </br>

* Is the error reported by a customer or user that is using a browser ? 
  * If so, what are your observations when looking at the browser developer tools ?
  * Can you reproduce this error using the browser or `curl` ?
* Can you locate a network issue with the steps above ?
* Is the error reported by a developer from software like a service or API ?
  * If so, what are the observations from the behaviour or the logs of this API ?

Remember if this is a web application or a website, the client can be a browser. If the error is coming from a web application like a web service or API or microservice, then that is our client. </br>

<b>Patterns of Observations: </b> </br>

You will find that over time, and with experience that issues are generally repeating certain patterns. 
There are a number of error types and observations you will encounter which will point us in the right direction when it comes to troubleshooting HTTP requests. </br>

Firstly, based on all the above mentioned cases, <b>where are we seeing these errors ?</b>

For the demonstration in our video, I have made a list of errors we can see using a client like `curl` with the aim of showing you if the issue is on the client or not.
These are some of the most common patterns and error messages that I have dealt with and continue to see in the world of HTTP:


| Example 1 | Command | Observation | 
|---|-----|--------|
| <b>Could not Resolve Host </b> | `curl http://nonexistent.example.com` | `curl: (6) Could not resolve host: nonexistent.example.com` |

<b>Interpretation:</b></br>

`Could not resolve host` means that the client was unable to resolve the domain or host name in our case `nonexistent.example.com`. 
Now let's be careful not to assume it's a problem with the client itself. 
If we read the error , "Could not" states the client was unable to, but does not indicate the reason or cause "why" the client could not. 

Yes, perhaps the client has internet trouble. 
In Chapter two we created a virtual server and we learned about DNS. Servers are ideally provided with DNS capabilities by the network they are on. 
In our case, your Virtual Server will get DNS from its host, and the host will get it from the operating system which may get it from your network router. If anything in that chain goes wrong, DNS could be impacted and any service or application running on our virtual server would not be able to resolve hostnames. </br>

Or maybe it's as simple as the client misspelled the hostname or got it wrong completely and it's a simple fix. </br>

| Example 2 | Command | Observation |
|----|----|--------|
| <b>Connection timed out</b> |`curl --max-time 0.1 http://example.com` | `curl: (28) Connection timed out after 100 milliseconds`|
<b>Interpretation:</b></br>

`Connection timed out` means that either the client or the server exceeded a timeout. 
Again, let's not assume where exactly the timeout occurred. In the case above, we specified a very low client timeout value and we know this value is too low to make a connection and retrieve content from a website, so in this case it's a client side timeout that was exceeded. </br>
It's important to understand the timeouts can be raised by servers as well as network devices in between. </br>
Client's can be configured to only wait a certain time and either give up, or retry. </br>
Servers also have timeout values that can be set on web servers. We know that a customer is not going to wait around for 5 minutes for a website to load, therefore we set our expectations on servers and throw a timeout if the request can not be honored in a specified time. </br>

Where things get tricky is when you have network devices like load balancers and proxies that sit in front of your web servers. 
These also have timeout values. In our module on HTTP Status code, a `504` Gateway timeout is an example of a timeout thrown by a proxy or load balancer service sitting in front of a website. </br>

| Example 3 | Command | Observation |
|----|----|--------|
| <b>Bad Request</b> |`curl -X INVALID http://example.com` -v | `HTTP/1.1 405 Method Not Allowed` OR `HTTP/1.1 400 Bad Request` |

<b>Interpretation:</b></br>
This is an example of a client making an invalid or bad request. </br> 
Generally speaking , this HTTP request should reach a server and the server rejects it </br>
However, this could also be a request that the web server could not understand. </br>
`405` status code, or `Method Not Allowed` means that the server does not allow the HTTP method that was sent. </br>
For example, a web service that accepts a payload, may only accept an HTTP method of `POST` and not allow `GET` requests. </br>
In our case `-X INVALID` is not a valid HTTP Method, so a client may even reject this request and may not even attempt to create an HTTP request at all. </br> 
`400` or `Bad Request` is exactly that, a bad, invalid or malformed request. </br>
In our example, `-X INVALID` is not a valid HTTP Method, so a server may reject it as a "Bad Request". </br>
A bad request can also be a request that has invalid payload in the HTTP Body field. For example if a web service expects JSON and you send invalid JSON or a field in the JSON that is not accepted 

| Example 4 | Command | Observation |
|----|----|--------|
| <b>Unauthorized</b> |`curl -H "Authorization: InvalidToken" http://example.com/protected-resource` | `HTTP/1.1 401 Unauthorized` |

<b>Interpretation:</b></br>

When resources on the web are protected and can only be accessed by credentials, a `401` status code can be returned when a client is not authorized to access the resource. </br>
`401` Generally means that credentials or tokens should be passed via an HTTP header called "Authorization" </br>
`403` Generally means "Forbidden" which means the request was rejected by the web server or network device in front of the web server. Usually this is because of some restrictions when a web server only allows traffic from certain IP addresses. </br>


| Example 5 | Command | Observation |
|---|-----|--------|
| <b>Hanging request</b> |`curl http://example.com:5000` | Request is hanging |

<b>Interpretation:</b></br>

When a request is hanging, it can mean a few things. </br>
In the example above I explicitly indicated a port `5000` which I know is not open on the target server. </br>
This is not always the case, but in my experience, if a web server is sitting behind a network firewall which is usually the case in the cloud, then the firewall will let the request hang until the client times out. </br>

If you do the same `curl` command and remove the port, `curl` will use port `80` which is open and the server will respond.
In Cloud environments, virtual servers and web servers are protected by firewalls which will make the request hang until timeout. 

This is very similar to `Connection Refused` we will walk through below.

| Example 6 | Command | Observation |
|----|----|--------|
| <b>Connection Refused</b> | `curl http://localhost:5000 -v`| `Couldn't connect to server` and `Connection refused` |

<b>Interpretation:</b></br>

`Connection Refused` is a very popular error message which is generally returned by servers when a connection to a specified port is refused. </br>
This can be similar to above where I mentioned a connection is hanging, but the distinct difference is when you have a connection refused, the connection has reached a server, but the server refused it because there is no process listening on the specified port. Whereas the connection hanging can happen when a connection reaches a firewall which does not allow it through. </br>

This can also either be a client side or a server side issue. </br>
For example, a client might try to connect to the wrong port by accident or misconfiguration. </br> 
A server could be misconfigured to listen on the wrong port too, so either or needs to be verified.

In the above `curl` request, the general error was `Couldn't connect to server` which is too vague. </br>
There is no reason stated in the error message and there could be a lot of reasons. So we should know when you see such a generic error, to continue troubleshooting to work out the root cause. </br>
In this case, I use the `-v` option in `curl` to get verbose information about the HTTP request, and in the output we can see the reason is `Connection Refused` </br>
This would be overlooked if I started making assumptions when looking at the initial output. </br>

## 3. Collecting Facts

When it comes to the observations and our interpretations of it, you can now see that the answer to a root cause is not always exactly black and white and doesn't always mean a problem sits with the client. </br>
For example `Could not resolve host` means that the client could not resolve the host. This could be an issue on the client's end with internet connectivity, its DNS settings etc. In this case it could be only one client affected, and other clients would be able to connect. However, if these clients are running in a platform where a SRE team provides DNS servers, then the problem could be on a platform side. </br>

Similarly when a client times out, it could be that the server is taking too long and an issue on the server-side causing the clients to timeout. Or it could just be a misconfiguration on the client side where timeout values are set too low. </br>

This is where intuition and experience once again comes in. </br>
It's very easy to make assumptions in the above scenarios. </br>
Which brings us to this topic that is all about collecting facts. 
We want to turn our observations into facts, by looking at the observations, client errors, server errors, etc and be able to map that to evidence. 

Evidence can be found in web server logs. We learned how to run and configure a basic web server in this chapter and how to configure its logging. The web server logs will have events based on all the HTTP requests that were processed. 
The web server logs will help us further establish if the error is a client or server issue. 

### Understanding HTTP status codes and meaning of errors

When we go through the cycle of HTTP troubleshooting, from observations to collecting facts and evidence, we need to have a strong understanding of HTTP status codes. This fundamentally forms a foundation of monitoring for the web. It gives us the most clues as to what is going wrong and points us into the direction on where to look next. </br>

Let's take everything we learned so far, and recap the HTTP status codes covered in the earlier module on 
[Module: HTTP](../README.md)

## 4. Validation 

Once we have made our observations and collected our evidence, it's important to validate these. We can validate our findings using tools like `curl` and the browser's development tools etc to reproduce issues and validate our findings. We can also verify if the server has received our HTTP request, but accessing web server logs. So for example, if we see a `404` on our client when making an HTTP request, we can validate that by checking the web server access logs for a corresponding `404` event and validate that the `404` came from our server. </br>

Another example can be a `504` Gateway timeout. We see the `504` on the client side or Browser page and facts may be our output in the client like `curl`, as well as logs on the web server side. </br>
However, we want to validate our assumptions here, by checking logs on each hop, not only the client. So we would check the network devices in between such as proxy logs, or load balancing service logs. Whichever device acted as a gateway. The logs should show to `504` as well as the upstream server that was contacted and the time it took. </br>
We can then check the logs of that upstream web server and try to see if we can find the request event in the web server access logs.

![diagram](troubleshooting.drawio.svg)

# Key Takeaways 

* Practise, Practise, Practise
  * Use command line, `curl` and browser tools to test out your companies systems and services
  * Make HTTP web requests, interpret successful responses, validate and practise. 
* You will build a fundamental knowledge and experience of a working system, which helps you validate a non-working system when things go wrong
* Intuition comes with experience which comes with time. Don't wait for things to go wrong. Try to follow the above HTTP troubleshooting cycle a few times a day for practice.
* Observe web server and proxy logs as often as you can
