
# Introduction to Datree

<a href="https://youtu.be/aqiOyXPPadk" title="Kubernetes"><img src="https://i.ytimg.com/vi/aqiOyXPPadk/hqdefault.jpg" width="20%" alt="Kubernetes Guide" /></a> 

## Installation 

Best place to start is the [documentation](https://hub.datree.io/)

I like to start all my work inside a docker container. </br>
Let's run a small Alpine linux container

```
docker run -it -v ${PWD}:/work -v ${HOME}/.kube/:/root/.kube/ -w /work --net host alpine sh 
```

### Install some dependancies 

Let's install `curl` and `unzip` because the installation script uses those. <br/>
We will also install `sudo` since we are running in a container as root and install scripts have `sudo` commands in them.

```
apk add curl unzip bash sudo
```

### Automatic Installation

We can install the latest version of Datree with the command advertised:

```
curl https://get.datree.io | /bin/bash
```

### Manual Installation

Or we can grab a specific version of `datree` on the GitHub releases page. </br>
For example: [1.5.20](https://github.com/datreeio/datree/releases/tag/1.5.20) binary

```
curl -L https://github.com/datreeio/datree/releases/download/1.5.20/datree-cli_1.5.20_Linux_x86_64.zip -o /tmp/datree.zip

unzip /tmp/datree.zip -d /tmp && \
chmod +x /tmp/datree && \
mv /tmp/datree /usr/local/bin/datree

```

Now we can run the `datree` command:

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

## Testing Kubernetes Manifests

We have a number of Kubernetes manifests in this repo. </br>
Datree does a few things for us: </br>
* YAML validation ( Is this YAML well formatted ? )
* Schema validation. ( Is this a Kubernetes YAML file ? For the right version ? )
* Policy checks ( Checks YAML to ensure good practises are followed )

</br>

Let's test my example manifests under our datree folder `kubernetes\datree\example`

### YAML validation

If we break the YAML file format, we can detect that with the YAML validation feature

```
datree test ./kubernetes/datree/example/deployment.yaml
```

### Policy checks

When we fix our YAML file, notice if we run `datree test` again, we get some policy checks failing

```
datree test ./kubernetes/datree/example/deployment.yaml

```

Let's test some other types of Kubernetes objects

```
datree test ./kubernetes/services/service.yaml
datree test ./kubernetes/configmaps/configmap.yaml
datree test ./kubernetes/statefulsets/statefulset.yaml
datree test ./kubernetes/ingress/ingress.yaml
```

### Schema validation

Datree can also check if our YAML matches the target Kubernetes version schema.
For example, our Ingress YAML is a newer version of Kubernetes

```
datree test --schema-version 1.14.0 ./kubernetes/ingress/ingress-nginx-example.yaml
datree test --schema-version 1.19.0 ./kubernetes/ingress/ingress-nginx-example.yaml

```

We can also test a directory of YAML files and include `*` wildcard in your scans. </br>
Let's test my latest Kubernetes tutorial that contains a Wordpress + MySQL + Ingress setup:

```
datree test kubernetes/tutorials/basics/yaml/*.y*ml
```

# Policies

Now if we take a look at the CLI output of `datree` we notice a link in the Summary output. </br>
The URL is in the form of `https://app.datree.io/login?t=<token>` </br>

```
(Summary)

- Passing YAML validation: 4/4

- Passing Kubernetes (1.20.0) schema validation: 4/4

- Passing policy check: 2/4

+-----------------------------------+------------------------------------------------------+
| Enabled rules in policy "Default" | 21                                                   |
| Configs tested against policy     | 5                                                    |
| Total rules evaluated             | 84                                                   |
| Total rules skipped               | 0                                                    |
| Total rules failed                | 14                                                   |
| Total rules passed                | 70                                                   |
| See all rules in policy           | https://app.datree.io/login?t=xxxxxxxxxxxxxxxxxxxxxx |
+-----------------------------------+------------------------------------------------------+
```

We can use this URL to access the Datree UI to get a view of the policy management screens </br>
Checkout the link to access the UI which helps us manage our policies. </br>

## Policy examples

One of the key features about policies is that we can apply rule sets for specific environments. </br>
Perhaps you have a development environment where policies are a little loose and a staging server that has tighter restrictions to match production, or even a regulated environment that has very tight controls. </br>

We can use the Datree UI to create policies with different sets of rules. </br>
We can then tell `datree` about the policy we want it to test against:

```
datree test kubernetes/datree/example/deployment.yaml -p production
```

For a new policy, we notice that 0 rules are enabled, so now we have the flexibility to set up the rules we want to protect this environment. </br>

## Helm

What if I don't use `kubectl` and use `helm` instead ? </br>
Let's install `helm` in our container </br>

```
apk add tar git
curl -L https://get.helm.sh/helm-v3.5.4-linux-amd64.tar.gz -o /tmp/helm.tar.gz && \
tar -xzf /tmp/helm.tar.gz -C /tmp && \
chmod +x /tmp/linux-amd64/helm && \
mv /tmp/linux-amd64/helm /usr/local/bin/helm

```

Let's install the `helm` plugin for `datree` <br/>

```
helm plugin install https://github.com/datreeio/helm-datree 

```

Now we can test a `helm` chart we have in our repo from my `helm` tutorial </br>

```

cd kubernetes/helm

helm datree test example-app \
-- --values ./example-app/example-app-01.values.yaml
```

## Kustomize

What if I don't use `helm` and use `kustomize` instead ? <br/>
Datree has out the box built-in `kustomize` support <br/>
Let's test our `kustomize` template from a video I did on `kustomize`

```
datree kustomize test .\kubernetes\kustomize\application\
```

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

Note that we create a Kubernetes 1.23 cluster. </br>
So we want to use `datree` to validate and ensure our manifests comply with that version of Kubernetes. <br/>

```
kind create cluster --name datree --image kindest/node:v1.23.6
```

Let's also grab `kubectl`:

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.23.6/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

We'll need a `datree` token so our admission controller can read our policies

```
export DATREE_TOKEN=[your-token]

```

## Installation 

I will need some dependencies since I am running in a lightweight `alpine` container. </br>
OpenSSL is needed by the webhook install to generate certificates. </br>

```
apk add openssl
```

Let's grab the `datree` manifests
```
curl -L https://get.datree.io/admission-webhook -o datree.sh
chmod +x datree.sh
bash datree.sh
```

With the admission controller now deployed, `datree` will validate things coming into the cluster. <br/>
For example, if we bypass our CI/CD, `datree` will catch our deployment and run our policy checks

I have a separate example deployment in our datree folder that we can play with:

```
kubectl apply -f kubernetes/datree/example/deployment.yaml
```

Output:

```
kubectl apply -f kubernetes/deployments/deployment.yaml
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

- Passing Kubernetes (v1.23.6) schema validation: 1/1

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

Now to get this deployment fixed up, let's go ahead and comply to some of the policies </br>
Under the `deployment.yaml` I have included a `livenessProbe` as well as a `readinessProbe` </br>
Let's add those in. </br>
And finally we need to also add CPU and Memory requests and limit values. </br>

The last one is simple. We should avoid using the default namespace. So I will create an `example` namespace where I will keep all example apps.

```
kubectl create ns examples
```

And finally we can deploy our resource, and specify a namespace:

```
kubectl apply -n examples -f kubernetes/datree/example/deployment.yaml
deployment.apps/example-deploy created

```

## Kubectl

But what about resources already in your cluster ? </br>
Datree covers this with their `kubectl` plugin.

We can grab the install script right off the [GitHub Release](https://github.com/datreeio/kubectl-datree/releases) page. </br>
For this demo I'll grab the `v0.11` version </br>

Installation: 

```
curl -L https://github.com/datreeio/kubectl-datree/releases/download/v0.1.1/manual_install.sh -o /tmp/kubectl-plugin.sh
chmod +x /tmp/kubectl-plugin.sh
bash /tmp/kubectl-plugin.sh

```

Now we have datree inside `kubectl` and can perform checks in our cluster. </br>
We can check our entire namespace now, which should be pretty clean:

```
kubectl datree test -- --namespace examples
Fetching resources, this may take some time depending on the amount of resources in your cluster...

(Summary)

- Passing YAML validation: 1/1

- Passing Kubernetes (1.24.2) schema validation: 1/1

- Passing policy check: 1/1

+-----------------------------------+------------------------------------------------------+
| Enabled rules in policy "Default" | 21                                                   |
| Configs tested against policy     | 1                                                    |
| Total rules evaluated             | 21                                                   |
| Total rules skipped               | 0                                                    |
| Total rules failed                | 0                                                    |
| Total rules passed                | 21                                                   |
| See all rules in policy           | https://app.datree.io/login?t=xxxxxxxxxxxxxxxxxxxxxx |
+-----------------------------------+------------------------------------------------------+

The following cluster resources in namespace 'examples' were checked:

deployment.apps/example-deploy

```