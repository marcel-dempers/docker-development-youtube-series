

## Installation 

Best place to start is the [documentation](https://hub.datree.io/)

I like to start all my work inside a docker container. </br>
Let's run a small Alpine linux container

```
docker run -it -v ${PWD}:/work -w /work --net host alpine sh
```

Let's install `curl` and `unzip`

```
apk add curl unzip bash
```

We can install the latest version of Datree with the command advertised:

```
curl https://get.datree.io | /bin/bash
```


Or we can grab a specific version of `datree` on the GitHub releases page. </br>
For example: [1.5.9](https://github.com/datreeio/datree/releases/tag/1.5.9) binary

```
curl -L https://github.com/datreeio/datree/releases/download/1.5.9/datree-cli_1.5.9_Linux_x86_64.zip -o /tmp/datree.zip

unzip /tmp/datree.zip -d /tmp && \
chmod +x /tmp/datree && \
mv /tmp/datree /usr/local/bin/datree

```

Now we an run the `datree` command:

```
datree
Datree is a static code analysis tool for kubernetes files. Full code can be found at https://github.com/datreeio/datree

Usage:
  datree [command]

Available Commands:
  completion       Generate completion script for bash,zsh,fish,powershell
  config           Configuration management
  help             Help about any command
  kustomize        Render resources defined in a kustomization.yaml file and run a policy check against them
  publish          Publish policies configuration for given <fileName>.
  test             Execute static analysis for given <pattern>
  version          Print the version number

Flags:
  -h, --help   help for datree

Use "datree [command] --help" for more information about a command.

```

## Test Kubernetes Manifests

We have a number of Kubernetes manifests in this repo. </br>
Datree does a few things for us. </br>
* YAML validation ( Is this YAML well formatted ? )
* Schema validation. ( Is this a Kubernetes YAML file ? For the right version ? )
* Policy checks ( Checks YAML against best practise policies )

</br>

Let's test my example manifests under the `kubernetes` directory

### YAML validation

If we break the YAML file format, we can detect that with the YAML validation feature

```
datree test ./kubernetes/deployments/deployment.yaml
```

### Policy checks

When we fix our YAML file, notice if we run `datree test` again, we get some policy checks failing

```
datree test ./kubernetes/deployments/deployment.yaml

```

Let's test some other types of Kubernetes objects

```
datree test ./kubernetes/services/service.yaml
datree test ./kubernetes/configmaps/configmap.yaml
datree test ./kubernetes/statefulsets/statefulset.yaml
datree test ./kubernetes/ingress/ingress.yaml
```

### Schema validation

Datree kan also check if our YAML matches the target Kubernetes version schema.
For example, our Ingress YAML is a newer version of Kubernetes

```
datree test --schema-version 1.14.0 ./kubernetes/ingress/ingress-nginx-example.yaml
datree test --schema-version 1.19.0 ./kubernetes/ingress/ingress-nginx-example.yaml

```

We can also test a directory of YAML files. </br>
Let's test my latest Kubernetes tutorial that contains a Wordpress + MySQL + Ingress setup:

```
datree test kubernetes/tutorials/basics/yaml/*
```

# Policies

We can log into the Datree UI to get a view of the policy management screens

```
datree config set token <token>
```

Now that we have a token set, lets run a `datree test` command to see how `datree` checks our YAML against policies and provides us a UI for the output

```
datree test ./kubernetes/deployments/deployment.yaml
```

We can then review this test on the [Datree UI](https://hub.datree.io/)

# CI/CD examples

We can even run datree in GitHub Actions and various [CI/CD integrations](https://hub.datree.io/cicd-examples). </br>


# Admission Controller

So far, `datree` helps us detect misconfigurations on our local machine as well as at our CI level. </br>
But what about the things that don't flow via our CI ? </br>

When folks deploy stuff directly to our clusters via `kubectl` or `helm`. </br>
Datree now allows us to not only detect but prevent  misconfigurations being applied using a new admission controller feature. </br>

The admission controller is available [here](https://github.com/datreeio/admission-webhook-datree)

## Create a Kubernetes cluster

Let's start by creating a local `kind` [cluster](https://kind.sigs.k8s.io/)

Note that we create a Kubernetes 1.24 cluster. </br>
So we want to use `datree` to validate and ensure our manifests comply with that version of Kubernetes. <br/>

```
kind create cluster --name datree --image kindest/node:v1.24.2
```

We'll need a `datree` token so our admission controller can read our policies

```
DATREE_TOKEN=[your-token]

```

## Installation 

Let's grab the `datree` manifests
```
curl -L https://get.datree.io/admission-webhook -o datree.sh
chmod +x datree.sh
bash datree.sh
```

With the admission controller now deployed, `datree` will validate things coming into the cluster. <br/>
For example, if we bypass our CI/CD, `datree` will catch our deployment and run our policy checks


```
kubectl apply -f kubernetes/deployments/deployment.yaml
```

Output:

```
Error from server: error when creating "kubernetes/deployments/deployment.yaml": admission webhook "webhook-server.datree.svc" denied the request: 
---
webhook-example-deploy-Deployment.tmp.yaml

[V] YAML validation
[V] Kubernetes schema validation

[X] Policy check

❌  Ensure each container has a configured liveness probe  [1 occurrence]
    - metadata.name: example-deploy (kind: Deployment)
💡  Missing property object `livenessProbe` - add a properly configured livenessProbe to catch possible deadlocks

❌  Ensure each container has a configured readiness probe  [1 occurrence]
    - metadata.name: example-deploy (kind: Deployment)
💡  Missing property object `readinessProbe` - add a properly configured readinessProbe to notify kubelet your Pods are ready for traffic

❌  Prevent workload from using the default namespace  [1 occurrence]
    - metadata.name: example-deploy (kind: Deployment)
💡  Incorrect value for key `namespace` - use an explicit namespace instead of the default one (`default`)


(Summary)

- Passing YAML validation: 1/1

- Passing Kubernetes (v1.24.2) schema validation: 1/1

- Passing policy check: 0/1
+-----------------------------------+-----------------------+
| Enabled rules in policy "Default" | 21                    |
| Configs tested against policy     | 1                     |
| Total rules evaluated             | 21                    |
| Total rules skipped               | 0                     |
| Total rules failed                | 3                     |
| Total rules passed                | 18                    |
| See all rules in policy           | https://app.datree.io |
+-----------------------------------+-----------------------+

```
