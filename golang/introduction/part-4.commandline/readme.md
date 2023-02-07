# Introduction to Go: Command Line

<a href="https://youtu.be/CODqM_rzwtk" title="golang-part-4"><img src="https://i.ytimg.com/vi/CODqM_rzwtk/hqdefault.jpg" width="20%" alt="introduction to Go part 4" /></a>

Command line apps are a fundamental part of software development <br/>

Go has a built in Commandline parser package. The package can be found [here](https://golang.org/pkg/flag/) <br/>
We simply have to import the `flag` package: 

```
import (
  "flag"
)
```

[In part 1](../readme.md), we covered the fundamentals of writing basic Go <br/>
[In part 2](../part-2.json/readme.md), we've learn how to use basic data structures like `json` so we can send\receive data. <br/>
[Part 3](../part-3.http/readme.md) was about exposing data via a Web server.

We'll be combining these techniques so we can serve our `videos` data over a commandline application.

As always, let's start with our `dockerfile` , `main.go` and `videos.go` we created in Part 2

## Dev Environment

The same as Part 1+2+3, we start with a [dockerfile](./dockerfile) where we declare our version of `go`.

```
cd golang\introduction\part-4.commandline

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version
```

## Create our App

Create a new directory that holds defines our `repository` and holds our `module`

```
mkdir videos
cd videos
```

* Define a module path

```
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

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
  if err != nil {
		panic(err)
	}

}
```

## Flag Package 

https://golang.org/pkg/flag/

Package flag implements command-line flag parsing.
So we can run our videos app and pass it inputs like any other CLI.

Let's build a CLI tool that users our `getVideos()` and `saveVideos` functions.

To get all videos, perhaps we'd like a command
```
# get all videos
videos get --all

# get video by ID
videos get -id <video-id>

# add a video to our list
videos add -id -title -url -imageurl -desc 


```
To start, we import package `flag`

```
import (
	"flag"
	"fmt"
)
```


Let's define our subcommands in `main.go` :

```
    //'videos get' subcommand
    getCmd := flag.NewFlagSet("get", flag.ExitOnError)
```

`videos get` command will need two inputs, `--all` if the user wants to return all videos, or `--id` if the user only wants a specific video

```
  // inputs for `videos get` command
  getAll := getCmd.Bool("all", false, "Get all videos")
  getID := getCmd.String("id", "", "YouTube video ID")
```

`videos add` command will take a bit more inputs to be able to add a video to our storage.
```
  addCmd := flag.NewFlagSet("add", flag.ExitOnError)

  addID := addCmd.String("id", "", "YouTube video ID")
  addTitle := addCmd.String("title", "", "YouTube video Title")
  addUrl := addCmd.String("url", "", "YouTube video URL")
  addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
  addDesc := addCmd.String("desc", "", "YouTube video description")


```

When a user runs our videos CLI tool, we may need to validate that 
our application receives the right subcommands. So lets ensure a simple validation to check if the user has passed a subcommand

To check the arguments passed to our CLI, we use the ["os"](https://golang.org/pkg/os/) package. Check the Args variable, it holds usefull information passed to our application including its name.
`var Args []string`

We can do a simple check by ensuring the length of `os.Args` is atleast 2.

We firstly need to add `os` to our import section </br>
Followed by our check:

```
if len(os.Args) < 2 {
        fmt.Println("expected 'get' or 'add' subcommands")
        os.Exit(1)
  }
```

## Handling our subcommands

So to handle each sub command like `get` and `add`, we add a simple
`switch` statement that can branch into different pathways of execution,
based on a variables content.

Let's take a look at this simple `switch` statement:

```
//look at the 2nd argument's value
switch os.Args[1] {
  case "get": // if its the 'get' command
    //hande get here
  case "add": // if its the 'add' command
    //hande add here
  default: // if we don't understand the input
}
```

Let's create seperate handler functions for each sub command to keep our code tidy:

```
func HandleGet(getCmd *flag.FlagSet, all *bool, id *string){
}

func HandleAdd(addCmd *flag.FlagSet,id *string, title *string, url *string, imageUrl *string, description *string ){
}
```

Now that we have seperate functions for each subcommand, we can take appropriate actions in each. Let's firstly parse the command flags for each subcommand:

This allows us to parse everything after the `videos <subcommand>` arguments:

```
getCmd.Parse(os.Args[2:])
```

## Input Validation

For our `HandleGet` function, let's validate input to ensure its correct.

```
  if *all == false && *id == "" {
    fmt.Print("id is required or specify --all for all videos")
    getCmd.PrintDefaults()
    os.Exit(1)
  }
```

Let's handle the scenario if user passing `--all` flag:

```
if *all {
		//return all videos
		videos := getVideos()
		
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n",video.Id, video.Title, video.Url, video.Imageurl,video.Description)
		}

		return
}
```

Let's handle when user is searching for a video by ID

```
if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n",video.Id, video.Title, video.Url, video.Imageurl,video.Description)

			}
		}
} 
```

## Parsing multiple fields

For our `HandleAdd` function, we need to validate multiple inputs, create a `video` struct, append it to the existing video list and save it back to file

Let's create a `ValidateVideo()` function with similar inputs to our `HandleAdd()`:

```
func ValidateVideo(addCmd *flag.FlagSet,id *string, title *string, url *string, imageUrl *string, description *string ){
}

```

Let's simply validate all fields since they are all required:

```
if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
```

And we can now call this function in our add function:

```
ValidateVideo(addCmd, id,title,url, imageUrl, description)
```

## Adding our video

Now that we have some basic validation, not perfect, but good enough to get started, let's add our video to the existing file.

Define a video struct with the CLI input:

```
	video := video{
		Id: *id,
		Title: *title,
		Description: *description,
		Imageurl: *imageUrl, 
		Url: *url, 
	}
```

Get the existing videos:

```
videos := getVideos()
```

Append our video to the list:

```
videos = append(videos,video)
```

Save the new updated video list:

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
COPY --from=build /videos/videos /usr/local/bin/videos
COPY ./videos/videos.json /
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]
```

For our entrypoint, we need to create a shell script to accept all the arguments:

Let's create a script called `run.sh`

```
#!/bin/sh

videos $@

```

Build :
```
cd golang\introduction\part-4.commandline
docker build . -t videos
```

Run :
```
docker run -it videos get --help
```