# 🎬 The Basics of HTML Websites and Web Services or APIs 

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 4](../../../chapters/chapter-4-web-and-http/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.

## Introduction

In previous modules, we learned about networking, TCP, and HTTP. </br>
When clients and servers communicate over HTTP, we fundamentally exchange data in the HTTP request and response `body` fields. </br>

### Types of Web services and servers
In this module we expand a little bit about the content we receive in the HTTP response `body`. This will dictate what type of HTTP server we are talking to. For example:

* <b>Web server hosting a Website</b>
  - This type of web server will generally return HTML, JavaScript, StyleSheets etc, which is used to render a web site. Usually this type of content is created by front-end or UI developers and hosted by pre built web server products like the ones mentioned below in the examples. 
  - Examples
    - NGINX
    - Apache
* <b>Web server hosting a Web API (web service)</b>
  - This type of web server will generally return "data" which can be in the form of JSON, XML and other formats. This type of service is generally used to pass data between services, like a service that returns customer records from a database, or returns weather data for given cities. 
  
  Usually these web services or APIs are created by developers.
  
  - Examples of web services or APIs
    - Python Flask
    - NodeJS express
    - ASP.NET 
    - Go HTTP    

We will deepdive into Web servers in a separate module, however it's worth mentioning that the difference between APIs or Web services and Websites are simply the content that is returned by the server as you may have noticed above. </br>

![diagram](html.drawio.svg) 

If we run an example web server that I created which you can find in this folder [here](../.test), we can navigate to that folder using `cd` and start it using `./server` in the command line terminal

<i><b>Important Note:</b> Please note that links to these files can change without notice and may be different from what you read on screen in the video at the time of recording. Please use the links in this document.</i>

Once started you should see the following output in `stdout`

```
./server 
2025/04/23 11:32:34 Starting server on :8080
```

You should now be able to access the home page of the web site at [http://localhost:8080/home.html](http://localhost:8080/home.html). 
When you access that page you should see our home page. </br>
There is quite a bit going on here at once. </br> Remember what we learned in earlier modules. 

* 1 - The browser is our client
* 2 - The server is the binary `server` we started in a terminal
* 3 - The browser will make a network request (TCP) as an HTTP request to the server and request `home.html` content.
* 4 - The browser will render the HTML page, which has two additional links in it that the browser will fetch, called `style.css` and `script.js` 
* 5 - The browser makes two additional HTTP requests to get this content.

We can see all these requests in our server output logs in `stdout` in our terminal

```
2025/04/23 11:44:51 GET /home.html [::1]:42164 HTTP/1.1 200
2025/04/23 11:44:51 GET /script.js [::1]:42180 HTTP/1.1 200
2025/04/23 11:44:51 GET /style.css [::1]:42164 HTTP/1.1 200
2025/04/23 11:44:51 GET /favicon.ico [::1]:42164 HTTP/1.1 404
```

You may now also notice that this is a great way to monitor who requests what from our web server. </br>
Web server logs give us details about who connected, when they connect, the content they asked for (HTTP Request) and the status of the HTTP response returned. </br>
By reading this output we can start painting a picture in our mind about what is going on. </br>
It's very important to build a habit of practising the skills to interpret and understand HTTP web server logs. </br>

We should also understand how the client, in our case the browser works. </br>
If you open up the browser's developer tools and open the network tab you can record all network activity that the browser makes. </br>
If you force refresh the page, you should see the same above HTTP requests being made by the browser and you can also explore the request and responses in further details.

This is sometimes really helpful when testing a website or web service to make sure all resources are requested and loaded correctly. </br>
Take some time to have a deeper look at each HTTP request, its HTTP fields, headers etc as well as HTTP response fields, headers and body content. </br>

Take note that the content in the body matches the HTTP header called `Content-Type`. Our web server returns the content that is specified by 
this header. </br>
In this example, our application can be seen as a Website hosted by a Web server. </br>

#### Web service \ API example

If we open up a different URL like [http://localhost:8080/weather](http://localhost:8080/weather), our Web server now returns JSON as its content and 
`Content-Type` header. 
In this example, we are demonstrating a Web server running a Web API that returns weather data

We can observe this activity in the browsers network tab, as well as in our web server logs too. 

It's also helpful to make HTTP requests using `curl`

```
curl -v http://localhost:8080/weather
```
We can observe the response in the command line too and get more verbose details using the `-v` option.
This allows us to craft the HTTP request and few the HTTP response fields, headers and body. </br>


## HTML files, syntax and usage

Now that we know that web servers can return different content types, I wanted to give a short overview on important aspects of the Web that DevOps engineers need to know. </br>
DevOps engineers will often work with developers including front-end or UI developers who build and design UI and web interfaces. Basically Website development. 

The "Front End" of a website is the HTML, JavaScript and CSS or Style Sheets that make up the user interface. It's called the front-end because it's literally the "front" facing part of a system that the customer interacts with. </br>

![diagram](html-frontend-backend.drawio.svg)

Front-end or UI developers will build and design the HTML using popular frameworks like React.js, Node.js </br>
These frameworks make development, debugging and certain tasks easier for developers, but ultimately it all ends up being HTML, Javascript and CSS

* `HTML` : Markup language that the browser can render
* `JavaScript` : Scripting language that allows the developer to write functionality of a website.
* `CSS` : Cascading style sheets used to apply styles like fonts, colors, widths & heights etc. 

Developers will use tools like `npm` to "compile" all their above files into a package that can be deployed. </br>

This is where DevOps engineers come in. We're often tasked to assist with building the automation that helps with compiling, packaging and deployment of these files to web servers. </br>

Although it's not expected of DevOps engineers to know the ins and outs of Web development, I would highly recommend dipping your toes into these parts of the technology.
It's always good to understand the high level overview of what your customers need and Web developers will be the customers of DevOps teams

## HTML Overview

HTML markup language is all about defining markup elements using open and closing tags. </br>

For example, you can open a DIV, with `<div>` and close it with the `</div>` tag. This allows you to place elements inside other elements for example: `<div>content inside!</div>`

#### <head>

The HTML `<head>` tag holds key information about our web page. So things like the `title`, `meta` tags and `links` to scripts and resources that are needed to render our site correctly. 

The head tag contains mostly the following tags:

* `<title>` : The title that will show on the browser tab
* `<link>` : Resources that are required to render our page, such as CSS Stylesheets and icons
* `<meta>` : Provides metadata about the page, such as character encoding, viewport settings, and descriptions for search engines.
* `<script>` : Includes JavaScript files or inline scripts for functionality.

<i>Note: When an HTML file is rendered in the browser, it's rendered from top to bottom. Therefore it's sometimes required that script tags are located at the bottom of the page, especially if they need to access elements that need to be rendered first</i>

#### <body>

The `<body>` tag holds all the content of our web page. The body is mostly made up of elements used to make our website look pretty. </br>

HTML elements we will be using:

* `<div>` : Defines a division or a section in an HTML document.
* `<h>` : Header tags, allowing you to define headings for sections of your content. They range from `<h1>` to `<h6>`, from largest to smallest. 
* `<p>` : Paragraphs, which contain blocks of text.
* `<hr>` :  Represents a thematic break between paragraph-level elements, useful to create separation between paragraphs
* `<button>` : a Button element that can perform some action like change the appearance of the site, trigger JavaScript functions etc. Buttons are a core element that allows users to interact with our site. 

## JavaScript Overview

JavaScript is a programming language that enables interactive and dynamic functionality on websites. Javascripts run in the user's browser (client-side). 
It allows developers to manipulate HTML, CSS, handle user interactions, and fetch data from API web services without having to reload the page. </br>

* `Page Manipulation` : Modify HTML and CSS dynamically (e.g., show/hide elements, update content).
* `Handle User Interactions` : Respond to user actions like clicks, keypresses, or form submissions.
* `Programming` : API calls to fetch data from web services and back-ends.
* `Cross-Browser Support` : Runs on all modern browsers without additional plugins.
* `Extensibility` : Supports libraries (e.g., React, jQuery) and frameworks (e.g., Angular, Vue) for advanced functionality.


## CSS Stylesheets Overview

CSS or "Cascading Style Sheets" is a stylesheet language used to control the presentation and layout of HTML elements on a web page. It allows developers to define styles such as colors, fonts, spacing, and positioning, enabling consistent and visually appealing designs.


* `Styling` : Apply colors, fonts, borders, and backgrounds to elements.
* `Layout` : Control the positioning of elements using techniques like Flexbox, Grid, or floats.
* `Responsive Design` : Adapt layouts for different screen sizes using media queries.
* `Reusability` : Define reusable styles with classes and IDs.
* `Separation of Concerns` : Keeps design (CSS) separate from content (HTML) for better maintainability.


## DevOps, SRE & Platform Engineering Key takeaways

As a DevOps, SRE & Platform Engineer, your customers will often be developers who may develop both front-ends and back-ends </br>
It's important to know the difference between front-end and back-end and how front-end developers build the web interfaces, how these interfaces interact with back-ends like web services or APIs

* How to host HTML\JS\CSS on a web server, which we will do in the next module
* How websites render and rely on resources loaded from web servers using the network
* How to troubleshoot network calls between the client (browser) and the server (web server)
* Understanding HTTP requests and responses and interpreting them
* Understand that these HTML\JS\CSS resources will need to be deployed to web servers so clients (browsers) can fetch them
  * The build and deployment process for this will be covered in a later chapter on CI/CD pipelines
  * DevOps engineers generally help build and support these deployment pipelines
