# Introduction to Octopus Deploy

<a href="https://youtu.be/xxxxx" title="octo-intro"><img src="https://i.ytimg.com/vi/xxxxx/hqdefault.jpg" width="20%" alt="octo-intro" /></a> 

Start here 👉🏽[https://octopus.com/](https://octopus.com/) </br>
Documentation 👉🏽[https://octopus.com/docs](https://octopus.com/docs)

In this video, I sign up with the Free version of Octopus deploy to get my license key. </br>

## Create Kubernetes clusters

In my video I use `kind` to act as my DEV and PROD environment where I run microservices </br>
Let's create that:
```shell
kind create cluster --name dev --image kindest/node:v1.34.0
kind create cluster --name prod --image kindest/node:v1.34.0
```

## Octopus installation

In this video, I will explore the architecture of Octopus and the required components, so I'll be using the self-hosted option for learning. </br>

I highly recommend you checkout the Cloud version. </br>

My settings:
```
cd automation/cicd/octopus-deploy

export MASTER_KEY=$(openssl rand -base64 32)
export OCTOPUS_SERVER_BASE64_LICENSE=xxxxx
export ADMIN_USERNAME="octo-admin"
export ADMIN_PASSWORD="octo-admin123"
export ADMIN_EMAIL="octo-admin@admin.com"
```

Spin up a database for Octopus to use:

```
docker run -d --name octopus-database \
--net kind \
-u root \
-e ACCEPT_EULA="Y" \
-e MSSQL_SA_PASSWORD="OurStrongPassword123" \
-v ${PWD}/.data/mssql/:/var/opt/mssql/data \
mcr.microsoft.com/mssql/server:2022-latest
```

Create our first Octopus server:
```
docker run -it \
--name octopus \
--net kind \
-p 8080:8080 \
-p 8443:8443 \
-p 10943:10943 \
-e ACCEPT_EULA="Y" \
-e ADMIN_USERNAME=$ADMIN_USERNAME \
-e ADMIN_PASSWORD=$ADMIN_PASSWORD \
-e ADMIN_EMAIL=$ADMIN_EMAIL \
-e OCTOPUS_SERVER_BASE64_LICENSE=$OCTOPUS_SERVER_BASE64_LICENSE \
-e MASTER_KEY=$MASTER_KEY \
-e DB_CONNECTION_STRING="Server=octopus-database,1433;Database=OctopusDeploy;User=sa;Password=OurStrongPassword123" \
-v ${PWD}/.data/octopus/repository/:/repository \
-v ${PWD}/.data/octopus/artifacts/:/artifacts \
-v ${PWD}/.data/octopus/taskLogs/:/taskLogs \
-v ${PWD}/.data/octopus/cache/:/cache \
-v ${PWD}/.data/octopus/import/:/import \
-v ${PWD}/.data/octopus/eventExports/:/eventExports \
octopusdeploy/octopusdeploy:2025.3

```

## Access Octopus Dashboard

Let's get the IP of my machine so we can access Octopus:

```
ifconfig eth0 | grep 'inet' | awk '{print $2}' | cut -d: -f2
```

## Setup Environments 

The Octopus UI helps us setup and configure environments. </br>
In this video, we add two Kubernetes environments, `k8s-dev` and `k8s-prod` </br>

## Setup Deployments

So from the Application menu, we can add an application from a `git` repository. </br>
Let's add this repo:

```
https://github.com/marcel-dempers/docker-development-youtube-series.git
```

We also specify all our manifests path that portainer needs to deploy:

```
automation/cicd/octopus-deploy/example-application/deployment.yaml
automation/cicd/octopus-deploy/example-application/configmap.yaml
automation/cicd/octopus-deploy/example-application/service.yaml
automation/cicd/octopus-deploy/example-application/ingress.yaml
```

### Variable injection

spec:rules:0:host

```
spec:
  ingressClassName: nginx
  rules:
  - host: marcel.local
```






