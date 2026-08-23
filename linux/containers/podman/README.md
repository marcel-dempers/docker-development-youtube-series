# Introduction to Podman

<!-- #TODO: 
<a href="https://youtu.be/xxxxxx" title="topic"><img src="https://i.ytimg.com/vi/VIDEO_ID/hqdefault.jpg" width="40%" alt="topic" /></a>
 -->

In this guide, we will dive into Linux containers with Podman and explore the full container lifecycle. </br>
We will cover what Podman is, how it differs from Docker, how to install it, and how to build and run containers with it. </br>

## What is Podman

[Podman](https://podman.io/) is a daemonless, open source, Linux native tool for working with Linux containers </br>

* Free and open source
* Fast, lightweight, and daemonless
* Can run Kubernetes pods directly
* Provides a Docker-compatible CLI (alternative to `docker`)

<i>Note: Podman does not require a background daemon. It uses a fork-exec model directly with the Linux kernel (via `runc`), which makes it easier to run rootless and more secure by default.</i>

Podman also supports "Pods", groups of containers similar to Kubernetes pods, and can generate Kubernetes YAML directly from running containers. </br>


[Documentation](https://docs.podman.io/)

## Installing Podman

Podman runs natively on Linux, and also works well inside WSL if you're on Windows. </br>
Running it on native Linux (or a Linux VM/WSL) gives the best portability and closest match to how it runs in production. </br>

<i>Note: Podman Desktop is a separate GUI tool for managing Podman on Windows/macOS. We'll cover it in a dedicated video.</i>

```shell
# Debian/Ubuntu
sudo apt-get update
sudo apt-get install -y podman
```

## Basic Commands

Podman's CLI mirrors Docker's, so most commands feel familiar right away:

```shell
podman --help
```

## Building a Container Image

Podman builds standard OCI images, so any existing `Dockerfile` works without changes. </br>
We'll quickly build images across a few different languages, C#, Node.js, Python, and Go, to show that the same Dockerfile format applies regardless of stack.

```shell
cd python
podman build -t aimvector/python:1.0.0

# note that podman is case sensitive and looking for Dockerfile instread of dockerfile .
podman build -f dockerfile -t aimvector/python:1.0.0

cd nodejs
podman build -f dockerfile -t aimvector/nodejs:1.0.0

```

Notice how podman prefixes our tag with `localhost` for local images:
Docker uses default `docker.io` as prefix. </br>
Remember Kubernetes used to default to `docker.io` too, and no longer does. </br>
It's good practice to get used to referencing full image names (e.g. `docker.io/library/nginx` instead of just `nginx`) </br>

## Running Containers

Run a container the same way you would with Docker:

```shell
#interactive with -it or background with -d
podman run -d -p 80:5000 localhost/aimvector/python:1.0.0

# view containers
podman ps 

# check container logs
podman logs

# stop containers
podman stop

# view stopped containers
podman ps -a 

# remove containers
podman rm 

```

## Compose 

Podman also supports Compose files through `podman compose`:
Note in my demo its not found, as podman "wraps" docker compose under the hood if it exists. </br>

We can install it with `sudo apt-get install -y podman-compose`

```shell
podman compose build
```