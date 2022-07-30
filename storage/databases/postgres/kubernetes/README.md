
https://github.com/CrunchyData/postgres-operator

https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/

* Create a cluster 

```
kind create cluster --name postgres-example --image kindest/node:v1.23.5
```


docker run -it -v ${PWD}:/work -v ${HOME}/.kube/:/root/.kube/ -w /work --net host alpine sh 

apk add curl git

curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.23.5/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl


commit=c35b44b9bcabe6c1fea896bde043ff0e2d4bb43e
cd /tmp
git init
git remote add origin https://github.com/CrunchyData/postgres-operator-examples.git
git fetch --depth 1 origin $commit
git checkout $commit

