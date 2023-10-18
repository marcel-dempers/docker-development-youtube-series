<a href="https://youtu.be/RcHGqCBofvw" title="githubactions"><img src="https://i.ytimg.com/vi/RcHGqCBofvw/hqdefault.jpg" width="20%" alt="introduction to github actions runners" /></a> 

# Introduction to GitHub Actions: Self hosted runners

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name githubactions --image kindest/node:v1.28.0@sha256:b7a4cad12c197af3ba43202d3efe03246b3f0793f162afb40a33c923952d5b31
```

Let's test our cluster:
```
kubectl get nodes
NAME                          STATUS   ROLES           AGE     VERSION
githubactions-control-plane   Ready    control-plane   2m53s   v1.28.0
```

## Running the Runner in Docker 

We can simply install this directly on to virtual machines , but for this demo, I'd like to run it in Kubernetes inside a container. </br>

### Security notes

* Running in Docker needs high priviledges.
* Would not recommend to use these on public repositories.
* Would recommend to always run your CI systems in seperate Kubernetes clusters.

### Creating a Dockerfile

* Installing Docker CLI 
For this to work we need a `dockerfile` and follow instructions to [Install Docker](https://docs.docker.com/engine/install/debian/).
I would grab the content and create statements for my `dockerfile` </br>

Now notice that we only install the `docker` CLI. </br> 
This is because we want our running to be able to run docker commands , but the actual docker server runs elsewhere </br>
This gives you flexibility to tighten security by running docker on the host itself and potentially run the container runtime in a non-root environment </br>

* Installing Github Actions Runner 

Next up we will need to install the [GitHub actions runner](https://github.com/actions/runner) in our `dockerfile`
Now to give you a "behind the scenes" of how I usually build my `dockerfile`s, I run a container to test my installs: 

```
docker build . -t github-runner:latest 
docker run -it github-runner bash
```

Next steps:

* Now we can see `docker` is installed 
* To see how a runner is installed, lets go to our repo | runner and click "New self-hosted runner"
* Try these steps in the container
* We will needfew dependencies
* We download the runner
* TODO


Finally lets test the runner in `docker` 

```
docker run -it -e GITHUB_PERSONAL_TOKEN="" -e GITHUB_OWNER=marcel-dempers -e GITHUB_REPOSITORY=docker-development-youtube-series github-runner
```

## Deploy to Kubernetes 

Load our github runner image so we dont need to push it to a registry:

```
kind load docker-image github-runner:latest --name githubactions
```

Create a kubernetes secret with our github details 

```
kubectl create ns github
kubectl -n github create secret generic github-secret `
  --from-literal GITHUB_OWNER=marcel-dempers `
  --from-literal GITHUB_REPOSITORY=docker-development-youtube-series `
  --from-literal GITHUB_PERSONAL_TOKEN=""

kubectl -n github apply -f kubernetes.yaml
```
