# How to replicate PostgreSQL

So far we've learnt how to run PostgreSQL as per [chapter 1](../1-introduction/README.md). </br>
We've also learnt how to configure PostgreSQL with our custom configuration file as per [chapter 2](../2-configuration/README.md). </br>

In this chapter we will setup a second PostgreSQL instance. </br>
Then we will learn how to configure our first PostgreSQL instance to replicate its data to the new second instance. </br>
This will essentially give us a primary and a secondary server for better availability in case we lose our primary server. </br>


* Documentation on HA \ replication \ failover modes


## Get our Primary PostgreSQL up and running

Let's start by running our PostgreSQL in docker </br>

Few things to note here: </br>
* We start our instance with a different name to identify it as the first instance with the `--name postgres-1` flag and `2` for the second instance
* Set unique data volumes for data between instances
* Set unique config files for each instance
* Create and run our docker containers on the same network 

Create a new network so our instances can talk with each other:

```
docker network create postgres
```

Start with instance 1: 

```
cd storage/databases/postgres/3-replication

docker run -it --rm --name postgres-1 `
--net postgres `
-e POSTGRES_USER=postgresadmin `
-e POSTGRES_PASSWORD=admin123 `
-e POSTGRES_DB=postgresdb `
-e PGDATA="/data" `
-v ${PWD}/postgres-1/pgdata:/data `
-v ${PWD}/postgres-1/config:/config `
-p 5000:5432 `
postgres:15.0 -c 'config_file=/config/postgresql.conf'
```

Start instance 2:

```
cd storage/databases/postgres/3-replication

docker run -it --rm --name postgres-2 `
--net postgres `
-e POSTGRES_USER=postgresadmin `
-e POSTGRES_PASSWORD=admin123 `
-e POSTGRES_DB=postgresdb `
-e PGDATA="/data" `
-v ${PWD}/postgres-2/pgdata:/data `
-v ${PWD}/postgres-2/config:/config `
-p 5000:5432 `
postgres:15.0 -c 'config_file=/config/postgresql.conf'
```