

## Installation 

Best place to start is the [documentation](https://hub.datree.io/)

I like to start all my work inside a docker container. </br>
Let's run a small Alpine linux container

```
docker run -it -v ${PWD}:/work -w /work --net host alpine sh
```

Let's install `curl` and `unzip`

```
apk add curl unzip
```

And finally grab the `datree` [1.5.9](https://github.com/datreeio/datree/releases/tag/1.5.9) binary

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
* YAML validation
* Schema validation. 
* Policy checks (there are 21 built-in policies at time of this demo)

</br>

Let's test my example manifests under the `kubernetes` directory

```
datree test ./kubernetes/deployments/deployment.yaml
datree test ./kubernetes/services/service.yaml
datree test ./kubernetes/configmaps/configmap.yaml
datree test ./kubernetes/statefulsets/statefulset.yaml
datree test ./kubernetes/ingress/ingress.yaml
```

Notice on my `ingress.yaml` the schema validation fails. </br>
This is a neat feature of `datree` since it checks for a few things: </br>

* Ensures the YAML is Kubernetes friendly. 
* Ensures its compatible with a Kubernetes version

It defaults to `1.19.0` as per time of this demo, and we can also change that on our account, or on the CLI

```
datree test --schema-version "1.19.0" ./kubernetes/ingress/ingress.yaml
datree test --schema-version "1.14.0" ./kubernetes/ingress/ingress.yaml
```

We can also test a directory of YAML files. </br>
Let's test my latest Kubernetes tutorial that contains a Wordpress + MySQL + Ingress setup:

```
datree test kubernetes/tutorials/basics/yaml/*
```

