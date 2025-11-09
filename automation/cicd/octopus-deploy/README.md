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

## Octopus Components & Installation

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

[notes: ports to expose](https://octopus.com/docs/installation/octopus-server-linux-container#exposed-container-ports)
```
docker run -d \
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

We'll setup targets that will allow us to deploy to each environment. 

## Setup Targets

The Environment page allows us to configure deployment targets. </br>
The Octopus UI guides us to install the Kubernetes Agent using Helm. </br>
In this video, we will add targets for each Kubernetes environment so we can target those environments for deployment. </br>

## Script Modules

Octopus allows us to create script modules that we can use during deployments. 
Instead of placing raw scripts into deployment steps, we can create re-usable modules

Let's create a simple module called `create_namespace` that creates a kubernetes namespace: 

```bash
create_namespace() {
NAMESPACE_NAME=$1

# Check if the namespace exists.
# The 'get namespace' command will fail if the namespace is not found.
# The '>/dev/null 2>&1' part suppresses all output (success or failure) 
# so the script remains silent unless a new namespace is created.

kubectl get namespace $NAMESPACE_NAME >/dev/null 2>&1

# $? is the exit status of the last command (0 for success, non-zero for failure)
if [ $? -ne 0 ]; then
    echo "Namespace '$NAMESPACE_NAME' not found. Creating it now..."
    kubectl create namespace $NAMESPACE_NAME
    echo "Namespace '$NAMESPACE_NAME' created successfully."
else
    echo "Namespace '$NAMESPACE_NAME' already exists. Skipping creation."
fi
}

```

Now we'll be able to source that later in a deployment step. </br>

## Setup Project 

We'll setup a project to deploy our service. A project holds the process & steps as well as variables and features that allow us to deploy our software. </br>

## Setup Deployment Steps

We can now use a Script step template to run our script module 

```bash
source create_namespace.sh
create_namespace $(get_octopusvariable "NAMESPACE")
```

## Setup Variables

Variables allow us to configure almost anything about our steps, and scripts. They can be config values or even secrets. </br>
In this video we'll add a `NAMESPACE` variable for our `product` namespace. </br>

## Setup Kubernetes Step

Let's add a new step to "Deploy Kubernetes YAML"
We can refernce YAML in many ways, in this video we use Github:
```
https://github.com/marcel-dempers/docker-development-youtube-series.git
```

We also specify all our manifests path that portainer needs to deploy:

```
automation/cicd/octopus-deploy/example-application/deployment.yaml
automation/cicd/octopus-deploy/example-application/service.yaml
automation/cicd/octopus-deploy/example-application/ingress.yaml
```

### Variable injection

Octopus allows us to manipulate files, and flows and achieve different outcomes in many ways. One powerful feature is injecting values into files. </br>

We can use this for configuration value injection, secret injection etc. </br>

Let's update contents in our `configmap.yaml` based on our environment.
We can inject values into our JSON block based on the environment we are deploying too.

spec:rules:0:host

```
spec:
  ingressClassName: nginx
  rules:
  - host: marcel.local
```

## Create a step template

In this video, we also create a new step template which allows us to write modular templates that can be re-used throughout Octopus.
Since may microservices may use `ConfigMap` objects, we can create a step template for it </br>

Our step template name: `Deploy ConfigMap to Kubernetes` </br>
Parameters:
* `config_name` of the ConfigMap to create
* `config_namespace` name of the Namespace to deploy the ConfigMap to
* `config_body` the contents of the ConfigMap as Multi-line text

Inline Source Code:

```bash
NAMESPACE=$(get_octopusvariable "config_namespace")
CONFIGMAP_NAME=$(get_octopusvariable "config_name")
BODY_JSON=$(get_octopusvariable "config_body")

CONFIG_KEY=config.json
# Check if arguments are provided
if [ -z "$NAMESPACE" ] || [ -z "$CONFIGMAP_NAME" ] || [ -z "$BODY_JSON" ]; then
    echo "Error: Missing one or more required arguments."
    usage
fi

kubectl create configmap "${CONFIGMAP_NAME}" --namespace "${NAMESPACE}" \
  --from-literal="${CONFIG_KEY}=${BODY_JSON}" \
  --dry-run=client -o yaml | kubectl apply -f -
```

We can then proceed to add this step to our deployment process. </br>

We can do the same and create a Secret deployment template:

```bash
NAMESPACE=$(get_octopusvariable "secret_namespace")
SECRET_NAME=$(get_octopusvariable "secret_name")
BODY_JSON=$(get_octopusvariable "secret_body")

SECRET_KEY="secret.json"

# Check if arguments are provided
if [ -z "$NAMESPACE" ] || [ -z "$SECRET_NAME" ] || [ -z "$BODY_JSON" ]; then
    echo "Error: Missing one or more required arguments."
    usage
fi

ENCODED_BODY=$(echo -n "$BODY_JSON" | base64)

kubectl create secret generic "${SECRET_NAME}" --namespace "${NAMESPACE}" \
  --from-literal="${SECRET_KEY}=${ENCODED_BODY}" \
  --dry-run=client -o yaml | kubectl apply -f -

```