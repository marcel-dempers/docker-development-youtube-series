# Introduction to Flux CD v2

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name fluxcd --image kindest/node:v1.23.5
```

See cluster up and running:

```
kubectl get nodes
NAME                  STATUS   ROLES                  AGE     VERSION
fluxcd-control-plane   Ready    control-plane,master   2m12s   v1.23.5
```

## Run a container to work in

### run Alpine Linux: 
```
docker run -it --rm -v ${HOME}:/root/ -v ${PWD}:/work -w /work --net host alpine sh
```

### install some tools

```
# install curl 
apk add --no-cache curl

# install kubectl 
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

# install helm 

curl -o /tmp/helm.tar.gz -LO https://get.helm.sh/helm-v3.10.1-linux-amd64.tar.gz
tar -C /tmp/ -zxvf /tmp/helm.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm

```

### test cluster access:
```
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
fluxcd-control-plane   Ready    control-plane,master   3m26s   v1.23.5
```

## Flux CD 

### get flux command-line tool

Let's download the `flux` command-line utility. </br>
We can get this utility from the GitHub [Releases page](https://github.com/fluxcd/flux2/releases) </br>

It's also worth noting that you want to ensure you get a compatible version of flux which supports your version of Kubernetes. Checkout the [prerequisites](https://fluxcd.io/flux/installation/#prerequisites) page. </br>

```
curl -o /tmp/flux.tar.gz -LO https://github.com/fluxcd/flux2/releases/download/v0.36.0/flux_0.36.0_linux_amd64.tar.gz
tar -C /tmp/ -zxvf /tmp/flux.tar.gz
mv /tmp/flux /usr/local/bin/flux
chmod +x /usr/local/bin/flux
```

Now we can run `flux --help` to see its installed

## Documentation

As with every guide, we start with the documentation </br>
The [Core Concepts](https://fluxcd.io/flux/concepts/) is a good place to start. </br>

We begin by following the steps under the [bootstrap](https://fluxcd.io/flux/installation/#bootstrap) section for GitHub </br>

We'll need to generate a personal access token (PAT) that can create repositories by checking all permissions under `repo`.  </br>

Once we have a token, we can set it: 

```
export GITHUB_TOKEN=<your-token>
```

Then we can bootstrap it using the GitHub bootstrap method

```
flux bootstrap github \
  --owner=my-github-username \
  --repository=my-repository \
  --path=clusters/my-cluster \
  --personal
```
