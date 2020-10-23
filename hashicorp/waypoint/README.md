# Introduction to Waypoint

I will do the following: <br/>

* Run a local Kubernetes cluster with [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
* Run a container to extract the windows `waypoint` binary + move it to $PATH


## We need a Kubernetes cluster

Lets create a Kubernetes cluster to play with using [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```
kind create cluster --name waypoint --image kindest/node:v1.19.1
```

## Get the Waypoint binary for Windows

<br/>
Feel free to follow the official install guide using the [docs](https://learn.hashicorp.com/tutorials/waypoint/get-started-install) <br/>

<br/>

Use a container to download and extract the binary:

```
docker run -it --rm -v ${PWD}:/work -w /work alpine sh -c 'apk add --no-cache curl unzip && curl -LO https://releases.hashicorp.com/waypoint/0.1.3/waypoint_0.1.3_windows_amd64.zip && unzip waypoint_0.1.3_windows_amd64.zip && rm waypoint_0.1.3_windows_amd64.zip'

```

I've setup my Windows $PATH environment to point to a folder `C:\kubectl\`

```
mv waypoint.exe C:\kubectl\waypoint.exe

#open new powershell terminal

waypoint --help
```

<br/>

## Install Waypoint server 

```
kubectl port-forward svc/waypoint 9701:9701
waypoint install --platform=kubernetes -accept-tos
```


## Waypoint Kubernetes + Python example

```
cd hashicorp/waypoint/python/

waypoint init
waypoint up
```