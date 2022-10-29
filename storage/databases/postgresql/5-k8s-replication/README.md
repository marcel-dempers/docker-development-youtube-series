# Running PostgreSQL in Kubernetes (Primary and Standby)


## Create a cluster

```

kind create cluster --name postgresql --image kindest/node:v1.23.5

kubectl get nodes
NAME                       STATUS   ROLES                  AGE   VERSION
postgresql-control-plane   Ready    control-plane,master   31s   v1.23.5

```

## Deploy our first PostgreSQL instance

Deploy a namespace to hold our resources: 

```
kubectl create ns postgresql
```

Create our secret for our first PostgreSQL instance:

```
kubectl -n postgresql create secret generic postgresql `
  --from-literal POSTGRES_USER="postgresadmin" `
  --from-literal POSTGRES_PASSWORD='admin123' `
  --from-literal POSTGRES_DB="postgresdb" `
  --from-literal REPLICATION_USER="replicationuser" `
  --from-literal REPLICATION_PASSWORD='replicationPassword'
```

Deploy our first instance which will act as our primary:

```
kubectl -n postgresql apply -f storage/databases/postgresql/4-kubernetes/yaml/postgres-1.yaml
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

## Deploy our Standby Server

```
kubectl -n postgresql apply -f storage/databases/postgresql/4-kubernetes/yaml/postgres-2.yaml
```


runuser -u postgres -- pg_ctl reload



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

```
  #replication
  wal_level = replica
  archive_mode = on
  archive_command = 'test ! -f /data/archive/%f && cp %p /data/archive/%f'
  max_wal_senders = 3
```

Reconfigure the `postgres-2` instance:

```
kubectl -n postgresql apply -f storage/databases/postgresql/4-kubernetes/yaml/postgres-2.yaml

# check our new instance
kubectl -n postgresql get pods
kubectl -n postgresql logs postgres-2-0 -c init
kubectl -n postgresql logs postgres-2-0
kubectl -n postgresql exec -it postgres-2-0 -- bash
```

That's it for chapter three! </br>
Now we understand how to [run PostgreSQL](../1-introduction/README.md), how to [configure PostgreSQL](../2-configuration/README.md) and how to setup replication for better availability.
