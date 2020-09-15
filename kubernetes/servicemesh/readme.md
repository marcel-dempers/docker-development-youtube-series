# Introduction to Service Mesh

To understand service mesh, we need a good use case. <br/>
We need some service-to-service communication. <br/>
A basic microservice architecture will do. <br/>

## A simple Web UI: videos-web
<hr/>
<br/>

Consider `videos-web` <br/>
It's an HTML application that lists a bunch of playlists with videos in them.

```
+------------+
| videos-web |
|            |
+------------+
```
<br/>

## A simple API: playlists-api
<hr/>
<br/>

For `videos-web` to get any content, it needs to make a call to `playlists-api`

```
+------------+     +---------------+
| videos-web +---->+ playlists-api |
|            |     |               |
+------------+     +---------------+

```

Playlists consist of data like `title`, `description` etc, and a list of `videos`. <br/>
Playlists are stored in a database. <br/>
`playlists-api` stores its data in a database

```
+------------+     +---------------+    +--------------+
| videos-web +---->+ playlists-api +--->+ playlists-db |
|            |     |               |    |              |
+------------+     +---------------+    +--------------+

```

<br/>

## A little complexity
<hr/>
<br/>

Each playlist item contains only a list of video id's. <br/>
A playlist does not have the full metadata of each video. <br/>

Example `playlist`:
```
{
  "id" : "playlist-01",
  "title": "Cool playlist",
  "videos" : [ "video-1", "video-x" , "video-b"]
}
```
Take not above `videos: []` is a list of video id's <br/>

Videos have their own `title` and `description` and other metadata. <br/>

To get this data, we need a `videos-api` <br/>
This `videos-api` has its own database too <br/>

```
+------------+       +-----------+
| videos-api +------>+ videos-db |
|            |       |           |
+------------+       +-----------+
```

For the `playlists-api` to load all the video data, it needs to call `videos-api` for each video ID it has.<br/>
<br/>

## Traffic flow
<hr/>
<br/>
A single `GET` request to the `playlists-api` will get all the playlists 
from its database with a single DB call <br/>

For every playlist and every video in each list, a separate `GET` call will be made to the `videos-api` which will
retrieve the video metadata from its database. <br/>

This will result in many network fanouts between `playlists-api` and `videos-api` and many call to its database. <br/>
This is intentional to demonstrate a busy network.

<br/>

## Full architecture
<hr/>
<br/>

```

+------------+     +---------------+    +--------------+
| videos-web +---->+ playlists-api +--->+ playlists-db |
|            |     |               |    |              |
+------------+     +-----+---------+    +--------------+
                         |
                         v
                   +-----+------+       +-----------+
                   | videos-api +------>+ videos-db |
                   |            |       |           |
                   +------------+       +-----------+

```

## Run the apps

There is a `docker-compose.yaml`  in this directory. <br/>
Change your terminal to this folder and run:

```
docker-compose build

docker-compose up

```

You can access the app on `http://localhost` 