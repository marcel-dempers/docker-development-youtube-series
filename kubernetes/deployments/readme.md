# Introduction to Kubernetes: Deployments

<a href="https://youtu.be/DMpEZEakYVc" title="k8s-deployments"><img src="https://i.ytimg.com/vi/DMpEZEakYVc/hqdefault.jpg" width="20%" alt="k8s-deployments" /></a> 

Build an example app:

```
#   Important!
#   make sure you are at root of the repository
#   in your terminal

#   you can choose which app you want to build!


#   aimvector/golang:1.0.0
docker-compose build golang


#   aimvector/csharp:1.0.0
docker-compose build csharp

#   aimvector/nodejs:1.0.0
docker-compose build nodejs

#   aimvector/python:1.0.0
docker-compose build python

```

Take a look at example [deployment yaml](./deployment.yaml)