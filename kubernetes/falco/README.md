https://falco.org/

* Learning environment

https://falco.org/docs/getting-started/third-party/learning/


* Create a cluster 

```
cd kubernetes/falco
kind create cluster --name falco --image kindest/node:v1.23.5 --config kind.yaml
```

* Install helm

```
curl -LO https://get.helm.sh/helm-v3.7.2-linux-amd64.tar.gz
tar -C /tmp/ -zxvf helm-v3.7.2-linux-amd64.tar.gz
rm helm-v3.7.2-linux-amd64.tar.gz
mv /tmp/linux-amd64/helm /usr/local/bin/helm
chmod +x /usr/local/bin/helm
```

Add falcosecurity repository
Before installing the chart, add the falcosecurity charts repository:

```
helm repo add falcosecurity https://falcosecurity.github.io/charts
helm repo update
```

Let's find what versions of vault are available:

```
helm search repo falcosecurity --versions
```

I will be using the below version in this demo

```
NAME                        	CHART VERSION	APP VERSION	DESCRIPTION                                       
falcosecurity/falco         	1.18.5       	0.31.1     	Falco 
```

* Get YAML template 

Let's grab the manifests:

```
mkdir manifests
helm template falcosecurity  falcosecurity/falco \
  --namespace falco \
  --version 1.18.5 \
  > ./manifests/falco.yaml
```


* Daemonset 

https://github.com/falcosecurity/evolution/tree/master/deploy/kubernetes/falco/templates


```
kubectl create ns falco
kubectl -n falco apply -f ./kubernetes/falco/templates/
```