---
name: k8s-local-provisioner
description: Use this skill to provision local kubernetes clusters 
---

# K8s Local Provisioner

This skills allows an agent to provision a local kubernetes cluster. 

## Requirements

In order to be able to provision a local kubernetes cluster, we need a few tools:

* curl
* kind
* kubectl 
* docker ce cli

If any of these tools are not installed, use the provided scripts from the skill to install these 

You must use the scripts provided to create the local kubernetes cluster.