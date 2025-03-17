# Introduction to Kubernetes: Secrets

<a href="https://youtu.be/EkUN4V4Hmws" title="k8s-secrets"><img src="https://i.ytimg.com/vi/EkUN4V4Hmws/hqdefault.jpg" width="20%" alt="k8s-secrets" /></a> 

## Create a cluster with Kind

```
kind create cluster --name secrets --image kindest/node:v1.31.1
```

## Our Secret

We have a secret under `kubernetes/secrets/secret.json`

```
cat kubernetes/secrets/secret.json
```

## Using our secret in a container

As a file:
```
docker run -it -v $PWD/kubernetes/secrets/secret.json:/secrets/secret.json ubuntu:latest bash

cat /secrets/secret.json
```

As environment variables:

```
api_key="somesecretgoeshere"
docker run -it -e API_KEY=$api_key ubuntu:latest bash

echo $API_KEY
```

## Kubernetes Secret

Read more about [Kubernetes secrets](https://kubernetes.io/docs/concepts/configuration/secret/)


## Create our secret

There are two main ways we can create a Kubernetes secret. </br>
Either by creating the secret object with `kubectl create secret` or apply\create it declaratively using YAML with `kubectl apply -f`

`kubectl create secret`:

```
kubectl create secret generic mysecret --from-file kubernetes/secrets/secret.json
```

`kubectl apply -f` or `kubectl create -f` allows us to define things declaratively using YAML files:

```
kubectl apply -f kubernetes/secrets/secret.yaml
```

## Use our secret

In order to use our secret we add a `volume` to our pod spec and then mount that using a `volumeMount` </br>
We can also use a secret references as `env` variable </br>


```
kubectl apply -f kubernetes/secrets/pod.yaml
```



