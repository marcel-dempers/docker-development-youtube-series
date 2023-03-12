# Introduction to Go: Storing data in Redis Database

<a href="https://youtu.be/6lJCyKwoQaQ" title="golang-part-5"><img src="https://i.ytimg.com/vi/6lJCyKwoQaQ/hqdefault.jpg" width="20%" alt="introduction to Go part 5" /></a>

Up until now, we've learned the fundamentals of Go and built a small web microservice that handles our video data.
Our service has a `/` `GET` endpoint for returning all videos, as well as a simple `/update` endpoint for updating our list of videos.

We've learnt how to read and write from files and learn how to work with `json` data. </br>
This is important for learning Go, however there are a few challenges for using a file as storage. <br/>

* It can be problematic if we have more than one instance of our service writing to the same file
* It's important to keep state out of our application, so if our application crashes, we don't lose our data

[In part 1](../readme.md), we covered the fundamentals of writing basic Go <br/>
[In part 2](../part-2.json/readme.md), we've learnt how to use basic data structures like `json` so we can send\receive data. <br/>
[In part 3](../part-3.http/readme.md), we've learnt how to write a HTTP service to expose our videos data <br/>

## Start up a Redis Cluster

Follow my Redis clustering Tutorial </br>

<a href="https://youtube.com/playlist?list=PLHq1uqvAteVtlgFkmOlIqWro3XP26y_oW" title="Redis"><img src="https://i.ytimg.com/vi/L3zp347cWNw/hqdefault.jpg" width="30%" alt="Redis Guide" /></a>

Code is over [here](../../../storage/redis/clustering/readme.md)

## Go Dev Environment

The same as Part 1+2+3+4, we start with a [dockerfile](./dockerfile) where we declare our version of `go`.

The `dockerfile`:

```
FROM golang:1.15-alpine as dev

WORKDIR /work
```

Let's build and start our container: 

```
cd golang\introduction\part-5.database.redis

docker build --target dev . -t go
docker run -it -p 80:80 -v ${PWD}:/work go sh
go version
```

## Our application

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

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func main() {
	
	http.HandleFunc("/", HandleGetVideos)
	http.HandleFunc("/update", HandleUpdateVideos)

	http.ListenAndServe(":80", nil)
}
```

Now before we write these handler functions, we need our videos application

## Create our Videos app

Firstly, we create a separate code file `videos.go` that deals with our YouTube videos <br/>

The `videos.go` file defines what a video `struct` looks like, a `getVideos()` function to retrieve <br/>
videos list as a slice and a `saveVideos()` function to save videos to a file locally. <br/>

We start with our dependencies. <br> 
We import 2 packages, 1 for reading and writing files, and another for dealing with `json`

```
package main

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

## HTTP Handlers

Now we have to define our handler functions `HandleGetVideos` and `HandleUpdateVideos`, 

```
func HandleGetVideos(w http.ResponseWriter, r *http.Request){
	
	videos := getVideos()
	videoBytes, err  := json.Marshal(videos)

	if err != nil {
  	panic(err)
	}

	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var videos []video
			err = json.Unmarshal(body, &videos)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Bad request")
			}

			saveVideos(videos)

		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method not Supported!")
		}
}

```

Now so far, we have a web HTTP application that can list and update our youtube videos. <br/>
Let's build and run it 

```
go build

./videos

```

If we head over to `http://localhost` in the browser we can see our application returns our 2 videos. </br>.
We can use tools like PostMan to generate a `POST` request to update our videos too.


## Redis Go Package 

Now instead of reading and writing our records to a `json` file, we are going to read and write records to Redis.

https://github.com/go-redis/redis/

```
go get github.com/go-redis/redis/v8

```

And to use the library, we have to import it

```
import (
  "github.com/go-redis/redis/v8"
)
```

Now to connect to any database, you're going to need a bit of information:

* Database Address
* Database Port
* Database Username\Password

We can define these as environment variables and read those in our code

```
import (
  "os"
)
//main.go (global variables)

var redis_sentinels = os.Getenv("REDIS_SENTINELS")
var redis_master = os.Getenv("REDIS_MASTER_NAME")
var redis_password = os.Getenv("REDIS_PASSWORD")

```

We define an empty context and Redis client
Context can be used to control timeouts and deadlines for our application. We can set up an empty context for now.

```
import (
  "context"
)

var ctx = context.Background() 
var redisClient *redis.Client

```

## Redis Sentinel Client

https://redis.uptrace.dev/guide/sentinel.html#redis-server-client


```
import (
  "strings"
)

sentinelAddrs := strings.Split(redis_sentinels, ",")
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
  MasterName:     redis_master,
  SentinelAddrs:  sentinelAddrs,
  Password: redis_password,
  DB: 0,
})

redisClient = rdb

rdb.Ping(ctx)
```

We can also add the `Ping()` to our handler functions to ensure it can connect using the global redis client variable

## Networking

Now we need to remember our go container may not be able to talk to the redis containers because they are running on different networks.

If you took note, we started our Redis containers on a `redis` network by passing `--net redis` as a flag to our `docker run` commands.  </br>

Let's restart our Go container on the same network

We also need to set our `ENV` variables to point our container to the redis sentinels. </br>

If we look at our sentinel configuration, our master alias is set to `mymaster`

```
docker run -it -p 80:80 `
  --net redis `
  -e REDIS_SENTINELS="sentinel-0:5000,sentinel-1:5000,sentinel-2:5000" `
  -e REDIS_MASTER_NAME="mymaster" `
  -e REDIS_PASSWORD="a-very-complex-password-here" `
  -v ${PWD}:/work go sh
```

We can now observe our container is connected to Redis. </br>
Our application: [http://localhost](http://localhost)

# Store our Data

So now we can store our video data in Redis instead of a local `json` file. </br>
We'll write the `json` document of a video (struct) to Redis. </br>

Let's create a `saveVideo()` function that stores a single record.

```
func saveVideo(video video)(){

  videoBytes, err  := json.Marshal(video)
  if err != nil {
		panic(err)
	}

	err = redisClient.Set(ctx, video.Id, videoBytes, 0).Err()
	if err != nil {
		panic(err)
	}
}
```

Now since we have one endpoint that saves all videos, we need to update it to save each video it has.

```
func saveVideos(videos []video)(){
	for _, video := range videos {
		saveVideo(video)
	}
}
```

To get the videos, let's create a function to retrieve a single record:

```
func getVideo(id string)(video video) {
	
	value, err := redisClient.Get(ctx, id).Result()

	if err != nil {
		panic(err)
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &video)
	}
	
	return video
}

```

And finally we need to update our `GetVideos()` function to loop all the video keys and return all videos

```
func getVideos()(videos []video){
	
	keys, err := redisClient.Keys(ctx,"*").Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		video := getVideo(key)
		videos = append(videos, video)
	}

	return videos
}

```

And we need all our imports

```
import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
)
```

Now if we rebuild our all and access it, we get `null` as there are no videos in Redis. Let's add two!

Let's `POST` the following using PostMan to url `http://localhost/update`
```
[
	{
		"id" : "QThadS3Soig",
		"title" : "Kubernetes on Amazon",
		"imageurl" : "https://i.ytimg.com/vi/QThadS3Soig/sddefault.jpg",
		"url" : "https://youtu.be/QThadS3Soig",
		"description" : "TEST"
	},
	{
		"id" : "eyvLwK5C2dw",
		"title" : "Kubernetes on Azure",
		"imageurl" : "https://i.ytimg.com/vi/eyvLwK5C2dw/mqdefault.jpg?sqp=CISC_PoF&rs=AOn4CLDo7kizrJozB0pxBhxL9JbyiW_EPw",
		"url" : "https://youtu.be/eyvLwK5C2dw",
		"description" : "TEST"
	}
]

```

If we refresh our page, we can now see two records!

# Improvements

Now that you have the fundamental knowledge of HTTP and Redis, </br>
you can update the code to retrieve 1 video by ID, or delete a video by ID. </br>
You can add search functionality and more! </br>

Let's update our `/ GET` handler to be able to return a single video

```
func HandleGetVideos(w http.ResponseWriter, r *http.Request){
	
	id, ok := r.URL.Query()["id"]
  
	if ok {

    videoID := id[0]
    video := getVideo(videoID)
    
		if video.Id == "" { //video not found, or empty ID
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{}"))
			return 
		}

		videoBytes, err  := json.Marshal(video)
		if err != nil {
  		panic(err)
		}

		w.Write(videoBytes)
    return
		
	}

	videos := getVideos()
	videoBytes, err  := json.Marshal(videos)

	if err != nil {
  	panic(err)
	}

	w.Write(videoBytes)
}
```

We can also update our `/update POST` endpoint to be able to update a single video

```
func HandleUpdateVideos(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			_, ok := r.URL.Query()["id"]
			if ok {
				
				var video video
				err = json.Unmarshal(body, &video)
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, "Bad request")
				}

				saveVideo(video)
				return

			}
			
			var videos []video
			err = json.Unmarshal(body, &videos)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Bad request")
			}

			saveVideos(videos)
			return

		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method not Supported!")
		}
}
```

# Build our Docker container

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
CMD ./videos

```

Build :
```
cd golang\introduction\part-5.database.redis
docker build . -t videos
```

Run :
```
docker run -it -p 80:80 `
  --net redis `
  -e REDIS_SENTINELS="sentinel-0:5000,sentinel-1:5000,sentinel-2:5000" `
  -e REDIS_MASTER_NAME="mymaster" `
  -e REDIS_PASSWORD="a-very-complex-password-here" `
  videos
```
