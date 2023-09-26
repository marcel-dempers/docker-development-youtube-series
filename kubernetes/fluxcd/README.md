# Introduction to Flux CD v2

## Create a kubernetes cluster

In this guide we we''ll need a Kubernetes cluster for testing. Let's create one using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name fluxcd --image kindest/node:v1.26.3
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
curl -sLO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

```

### test cluster access:
```
/work # kubectl get nodes
NAME                    STATUS   ROLES    AGE   VERSION
fluxcd-control-plane   Ready    control-plane   54s   v1.26.3
```

## Get the Flux CLI

Let's download the `flux` command-line utility. </br>
We can get this utility from the GitHub [Releases page](https://github.com/fluxcd/flux2/releases) </br>

It's also worth noting that you want to ensure you get a compatible version of flux which supports your version of Kubernetes. Checkout the [prerequisites](https://fluxcd.io/flux/installation/#prerequisites) page. </br>

```
curl -o /tmp/flux.tar.gz -sLO https://github.com/fluxcd/flux2/releases/download/v2.1.1/flux_2.1.1_linux_amd64.tar.gz
tar -C /tmp/ -zxvf /tmp/flux.tar.gz
mv /tmp/flux /usr/local/bin/flux
chmod +x /usr/local/bin/flux
```

Now we can run `flux --help` to see its installed

## Check our cluster

```
flux check --pre
```

## Documentation

As with every guide, we start with the documentation </br>
The [Core Concepts](https://fluxcd.io/flux/concepts/) is a good place to start. </br>

We begin by following the steps under the [bootstrap](https://fluxcd.io/flux/installation/#bootstrap) section for GitHub </br>

We'll need to generate a [personal access token (PAT)](https://github.com/settings/tokens/new) that can create repositories by checking all permissions under `repo`.  </br>

Once we have a token, we can set it: 

```
export GITHUB_TOKEN=<your-token>
```

Then we can bootstrap it using the GitHub bootstrap method

```
flux bootstrap github \
  --token-auth \
  --owner=marcel-dempers \
  --repository=docker-development-youtube-series \
  --path=kubernetes/fluxcd/repositories/infra-repo/clusters/dev-cluster \
  --personal \
  --branch fluxcd-2022

flux check

# flux manages itself using GitOps objects:
kubectl -n flux-system get GitRepository
kubectl -n flux-system get Kustomization
```

Check the source code that `flux bootstrap` created 

```
git pull origin <branch-name>
```

# Understanding GitOps Repository structures

Generally, in GitOps you have a dedicated repo for infrastructure templates. </br>
Your infrastructure will "sync" from the this repo </br>

```
                                                    
 developer    +-----------+     +----------+           
              |           |     | CI       |           
  ----------> | REPO(code)|---> | PIPELINE |           
              +-----------+     +----------+           
                                         |  commit     
                                         v             
           +----------+ sync    +----------+           
           |  INFRA   |-------> |INFRA     |           
           |  (k8s)   |         |REPO(yaml)|           
           +----------+         +----------+           
                                                            
```

Flux repository structure [documentation](https://fluxcd.io/flux/guides/repository-structure/)

* Mono Repo (all k8s YAML in same "infra repo")
* Repo per team
* Repo per app 

Take note in this guide the folders under `kubernetes/fluxcd/repositories` represent different GIT repos

```
- repositories
  - infra-repo
  - example-app-1
  - example-app-2
```

## build our example apps

Let's say we have a microservice called `example-app-1` and it has its own GitHub repo somewhere. </br>
For demo, it's code is under `kubernetes/fluxcd/repositories/example-app-1/`

```
# go to our "git repo"
cd kubernetes/fluxcd/repositories/example-app-1
# check the files
ls

cd src
docker build . -t example-app-1:0.0.1

#load the image to our test cluster so we dont need to push to a registry
kind load docker-image example-app-1:0.0.1 --name fluxcd 
```

## setup our gitops pipeline

Now we will also have a "infra-repo" GitHub repo where infrastructure configuration files for GitOps live.

```
cd kubernetes/fluxcd

# tell flux where our Git repo is and where the YAML is
# this is once off
# flux will monitor the example-app-1 Git repo for when any infrastructure changes, it will sync
kubectl -n default apply -f repositories/infra-repo/apps/example-app-1/gitrepository.yaml
kubectl -n default apply -f repositories/infra-repo/apps/example-app-1/kustomization.yaml

# check our flux resources 
kubectl -n default describe gitrepository example-app-1
kubectl -n default describe kustomization example-app-1

# check deployed resources
kubectl get all

kubectl port-forward svc/example-app-1 80:80

```

Now we have setup CD, let's take a look at CI </br>

## changes to our example apps

Once we make changes to our `app.py` we can build a new image with a new tag </br>

```
docker build . -t example-app-1:0.0.2

#load the image to our test cluster so we dont need to push to a registry
kind load docker-image example-app-1:0.0.2 --name fluxcd 

# update our kubernetes deployment YAML image tag
# git commit with [skip ci] as the prefix of commit message & git push to branch!
```

If we wait a minute or so we can ` kubectl port-forward svc/example-app-1 80:80` again and see the changes

## automate deploy by updating manifest 

So all we did to update our app is to build a new image, push it to our registry and update the image tag in our kubernetes deployment YAML file and `flux` will sync it. </br>
This is generally the role of CI, where `flux` concern is mainly CD. </br>

Here is an example on [how to automate that](https://fluxcd.io/flux/use-cases/gh-actions-manifest-generation/) 

## automate deploy by image scanning

```                                                                                                                                               
                                              docker push                   
                                                                            
        developer    +-----------+     +----------+       +-------------+   
                     |           |     | CI       |       |IMAGE        |   
         ----------> | REPO(code)|---> | PIPELINE | ----->|REGISTRY     |   
                     +-----------+     +----------+       +-------------+   
                                                           ^                
                                                           |sync            
                                                           |                
                                   +----------+  commit   +----------+      
                                   |INFRA     | <-------- |  INFRA   |      
                                   |REPO(yaml)|           |  (k8s)   |      
                                   +----------+           +----------+      
                                                                         
```

An alternative method is to use your CI to build and push a newly tagged image to your registry (same as first option) and use [Flux image scanner](https://fluxcd.io/flux/guides/image-update/#configure-image-updates) to trigger the rollout instead of automating a commit to your config repo. </br>

We firstly need to enable image scanning as its not enabled by default. </br>
To do this we just need to re-bootstrap `flux` with an addition flag

```
flux bootstrap github \
  --token-auth \
  --owner=marcel-dempers \
  --repository=docker-development-youtube-series \
  --path=kubernetes/fluxcd/repositories/infra-repo/clusters/dev-cluster \
  --components-extra=image-reflector-controller,image-automation-controller \
  --personal \
  --branch fluxcd-2022
```
We need to create a image registry credential where we will push our image:

```
kubectl -n default create secret docker-registry dockerhub-credential --docker-username '' --docker-password '' --docker-email 'test@test.com'

```

# build and push example-app-2

```
cd kubernetes\fluxcd\repositories\example-app-2\
ls
cd src
ls
docker build . -t aimvector/example-app-2:0.0.1
docker push aimvector/example-app-2:0.0.1

```
We will need to tell Flux how to manage our image deployment </br>
Note that this time our Kubernetes YAML is in the `configs` repo. </br>
This is because our application repo triggers it's CI which will build and push a new image to our cluster </br>
Flux will then detect the new image tag and update our Kubernetes YAML in our configs repo. </br>
If Flux pushed the update to our application repo, it will cause a CI/CD loop.

## add image policy and repository

```

kubectl -n default apply -f kubernetes/fluxcd/repositories/infra-repo/apps/example-app-2/gitrepository.yaml
kubectl -n default apply -f kubernetes/fluxcd/repositories/infra-repo/apps/example-app-2/kustomization.yaml

# see our application 
kubectl get deploy
kubectl get pods

# tell flux about our image update policy
kubectl -n default apply -f kubernetes/fluxcd/repositories/infra-repo/apps/example-app-2/imagerepository.yaml
kubectl -n default apply -f kubernetes/fluxcd/repositories/infra-repo/apps/example-app-2/imagepolicy.yaml
kubectl -n default apply -f kubernetes/fluxcd/repositories/infra-repo/apps/example-app-2/imageupdateautomation.yaml

# we will also need to provide authentication for our git repo
flux create secret git example-app-2-github --url https://github.com/marcel-dempers/docker-development-youtube-series --username '' --password '' --namespace default
```

There are a number of ways to authenticate with [GitRepositories](https://fluxcd.io/flux/components/source/gitrepositories/#secret-reference)

```
kubectl describe imagepolicy example-app-2
kubectl describe  imagerepository example-app-2
kubectl describe imageupdateautomation example-app-2
```

## Build and push our example-app-2

```
#make application changes and rebuild + push

docker build . -t aimvector/example-app-2:0.0.2
docker push aimvector/example-app-2:0.0.2


#see changes new tags
kubectl describe imagerepository

#see image being updated
kubectl describe imagepolicy example-app-2

# see flux commiting back to the repo
kubectl describe imageupdateautomation example-app-2

```