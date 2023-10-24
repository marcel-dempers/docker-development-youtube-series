# Running PostgreSQL in Kubernetes (Basic)

In chapters [one](../1-introduction/README.md), [two](../2-configuration/README.md) and [three](../3-replication/README.md) we've managed to stand up a Primary and Stand-By PostgreSQL instances using containers. </br>

We've learnt the fundamentals of how to persist and store data, how to configure instances and how to setup streaming replication from a Primary container to a Stand-by container. </br>

## The challenges

We have encountered a few challenges along the way, but running PostgreSQL in a container is pretty similar to just running it on a server outside of a container. </br>

</hr>

Kubernetes will add a bunch more complexity which we'll cover in this chapter. </br>
A few points to note:

* If you are not familiar with running PostreSQL in a container, this chapter is not for you. Please go back to [Chapter 1](../1-introduction/README.md)
* If you are not familiar with configuration of PostreSQL, do not attempt to run it in Kubernetes. Please go back to [Chapter 2](../2-configuration/README.md)
* If you are not familiar with Streaming Replication, Do not attempt to run PostreSQL in Kubernetes. Please go back to [Chapter 3](../3-replication/README.md)
* If you are not familiar with [StatefulSets](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/), Do not attempt to run PostreSQL in Kubernetes
* We will not be using Popular PostgreSQL controllers\operators or Helm charts in this guide. Operators and controllers simply automate things, and those open source tooling assumes you understand all the above mentioned tech.

One caveat to think of before running PostgreSQL in Kubernetes, or any database for that matter, is how would you handle cluster upgrades? </br>
Most cloud providers uprade by rolling new nodes and deleting old nodes, meaning your primary server may be deleted and start on a new node without any data. </br> If you don't have a strategy here, you will lose your data. </br>

If something goes wrong and you're using operators or controllers and don't have a background in how PostgreSQL works, you will lose data. </br>

And finally - The work in this guide has not been tested for Production workloads and written purely for educational purposes. </br>

## Create a Kubernetes cluster

In this chapter, we will start by creating a test Kubernetes cluster using [kind](https://kind.sigs.k8s.io/) </br>

```
kind create cluster --name postgresql --image kindest/node:v1.28.0

kubectl get nodes
NAME                       STATUS   ROLES                  AGE   VERSION
postgresql-control-plane   Ready    control-plane,master   31s   v1.28.0
```

## Setting up our PostgreSQL environment

Deploy a namespace to hold our resources: 

```
kubectl create ns postgresql
```

In [Chapter 3](../3-replication/README.md), we defined a few environment variables in our `docker run` command. </br>
Some of those values are sensitive, so in Kubernetes we'll place sensitive values in a Kubernetes [secret](https://kubernetes.io/docs/concepts/configuration/secret/). </br>


Create our secret for our first PostgreSQL instance:

```
kubectl -n postgresql create secret generic postgresql `
  --from-literal POSTGRES_USER="postgresadmin" `
  --from-literal POSTGRES_PASSWORD='admin123' `
  --from-literal POSTGRES_DB="postgresdb" `
  --from-literal REPLICATION_USER="replicationuser" `
  --from-literal REPLICATION_PASSWORD='replicationPassword'
```

## Deploy our first PostgreSQL instance

### Statefulsets

As we know we're going to need state and persist data, we'll go create a [StatefulSet](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)

I've taken a copy of the Statefulset from the Kubernetes site and created a [statefulset.yaml](./yaml/statefulset.yaml) for reference. </br>

In the video, we'll replace some of these values so our PostgreSQL will fit. </br>

* We replace the name of `nginx` with `postgres`
* Tweak the k8s service
* Tweak the [statefulset.yaml](./yaml/statefulset.yaml)
* Add environment variables and secret mappings.
* Add our configurations in a [Configmap](https://kubernetes.io/docs/concepts/configuration/configmap/)

We will take a look at Replication in the following chapter, so our replication user will not exist in the database just yet. </br>

Deploy our PostgreSQL instance:

```
kubectl -n postgresql apply -f storage/databases/postgresql/4-k8s-basic/yaml/statefulset.yaml
```

### Check our installation

```
kubectl -n postgresql get pods

# check the database logs
kubectl -n postgresql logs postgres-0

```
You will notice archive errors in the logs, because the archive directory does not exist in our volume. </br>
We will address this soon. </br>

Let's check our instance further:

```
kubectl -n postgresql exec -it postgres-0 -- bash

# login to postgres
psql --username=postgresadmin postgresdb

# see our replication user created
\du

#create a table
CREATE TABLE customers (firstname text, customer_id serial, date_created timestamp);

#show the table
\dt

# quit out of postgresql
\q

# check the data directory
ls -l /data/pgdata

# check the archive (does not exist!)
ls -l /data/archive
```

### Init containers

[Init containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) play a big role in fulfilling specific needs in IT workloads </br>

Init containers run before other containers in our pods. It can greatly assist when we need to do manual tasks. Like creating users, setting up tables, etc. </br>
Init containers can help us initialise things, like creating this `/data/archive` directory </br>

Now it may seem overkill for simply creating a directory, however this init container will play a big role in our next chapter on replication. </br>

We can use init containers to setup our postgres as a primary, or a standby server. Stay tuned! </br>

Let's create our init container:

```
initContainers:
- name: init
  image: postgres:15.0
  command: [ "bash", "-c" ]
  args:
  - |
    #create archive directory
    mkdir -p /data/archive && chown -R 999:999 /data/archive
```

This init container also needs to share the volume of the database container:

```
volumeMounts:
- name: data
  mountPath: /data
  readOnly: false
```

And redeploy!

```
kubectl -n postgresql apply -f storage/databases/postgresql/4-k8s-basic/yaml/statefulset.yaml

# check our install

kubectl -n postgresql get pods
kubectl -n postgresql logs postgres-0
kubectl -n postgresql exec -it postgres-0 -- bash
ls /data
ls /data/archive/

# check if our table was persisted!
psql --username=postgresadmin postgresdb
\dt
\q
```

That's it for chapter four! </br>
Now we've successfully managed to lift our PostgreSQL container and deploy it to Kubernetes. </br>
In the next chapter we'll take what we've learnt here and combine our previous studies to setup a Primary and Stand-By instance of PostgreSQL in Kubernetes </br>
