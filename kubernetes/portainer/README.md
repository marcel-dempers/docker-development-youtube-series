# Introduction to Portainer

<a href="https://youtu.be/FC8pABzxZVU" title="k8s-portainer"><img src="https://i.ytimg.com/vi/FC8pABzxZVU/hqdefault.jpg" width="20%" alt="k8s-portainer" /></a> 

Start here 👉🏽[https://www.portainer.io/](https://www.portainer.io/) </br>
Documentation 👉🏽[https://docs.portainer.io/](https://docs.portainer.io/)

## Portainer installation

In this demo, I will be running Kubernetes 1.22 using `kind` </br>
Which is compatible with portainer 2.11.1 </br>

Let's go ahead with a local docker install: 

```
cd kubernetes\portainer
mkdir volume-ce 

docker run -d -p 9443:9443 -p 8000:8000 --name portainer-ce `
--restart=always `
-v /var/run/docker.sock:/var/run/docker.sock `
-v ${PWD}/volume-ce:/data `
portainer/portainer-ce:2.11.1
```

## SSL & DOMAIN

We can also upload SSL certificates for our portainer.</br> 
In this demo, portainer will issue self signed certificates. </br> 
We will need a domain for our portainer server so our clusters can contact it. </br>
Let's use [nip.io](https://nip.io/) to create a public endpoint for portainer.

## Create Kubernetes Cluster

Let's start by creating a local `kind` [cluster](https://kind.sigs.k8s.io/)

For local clusters, we can use the public endpoint Agent. </br>
We can get a public endpoint for the portainer agent by: </br>

* Ingress
* LoadBalancer
* NodePort

So we'll deploy portainer agent with `NodePort` for local </br>

For production environments, I would recommend not to expose the portainer agent. </br>
In this case, for Production, we'll use the portainer edge agent.  </br>


To get `NodePort` exposed in `kind`, we'll open a host port with a [kind.yaml](./kind.yaml) config

```
kind create cluster --name local --config kind.yaml
```

## Manage Kubernetes Environments 

The portainer UI gives us a one line command to deploy the portainer agent. </br>
Note that in the video, we pick the `node port` option.

## Local: Portainer Agent

I download the YAML from [here](https://downloads.portainer.io/portainer-agent-ce211-k8s-nodeport.yaml) to take a closer look at what it is deploying </br>

Deploy the portainer agent in my `kind` cluster:

```
kubectl apply -f portainer-agent-ce211-k8s-nodeport.yaml
```

See the agent: 

```
kubectl -n portainer get pods
```

See the service with the endpoint it exposes: 

```
kubectl -n portainer get svc
```

Now since we dont have a public load balancer and using nodeport, our service will be exposed on the node IP. </br>
Since the Kubernetes node is our local machine, we should be able to access the portainer agent on `<computer-IP>:30778` </br>

We can obtain our local IP with `ipconfig` </br>
The IP and NodePort will be used to connect our portainer server to the new agent. </br>

## Production: Portainer Edge Agent

For the Edge agent, we get the command in the portainer UI. </br>
Once deployed, we can see the egde agent in our AKS cluster:

```
kubectl -n portainer get pods
```

## Helm 

Let's showcase how to deploy helm charts. </br>
Most folks would have helm charts for their ingress controllers, monitoring, logging and other 
platform dependencies.</br>

Let's add Kubernetes NGINX Ingress repo:

```
https://kubernetes.github.io/ingress-nginx
```

## GitOps

So from the Application menu, we can add an application from a `git` repository. </br>
Let's add this repo:

```
https://github.com/marcel-dempers/docker-development-youtube-series
```

We also specify all our manifests path that portainer needs to deploy:

* kubernetes/portainer/example-application/deployment.yaml
* kubernetes/portainer/example-application/configmap.yaml
* kubernetes/portainer/example-application/service.yaml
* kubernetes/portainer/example-application/ingress.yaml

Portainer will now poll our repo and deploy any updates, GitOps style!

## Oauth Setup example | Business edition

Here are the values I used for the Oauth settings:

| Field | Value |
|-------|-------|
|  Client ID     |  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx     |
|  Client secret     |   xxxxxxxxxxxxxxxxxxxx    |
|  Authorization URL  |  https://login.microsoftonline.com/`<tenant-id>`/oauth2/authorize  |    
|  Access token URL  | https://login.microsoftonline.com/`<tenant-id>`/oauth2/token   |     
|  Resource URL   | https://login.microsoftonline.com/`<tenant-id>`/openid/userinfo   |     
|  Redirect URL   | https://localhost:9443/   |     
|  Logout URL   |  https://login.microsoftonline.com/`<tenant-id>`/oauth2/logout  |     
|  User identifier   |  unique_name  |     
|  Scopes   | openid profile   |  






