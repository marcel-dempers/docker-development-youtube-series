# Introduction to Crossplane

[Crossplane](https://www.crossplane.io/) </br>
[Crossplane Documentation](https://docs.crossplane.io/v1.19/) </br>

## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name crossplane --image kindest/node:v1.33.0
```

## Installing Crossplane

In this guide we will reference the official document steps in the links above. </br>
I've recorded the commands we follow in the video too 


```
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm search repo crossplane-stable --versions
```

We'll install version `1.19.1` at the time of this guide 

```
VERSION=1.19.1

helm install crossplane \
crossplane-stable/crossplane \
--namespace crossplane-system \
--version $VERSION \
--create-namespace
```

View our install: 

```
kubectl get pods -n crossplane-system
kubectl get deployments -n crossplane-system
```

Once the pods are all running, we can see the `api-versions` 

```
kubectl api-versions  | grep crossplane
```

We can also see the new k8s objects that are installed with 

```
kubectl api-resources | grep crossplane
```

## Providers 

[Providers](https://docs.crossplane.io/latest/concepts/providers/) allow us to setup external providers that helps provision infrastructure for external services. </br>

For example, our crossplane cluster may have providers for deploying Azure, AWS, GCP or any other external infrastructure </br>

Furthermore, there is marketplace that hosts many providers, configurations and extensions for Crossplane called [Upbound](https://marketplace.upbound.io/providers)


Install a Provider for a cloud provider Azure:

```
kubectl apply -f kubernetes/crossplane/provider-azure.yaml
```

Check our provider:

```
kubectl get provider
kubectl describe provider provider-family-azure
```

## Provider Configuration

Once we have a provider setup, we can configure it using a `ProviderConfig` in Kubernetes </br>
An impotrant configuration is to tell the Crossplane Provider how to authenticate with its external service. </br>

For example, when using an Azure Provider, you need an Azure Service Principal, and for AWS you may need a service account with AWS account id and key. </br>
Each provider will have their own supported authentication methods. </br>

### Create Provider credentials 

```
SUBSCRIPTION_ID=<subscription-id>
RESOURCE_GROUP=marcel-test

az account set -s $SUBSCRIPTION_ID
az group create -n $RESOURCE_GROUP -l australiaeast
az ad sp create-for-rbac --sdk-auth  \
-n marcel-test \
--role Contributor \
--scopes "/subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP" > azure-credentials.json
```

### Create Provider Kubernetes Secret 

```
kubectl create secret \
generic azure-secret \
-n crossplane-system \
--from-file=creds=./azure-credentials.json
```

### Deploy the Provider Configuration

```
kubectl apply -f kubernetes/crossplane/providerconfig-azure.yaml
```

### Create Provider Resources 

```
kubectl apply -f kubernetes/crossplane/resources/azure/resource-vnet.yaml

error: resource mapping not found for name: "marcel-test-vnet" namespace: "" from "kubernetes/crossplane/resources/azure/resource-vnet.yaml": no matches for kind "VirtualNetwork" in version "network.azure.upbound.io/v1beta1"
ensure CRDs are installed first

```
We see there is no CRD for Azure VNETs, that is because every type of resource in Azure is modularized into a separate provider, so we will need the networking provider first </br>

Install the Azure Network Provider:

```
kubectl apply -f kubernetes/crossplane/provider-azure-network.yaml
```

Retry the resource creation:

```
kubectl apply -f kubernetes/crossplane/resources/azure/resource-vnet.yaml
kubectl get virtualnetwork
```

### Deploy a Virtual Network Subnet 

```
kubectl apply -f kubernetes/crossplane/resources/azure/resource-subnet.yaml
kubectl get subnet
```

### Deploy a Virtual Network Card 

```
kubectl apply -f kubernetes/crossplane/resources/azure/resource-networkcard.yaml
kubectl get networkinterface
```

### Deploy a Virtual Machine 

Firstly need to add the compoute provider for Azure

```
kubectl apply -f kubernetes/crossplane/provider-azure-compute.yaml
```

Deploy a Virtual Machine:

```
kubectl apply -f kubernetes/crossplane/resources/azure/resource-virtualmachine.yaml
```

## Cleanup Resources

```
kubectl delete linuxvirtualmachine marcel-test
kubectl delete networkinterface marcel-test
kubectl delete subnet marcel-test
kubectl delete virtualnetwork marcel-test
```