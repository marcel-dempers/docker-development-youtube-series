# Introduction to Go: HTTP

<a href="https://youtu.be/MKkokYpGyTU" title="golang-part-3"><img src="https://i.ytimg.com/vi/MKkokYpGyTU/hqdefault.jpg" width="20%" alt="introduction to Go part 3" /></a>

HTTP is a fundamental part of Microservices and Web distributed systems <br/>

Go has a built in HTTP web server package. The package can be found [here](https://golang.org/pkg/net/http/) <br/>
We simply have to import the `http` package: 

```
import (
  "net/http"
)
```

[In part 1](../readme.md), we covered the fundamentals of writing basic Go <br/>
[In part 2](../part-2.json/readme.md), we've learn how to use basic data structures like `json` so we can send\receive data. <br/>

We'll be combining both these techniques so we can serve our `videos` data over a web endpoint.

As always, let's start with our `dockerfile` , `main.go` and `videos.go` we created in Part 2


## Dev Environment

The same as Part 1+2, we start with a [dockerfile](./dockerfile) where we declare our version of `go`.

```
cd golang\introduction\part-3.http

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```

## Create our App

Create a new directory that holds defines our `repository` and holds our `module`

```
mkdir videos

```

* Define a module path

```
# change directory to your application source code

cd videos

# create a go module file

go mod init videos

```

## Create our base code

We start out all our applications with a `main.go` defining our `package`, declaring our `import` dependencies <br/>
and our entrypoint `main()` function

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

## Create our Videos app

Firstly, we create a seperate code file `videos.go` that deals with our YouTube videos <br/>

The `videos.go` file defines what a video `struct` looks like, a `getVideos()` function to retrieve <br/>
videos list as a slice and a `saveVideos()` function to save videos to a file locally. <br/>

Let's copy the following content from Part 2 and create `videos.go` :

We want `videos.go` to be part of package main:

```
package main

```

We import 2 packages, 1 for reading and writing files, and another for dealing with `json`

```
import (
	"io/ioutil"
	"encoding/json"
)

```

Then we define what a video `struct` looks like:

```

type video struct {
	Id string  
	Title string  
	Description string 
	Imageurl string 
	Url string 
}

```

We have a function for retrieving `video` objects as a list of type `slice` :

```

func getVideos()(videos []video){
	
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

```

We also need to copy our `videos.json` file which contains our video data. <br/>

And finally, we have a function that accepts a list of type `slice` and stores the videos to a local file

```
func saveVideos(videos []video)(){

  videoBytes, err  := json.Marshal(videos)
  if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
  if err != nil {
		panic(err)
	}

}
```

## HTTP Package 

https://golang.org/pkg/net/http/

The HTTP package allows us to implement an HTTP client and a server.
A client is a component that makes HTTP calls.
A server is a component that receives or serves HTTP.

The HTTP package is capable of sending HTTP requests as well as defining a server
for receiving HTTP requests. 

We can use this to run an HTTP server to serve files, or serve data, like an API.

Let's define a server in `main.go` :

```
# just one line :)

http.ListenAndServe(":8080", nil)

# ListenAndServe starts an HTTP server with a given address and handler. 
# The handler is usually nil, which means to use DefaultServeMux. 
# Handle and HandleFunc add handlers to DefaultServeMux

```

Now before we run this, since we're running in Docker, we want to exit the container <br/>
and rerun it, but this time open port `8080`

```
docker run -it -p 8080:8080 -v ${PWD}:/work go sh
cd videos

go run main.go 
# you will notice the application pausing
```

We should see our server with a 404 on http://localhost:8080/

## Handle HTTP requests

In order to handle requests, we can tell the HTTP service that we want it to run a function </br>
for the request coming in.

We can see the `http` package has a `HandleFunc` function: https://golang.org/pkg/net/http/

To see this in action, lets create a `Hello()` function:

```
func Hello(){
}
```

And tell our `http` service to run it:

```
http.HandleFunc("/", Hello)
```

We cannot run this yet. As per `http` documentation, our `Hello` function needs to take in some inputs.
`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`

Therefore we need to add inputs to our function: 

```
func Hello(w http.ResponseWriter, r *http.Request){
}
```

This allows us to get the request, its `body`, `headers` and a write where we can send a response.
Run this in the browser and you will notice the 404 goes away, but we now get an empty response.

## HTTP Response

Let's write a reponse to the incoming request. </br>
The response write has a `Write()` function that takes a bunch of bytes. <br/>
We can convert string to bytes by casting a `string` to a `[]byte` <br/> like:
`[]byte("Hello!")`. Let's convert it and write "Hello" to the response:

```
w.Write([]byte("Hello!"))
```

IF we run this code, we can see "Hello!" in the response body

## HTTP Headers

Headers play an important role in HTTP communication. </br>
Lets access all the headers of the incoming request!
If we look at the Header definition [here](https://golang.org/pkg/net/http), we can see how to access it.
Let's use the `for` loop we learnt in [part 1](../readme.md)

```
for i, value := range r.Header {
}
```

We learn't from our loop, we have in indexer and a value.
For `i`, we can rename it to header since it represents the header key in the dictionary.
And the `value` is the value of type `[]string`, containing the value of the header:

```
for header, value := range r.Header {
  fmt.Printf("Key: %v \t Value: %v \n", header, value)
}
```

We can use `fmt` to print out the values and look at the headers.

We can also set headers on our response. <br/>
If we take a look at the `http` docs, we can see header is also a dictionary or strings.

```
w.Header().Add("TestHeader", "TestValue")
```

You can now see the headers in the response value if you use `curl` or your browser development tools

## HTTP Methods | GET

Web servers can serve data in a number of ways and support multiple type of HTTP methods.
`GET` is used to request data from a specified resource.
So far, our HTTP route for our Hello function is using the `GET` method.

Let's make our `GET` method more useful by serving our video data </br>
Let's rename our `Hello()` function to `HandleGetVideos()`. </br>
Our `/` route will return all videos:

```
videos := getVideos()
```

In [part 2](../part-2.json/readme.md) we covered `JSON`.
We need to convert our video `slice` of `struct`'s to `JSON` in order to return it to the client.

For this we learnt about the Marshall function: 

Import the `JSON` package: 
```
"encoding/json"
```

Convert our videos to `JSON` :

```
videoBytes, err  := json.Marshal(videos)

if err != nil {
  panic(err)
}

w.Write(videoBytes)

```

If we run this code and hit our `/` endpoint, we can now see `JSON` data being returned. <br/>
This is a core part of building an API in Go. <br/>

## HTTP Methods | POST

A `POST` method is used to send data to a server to create/update a resource. <br/>
Since we built a `saveVideos` function, lets use that so a client can update videos! <br/>

We need to define a new route, we can all it `/update` :

```
http.HandleFunc("/update", HandleUpdateVideos)
```

And we need to define an `HandleUpdateVideos()` function:

```
func HandleUpdateVideos(w http.ResponseWriter, r *http.Request){
}
```

Let's validate the request method to ensure its `POST`
We need to also ensure we send a status code to inform the user of method not allowed.

https://golang.org/pkg/net/http/#ResponseWriter

```
	if r.Method == "POST" {
		//update our videos here!
	} else {
	  w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported!")
	}
```

Now we need to accept `JSON` from the `POST` request body
https://golang.org/pkg/net/http/#Request

In the docs above, we can see the request Body is of type `Body io.ReadCloser`
To read that, we can use the `ioutil` package 
https://golang.org/pkg/io/ioutil/

```
import "io/ioutil"
```

Then we can read the body into a `slice` of `bytes`:

```
body, err := ioutil.ReadAll(r.Body)

if err != nil {
  panic(err)
}
```

Now that we have the body in a `[]byte`, we need to use our knowledge from [part 2](../part-2.json/readme.md) where we <br/>
convert `[]byte` to a `slice` of `video` items.

```

var videos []video
err = json.Unmarshal(body, &videos)
if err != nil {
  panic(err)
}
```

Creating our video objects allows us to do some validation if we wanted to. <br/>
We can ensure the request body adheres to our API contract for this videos API. <br/>
So instead of calling `panic`, lets return a `400` Bad request status code if we cannot <br/>
Unmarshal the `JSON` data. This might help with some basic validation.

```
w.WriteHeader(400)
fmt.Fprintf(w, "Bad request")
```

And Finally, let's update our videos file! :

```
saveVideos(videos)
```

# Build our Docker container

Let's uncomment all the build lines in the `dockerfile`
Full `dockerfile` : 

```
FROM golang:1.15-alpine as dev

WORKDIR /work

FROM golang:1.15-alpine as build

WORKDIR /videos
COPY ./videos/* /videos/
RUN go build -o videos


FROM alpine as runtime 
COPY --from=build /videos/videos /
COPY ./videos/videos.json /
CMD ./videos

```

Build :
```
cd golang\introduction\part-3.http
docker build . -t videos
```

Run :
```
docker run -it -p 8080:8080 videos
```

## Things to know

* SSL for secure web connection
* Authentication
* Good API validation
* Support a backwards compatible contract (Inputs remain consistent)
