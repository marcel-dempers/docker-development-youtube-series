

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

# CI/CD examples

We can even run datree in GitHub Actions and various [CI/CD integrations](https://hub.datree.io/cicd-examples). </br>



