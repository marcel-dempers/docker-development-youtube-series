# Introduction to KUBECTL

<a href="https://youtu.be/1zcXudjSVUs" title="k8s-kubectl"><img src="https://i.ytimg.com/vi/1zcXudjSVUs/hqdefault.jpg" width="20%" alt="k8s-kubectl" /></a> 

To start off this tutorial, we will be using [kind](https://kind.sigs.k8s.io/) to create our test cluster. </br>
You can use `minikube` or any Kubernetes cluster. </br>

Kind is an amazing tool for running test clusters locally as it runs in a container which makes it lightweight and easy to run throw-away clusters for testing purposes. </br>

## Download KUBECTL

We can download `kubectl` from the [Official Docs](https://kubernetes.io/docs/tasks/tools/) </br>

## Create a kubernetes cluster

In this guide we will run two clusters side by side so we can demonstrate cluster access. </br>
Create two clusters:

```
kind create cluster --name dev --image kindest/node:v1.23.5
kind create cluster --name prod --image kindest/node:v1.23.5

```

See cluster up and running:

```
kubectl get nodes
NAME                  STATUS   ROLES                  AGE     VERSION
prod-control-plane   Ready    control-plane,master   2m12s   v1.23.5
```

## Understanding the KUBECONFIG

Default location of the `kubeconfig` file is in `<users-directory>/.kube/config`

```
kind: Config
apiVersion: v1
clusters:
 - list of clusters (addresses \ endpoints) 
users:
 - list of users (thing that identifies us when accessing a cluster [certificate]) 
contexts:
 - list of contexts ( which user and cluster to use when running commands)
```

Commands to interact with `kubeconfig` are `kubectl config`. </br>
Key commands are telling `kubectl` which context to use 

```
kubectl config current-context
kubectl config get-contexts
kubectl config use-context <name>
```

You can also tell your `kubectl` to use different config files. </br>
This is useful to keep your production config separate from your development ones </br>

Set the `$KUBECONFIG` environment variable to a path:
```
#linux
export KUBECONFIG=<path>

#windows 
$ENV:KUBECONFIG="C:\Users\aimve\.kube\config"
```

We can export seperate configs using `kind` </br>
This is possible with cloud based clusters as well:

```
kind --name dev export kubeconfig --kubeconfig C:\Users\aimve\.kube\dev-config 

kind --name prod export kubeconfig --kubeconfig C:\Users\aimve\.kube\prod-config 

#switch to prod
$ENV:KUBECONFIG="C:\Users\aimve\.kube\prod-config"
kubectl get nodes
```

## Working with Kubernetes resources

Now that we have cluster access, next we can read resources from the cluster
with the `kubectl get` command.

## Namespaces 

Most kubernetes resources are namespace scoped:

```
kubectl get namespaces
```

By default, `kubectl` commands will run against the `default` namespace

## List resources in a namespace

```
kubectl get <resource>

kubectl get pods
kubectl get deployments
kubectl get services
kubectl get configmaps
kubectl get secrets
kubectl get ingress
```

## Create resources in a namespace

We can create a namespace with the `kubectl create` command:

```
kubectl create ns example-apps
```

Let's create a couple of resources:

```

kubectl -n example-apps create deployment webserver --image=nginx --port=80
kubectl -n example-apps get deploy
kubectl -n example-apps get pods

kubectl -n example-apps create service clusterip webserver --tcp 80:80
kubectl -n example-apps get service
kubectl -n example-apps port-forward svc/webserver 80
# we can access http://localhost/

kubectl -n example-apps create configmap webserver-config --from-file config.json=./kubernetes/kubectl/config.json
kubectl -n example-apps get cm

kubectl -n example-apps create secret generic webserver-secret --from-file secret.json=./kubernetes/kubectl/secret.json
kubectl -n example-apps get secret

```

## Working with YAML

As you can see we can create resources with `kubectl` but this is only for basic testing purposes.
Kubernetes is a declarative platform, meaning we should provide it what to create instead
of running imperative line-by-line commands. </br>

We can also get the YAML of pre-existing objects in our cluster with the `-o yaml` flag on the `get` command </br>

Let's output all our YAML to a `yaml` folder:

```
kubectl -n example-apps get cm webserver-config -o yaml > .\kubernetes\kubectl\yaml\config.yaml
kubectl -n example-apps get secret webserver-secret -o yaml > .\kubernetes\kubectl\yaml\secret.yaml
kubectl -n example-apps get deploy webserver -o yaml > .\kubernetes\kubectl\yaml\deployment.yaml
kubectl -n example-apps get svc webserver -o yaml > .\kubernetes\kubectl\yaml\service.yaml   
```

## Create resources from YAML files

The most common and recommended way to create resources in Kubernetes is with the `kubectl apply` command. </br>
This command takes in declarative `YAML` files.

To show you how powerful it is, instead of creating things line-by-line, we can deploy all our infrastructure
with a single command. </br>

Let's deploy a Wordpress CMS site, with a back end MySQL database. </br>
This is a snippet taken from my `How to learn Kubernetes` video:

```
kubectl create ns wordpress-site
kubectl -n wordpress-site apply -f ./kubernetes/tutorials/basics/yaml/
```

We can checkout our site with the `port-forward` command:

```
kubectl -n wordpress-site get svc

NAME        TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
mysql       ClusterIP   10.96.146.75   <none>        3306/TCP   17s
wordpress   ClusterIP   10.96.157.6    <none>        80/TCP     17s

kubectl -n wordpress-site port-forward svc/wordpress 80
```

## Clean up

```
kind delete cluster --name dev
kind delete cluster --name prod

```