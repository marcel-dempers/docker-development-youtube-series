# Running PostgreSQL in Kubernetes (Primary and Standby)

In chapters [one](../1-introduction/README.md), [two](../2-configuration/README.md) and [three](../3-replication/README.md) we've managed to stand up a Primary and Stand-By PostgreSQL instances using containers. </br>

We've learnt the fundamentals of how to persist and store data, how to configure instances and how to setup streaming replication from a Primary container to a Stand-by container. </br>

In chapter [four](../4-k8s-basic/README.md) we lifted everything we've learnt so far into a Kubernetes cluster. </br>

We ended with a basic PostgreSQL running in Kubernetes with persistence, but without high availability. </br>

## The challenges - a reminder

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
kind create cluster --name postgresql --image kindest/node:v1.23.5

kubectl get nodes
NAME                       STATUS   ROLES                  AGE   VERSION
postgresql-control-plane   Ready    control-plane,master   31s   v1.23.5
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

## Review our first PostgreSQL instance

In [Chapter four](../4-k8s-basic/README.md) we created a basic PostgreSQL running in Kubernetes by combining everything we've learnt so far in the series annd creating a [statefulset.yaml](./yaml/statefulset.yaml) file. </br>

## Create our Primary PostgreSQL instance

Now what we will be doing is copying this file as it will be our starting point and call it `postgres-1.yaml`. </br>

We'll run through the file and will need to introduce a few things:


* Rename `postgres` to `postgres-1` to simbolize our first instance which will act as a Primary server
* Utilise `/docker-entrypoint-initdb.d/` initialization scripting
* Automate adding of the replication user

### Postgres Initialization

The Docker version of PostgreSQL has an initialization feature that will run any scripts in the folder `/docker-entrypoint-initdb.d/` as explained on [Docker Hub](https://hub.docker.com/_/postgres). </br>
It will only run when the data directory of PostgreSQL has not been initialized, making it perfect for things we want to setup only once at the first start and not upon every restart. </br>

To add this, we need a volume mount for out `init` step so the initContainer can do things and share files with the database container:

We will need this in both init and database container:
```
volumeMounts:
- mountPath: /docker-entrypoint-initdb.d
  name: initdb
```

Then we need a `volume` as an `emptyDir` which is just an empty directory shared between the two containers

```
volumes:
- name: initdb
  emptyDir: {}
```

Now our `initContainer` can generate a script that creates our replication user. PostgreSQL will run this on the first initialization of the data directory

```
#create a init template
echo "CREATE USER #REPLICATION_USER REPLICATION LOGIN ENCRYPTED PASSWORD '#REPLICATION_PASSWORD';" > init.sql

# add credentials
sed -i 's/#REPLICATION_USER/'${REPLICATION_USER}'/g' init.sql
sed -i 's/#REPLICATION_PASSWORD/'${REPLICATION_PASSWORD}'/g' init.sql

mkdir -p /docker-entrypoint-initdb.d/
cp init.sql /docker-entrypoint-initdb.d/init.sql
```

The init container will also need the replication user secrets to be mounted to environment variables:

```
env:
- name: REPLICATION_USER
  valueFrom:
    secretKeyRef:
      name: postgresql
      key: REPLICATION_USER
      optional: false
- name: REPLICATION_PASSWORD
  valueFrom:
    secretKeyRef:
      name: postgresql
      key: REPLICATION_PASSWORD
      optional: false
```

Now the `initContainer` will create our replication user!

## Deploy our Primary PostgreSQL instance

Deploy our first instance which will act as our primary:

```
kubectl -n postgresql apply -f storage/databases/postgresql/5-k8s-replication/yaml/postgres-1.yaml
```

## Check our installation

```
kubectl -n postgresql get pods

# check the initialization logs (should be clear)
kubectl -n postgresql logs postgres-1-0 -c init

# check the database logs
kubectl -n postgresql logs postgres-1-0

kubectl -n postgresql exec -it postgres-1-0 -- bash

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

# check the archive 
ls -l /data/archive
```

## Create our Standby Server

To create our Stand-By server, we will create a copy of `postgres-1.yaml` and rename it to `postgres-2.yaml`

Now there are a few things we'll do here:

* Carefully rename `postgres-1` to `postgres-2`
* Setup replication streaming from the Primary instance.

### Setting up our Stand by server automatically

As we learnt in previous chapters, a stand-by server is setup using `pg_basebackup` utility. </br>
We'll need to automate this as part of our initialization process. </br>

I'd like to keep our PostgreSQL YAML the same eventually so it could be packaged as a Helm chart. </br>

This will allow us to use the same YAML for Primary and Stand-by instances. </br>

To achieve this, we'll introduce a `STANDBY_MODE` variable which we can turn `"on"` if we want a Stand-by, and `"off"` if we want a Primary instance

```
if [ ${STANDBY_MODE} == "on" ]; 
then
  # initialize from backup if data dir is empty
  if [ -z "$(ls -A ${PGDATA})" ]; then
    export PGPASSWORD=${REPLICATION_PASSWORD}
    pg_basebackup -h ${PRIMARY_SERVER_ADDRESS} -p 5432 -U ${REPLICATION_USER} -D ${PGDATA} -Fp -Xs -R
  fi
else
  # previous logic here
fi
```

We we'll also need to introduce more environment variables

```
- name: STANDBY_MODE
  value: "on"
- name: PRIMARY_SERVER_ADDRESS
  value: "postgres-1"
- name: PGDATA
  value: "/data/pgdata"
```

## Deploy our Standby Server

Let's deploy the instance and see what happens:

```
kubectl -n postgresql apply -f storage/databases/postgresql/5-k8s-replication/yaml/postgres-2.yaml
```

Check our stand-by instance:

```
kubectl -n postgresql get pods

# check the initialization logs (you'll see `No such file or directory` because its not initialized yet!)
kubectl -n postgresql logs postgres-2-0 -c init

# check the database logs
kubectl -n postgresql logs postgres-2-0

kubectl -n postgresql exec -it postgres-2-0 -- bash

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

```

## Failover 

Now lets say `postgres-1` fails. </br>
PostgreSQL does not have built-in automated failver and recovery and requires tooling to perform this. </br>

When `postgres-1` fails, we would use a utility called [pg_ctl](https://www.postgresql.org/docs/current/app-pg-ctl.html) to promote our stand-by server to a new primary server. </br>

Then we have to build a new stand-by server just like we did in this guide. </br>
We would also need to configure replication on the new primary, the same way we did in this guide. </br>

Let's stop the primary server to simulate failure:

```
kubectl -n postgresql delete sts postgres-1

# notice the failure in replication from postgres-1 
 kubectl -n postgresql logs postgres-2-0
```

Then log into `postgres-2` and promote it to primary:
```
kubectl -n postgresql exec -it postgres-2-0 -- bash
psql --username=postgresadmin postgresdb

# confirm we cannot create a table as its a stand-by server
CREATE TABLE customers (firstname text, customer_id serial, date_created timestamp);

# quit out of postgresql
\q 

# run pg_ctl as postgres user (cannot be run as root!)
runuser -u postgres -- pg_ctl promote

# confirm we can create a table as its a primary server
CREATE TABLE customers (firstname text, customer_id serial, date_created timestamp);
```

## Setup Replication on the new Primary

Reconfigure the `postgres-2` instance as a primary:

```
env:
- name: STANDBY_MODE
  value: "off"
- name: PRIMARY_SERVER_ADDRESS
  value: ""
```


Deploy the changes:

```
kubectl -n postgresql apply -f storage/databases/postgresql/5-k8s-replication/yaml/postgres-2.yaml

# check our new instance
kubectl -n postgresql get pods
kubectl -n postgresql logs postgres-2-0 -c init
kubectl -n postgresql logs postgres-2-0
kubectl -n postgresql exec -it postgres-2-0 -- bash
```

If we wanted to deploy a new stand-by server, all we needed to do is create `postgres-3` with `STANDBY_MODE="on"` and set the `PRIMARY_SERVER_ADDRESS="postgres-2` and we would have a new stand-by server </br>

```
kubectl -n postgresql apply -f storage/databases/postgresql/5-k8s-replication/yaml/postgres-3.yaml
```

That's it for chapter chapter five! </br>
