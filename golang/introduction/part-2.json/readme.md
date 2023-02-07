# Introduction to Go: JSON

<a href="https://youtu.be/_ok29xwZ11k" title="golang-part-2"><img src="https://i.ytimg.com/vi/_ok29xwZ11k/hqdefault.jpg" width="20%" alt="introduction to Go part 2" /></a>

In programming languages, you will very often deal with data structures internally. <br/>
Sometimes, you need to pass data outside of your application or read data from another application, or even a file. <br/>

API, for example often expose data in `json`, `xml`, `grpc` and all sorts of formats.
Before we can really take a look at building APIs and even storing data in databases,
we need some fundamental knowledge about how to convert these data structures, into structures
that our application can understand. <br/>

In part 1, we dealt with [Variables]("https://tour.golang.org/basics/8") and more importantly, <br/> we dealt with `struct` data type, which made us build a `customer` object.

If we wanted to pass the `customer` to a database, or an external system, it's often required that we convert this to `json` format.


## Dev Environment

The same as Part 1, we start with a [dockerfile](./dockerfile) where we declare our version of `go`.

```
cd golang\introduction\part-2.json

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

Similar to Part 1, we define the file, we define what a video looks like using a `struct` <br/>
and a function for returning a slice of videos.


```
package main

import ()

type video struct {
	Id string  
	Title string  
	Description string 
	Imageurl string 
	Url string 
}

func getVideos()(videos []video){
	//Get our videos here,
	//and return them
	return videos
}


```

## Populating our Video Struct

Similar to our customers app, we created a `struct` and populated it with data: 
Let's add the following to our `getVideos()` function:

```
video1 := video{
  Id : "eyvLwK5C2dw",
  Title : "Kubernetes on Azure",
  Imageurl : "https://i.ytimg.com/vi/eyvLwK5C2dw/mqdefault.jpg?sqp=CISC_PoF&rs=AOn4CLDo7kizrJozB0pxBhxL9JbyiW_EPw",
  Url : "https://youtu.be/eyvLwK5C2dw",
  Description : "",
}

video2 := video{
  Id : "QThadS3Soig",
  Title : "Kubernetes on Amazon",
  Imageurl : "https://i.ytimg.com/vi/QThadS3Soig/sddefault.jpg",
  Url : "https://youtu.be/QThadS3Soig",
  Description : "",
}

return []video{ video1, video2}

```

We need to also invoke our function that will return the videos
Don't worry about where the videos will come from. <br/>
We always start with the building blocks and move on from there <br/>

In our `main()` function:  <br/>

```
func main() {
	videos := getVideos()
	fmt.Println(videos)
}
```

## Files

Now, Ideally, we do not want to "hard code" our videos like this. <br/>
Currently, our videos are embedded into the code and we can only return 2 videos. <br/>
If we want to introduce a new video, we have to rebuild the application. <br/>

In the real world, videos are our data. <br/>
Data lives outside of the application. Like in a database.

Technically, we can use a database, but that is too big of a step for now. <br/>
So let's start with a file instead.

### Introducing ioutil

It's important to learn how to navigate and read go documentation.
Let's go to https://golang.org/pkg/io/ioutil/

We can import this library since its part of the Go standard library set.

```
import (
	"io/ioutil"
)
```

Let's move our videos into a `json` file called `videos.json` :

```
[
  {
		"id" : "QThadS3Soig",
		"title" : "Kubernetes on Amazon",
		"imageurl" : "https://i.ytimg.com/vi/QThadS3Soig/sddefault.jpg",
		"url" : "https://youtu.be/QThadS3Soig",
		"description" : ""
	},
  {
		"id" : "eyvLwK5C2dw",
		"title" : "Kubernetes on Azure",
		"imageurl" : "https://i.ytimg.com/vi/eyvLwK5C2dw/mqdefault.jpg?sqp=CISC_PoF&rs=AOn4CLDo7kizrJozB0pxBhxL9JbyiW_EPw",
		"url" : "https://youtu.be/eyvLwK5C2dw",
		"description" : ""
	}
]

```

## Reading a file

Let's take another look at https://golang.org/pkg/io/ioutil/ <br/>
Notice the `ReadFile` function
We can read a file from disk:

```
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

  fileContent := string(fileBytes)
  fmt.Println(fileContent)
  
```

## Panic

Also notice the function `panic`. <br/>

"A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully." 
- https://gobyexample.com/panic

For now, we will panic on every potential error. <br/>
In a future video, we'll cover Error handling in more depth

## JSON 

Working with JSON is pretty straightforward in go. <br/>
`struct` objects are pretty well synergized with `json`

Let's take a look at the `json` package that is part of go.
https://golang.org/pkg/encoding/json/

```
"encoding/json"
```

This package allows us to marshal and unmarshal JSON. <br/>
This allows conversion between a `struct` or a go type, and `json` 


## Unmarshal

From JSON(bytes) ==> Struct\Go type

```
  err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
```
## Using our Data

Now in the real world, you would use your data to do something. <br/>
Let's say we'd like to update some common terms and conditions to the video description <br/>

In our `main` function, let's loop the videos & update the description.
Feel free to checkout Part 1 of the series on `loops`!

```
  for _, video := range videos {
  }
```

Now we dont want to override the description, we want to inject the terms and conditions on a new line

```
  video.Description = video.Description + "\n Remember to Like & Subscribe!"
```

Note that when we run this, it does not print our new description field </br>
This is because the loop `range` is giving us a copy of the video object. </br> 
This means we are updating a copy, but printing out the original. </br>
We need to use the loop indexer and update the original object in the slice. </br>


```
  for i, _ := range videos {
		videos[i].Description = videos[i].Description + "\n Remember to Like & Subscribe!"
	}
```

Run our app again and now it updates the original video!

## Marshal

https://golang.org/pkg/encoding/json/#Marshal
From Struct\Go type ==> JSON(bytes) 

Let's create a new function for saving our video back to file

```
func saveVideos(videos []video)(){

  videoBytes, err  := json.Marshal(videos)
  if err != nil {
		panic(err)
	}

}

```

## Writing to File

https://golang.org/pkg/io/ioutil/

```

  err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
  if err != nil {
		panic(err)
	}
```

Then in our `main` function we can save our videos after the update:

```
  saveVideos(videos)
```

## Mapping fields


In our example our `json` and our struct fields match exactly.
Sometimes this is not possible to maintain.
Sometimes we need to tell go, to map certain `json` fields to certain fields in our `struct`
We can do it like so:

```
type video struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageurl string `json:"imageurl"`
	Url string `json:"url"`
}
```