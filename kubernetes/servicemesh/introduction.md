# An Introduction to Service Mesh

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

## Full application architecture
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

## Adding an Ingress Controller

Adding an ingress controller allows us to route all our traffic. </br>
We setup a `host` file with entry `127.0.0.1  servicemesh.demo`
And `port-forward` to the `ingress-controller`


```
servicemesh.demo/home --> videos-web
servicemesh.demo/api/playlists --> playlists-api


                              servicemesh.demo/home/           +--------------+
                              +------------------------------> | videos-web   |
                              |                                |              |
servicemesh.demo/home/ +------+------------+                   +--------------+
   +------------------>+ingress-nginx      |
                       |Ingress controller |
                       +------+------------+                   +---------------+    +--------------+
                              |                                | playlists-api +--->+ playlists-db |
                              +------------------------------> |               |    |              |
                              servicemesh.demo/api/playlists   +-----+---------+    +--------------+
                                                                     |
                                                                     v
                                                               +-----+------+       +-----------+
                                                               | videos-api +------>+ videos-db |
                                                               |            |       |           |
                                                               +------------+       +-----------+



```
<br/>

## Run the apps: Docker
<hr/>
<br/>
There is a `docker-compose.yaml`  in this directory. <br/>
Change your terminal to this folder and run:

```
docker-compose build

docker-compose up

```

You can access the app on `http://localhost` 

<br/>

## Run the apps: Kubernetes 
<hr/>
<br/>

Create a cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name servicemesh --image kindest/node:v1.18.4
```
<br/>

### Deploy videos-web

<hr/>
<br/>

```
cd ./kubernetes/servicemesh/

kubectl apply -f applications/videos-web/deploy.yaml
kubectl port-forward svc/videos-web 80:80

```

You should see blank page at `http://localhost/` <br/>
It's blank because it needs the `playlists-api` to get data

<br/>

### Deploy playlists-api and database

<hr/>
<br/>

```
cd ./kubernetes/servicemesh/

kubectl apply -f applications/playlists-api/deploy.yaml
kubectl apply -f applications/playlists-db/
kubectl port-forward svc/playlists-api 81:80

```

You should see empty playlists page at `http://localhost/` <br/>
Playlists are empty because it needs the `videos-api` to get video data <br/>

<br/>

### Deploy videos-api and database

<hr/>
<br/>

```
cd ./kubernetes/servicemesh/

kubectl apply -f applications/videos-api/deploy.yaml
kubectl apply -f applications/videos-db/
```

Refresh page at `http://localhost/` <br/>
You should now see the complete architecture in the browser <br/>