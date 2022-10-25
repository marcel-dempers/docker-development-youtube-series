# How to configure PostgreSQL

This is part 2 of our PostgreSQL series. </br>
In this chapter, we learn about fundamentals of the Postgres configuration. </br>

Many people make the mistakes of relying directly on Kubernetes PostgreSQL controllers
and Helm charts without having any understanding of Databases. </br>

Let's start where we left off, and review our simple PostgreSQL database:

## Run a simple PostgreSQL database (docker)

```
cd storage/databases/postgres/2-configuration
docker run -it --rm --name postgres `
  -e POSTGRES_PASSWORD=admin123 `
  -v ${PWD}/pgdata:/var/lib/postgresql/data `
  -p 5000:5432 `
  postgres:15.0
```

## Environment Variables

Many settings can be specified using environment variables. </br>
I generally recommend not relying on default values and set most of the settings 
possible. </br>

I personally prefer most or all settings in a configuration file, so it can be committed to source control. </br>
This is where Environment variables are great because we can inject secrets there
and keep passwords out of our configuration files and out of source control. </br>

This will be important in Kubernetes later on. </br>

We will not learn all or even most of the configurations in this chapter, as PostgreSQL has a lot of depth. So we will only learn what we need, one step at a time. </br>

Let's take a look at some basic configurations [here](https://hub.docker.com/_/postgres)

Let's set a few things here:

| Environment Variable | Meaning |
|----------------------|---------|
| POSTGRES_USER        |  Username for the Postgres Admin       |
| POSTGRES_PASSWORD    |  Password for the Postgres Admin       |
| POSTGRES_DB          |  Default database for your Postgres Server       |
| PGDATA               |  Path where data is stored       |


## Configuration files

If we take a look at our `docker` mount that we defined in our `docker run` command: </br>

`-v ${PWD}/pgdata:/var/lib/postgresql/data ` </br>

The `{PWD}/pgdata` folder that we have mounted contains not only data, but some defaut configuration files that we can explore. </br>

Three files are important here:

 

|Configuration file | Meaning |  Documentation
|----------------------|---------|-------|
| pg_hba.conf        |  Host Based Authentication file       | [Official Documentation](https://www.postgresql.org/docs/current/auth-pg-hba-conf.html) |
| pg_ident.conf    |  User Mappings file       |  [Official Documentation](https://www.postgresql.org/docs/current/auth-username-maps.html)
| postgresql.conf          |  PostgreSQL main configuraiton       |

## The pg_hba.conf File

We'll start this guide with the host based authentication file. </br>
This file is automatically created in the data directory as we see. </br>
We should create a copy of this file and configure it ourselves. </br>

It controls who can access our PostgreSQL server. </br>
Let's refer to the official documentation as well as walk through the config. </br>
The config file itself has a great description of the contents. </br>

As mentioned in the previous chapter, it's always good not to rely on default configurations. So let's create our own `pg_hba.conf` file. </br>

We can grab the content from the default configuration and we may edit it as we go.

```
# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
# IPv6 local connections:
host    all             all             ::1/128                 trust
# Allow replication connections from localhost, by a user with the
# replication privilege.
local   replication     all                                     trust
host    replication     all             127.0.0.1/32            trust
host    replication     all             ::1/128                 trust

host all all all scram-sha-256
```

## The pg_ident.conf File

This config file is a mapping file between system users and database users. </br>
Let's refer to the official documentation and walk through the config. </br>
This is not a feature that we will need in this series, so we will skip this config for the time being. </br>

## The postgresql.conf File

This configuration file is the main one for PostgreSQL. </br>
As you can see this is a large file with in-depth tuning and customization capability. </br>

```
docker run -d --rm --name postgres-1 `
  --net postgres `
  -e POSTGRES_USER=postgresadmin `
  -e POSTGRES_PASSWORD=admin123 `
  -e POSTGRES_DB=postgresdb `
  -e PGDATA=/var/lib/postgresql/data/pgdata `
  -p 5000:5432 `
  -v ${PWD}/archive:/mnt/server/archive `
  -v ${PWD}/postgres-1/pgdata:/var/lib/postgresql/data/pgdata `
  -v ${PWD}/postgres-1/postgresql.conf:/etc/postgresql/postgresql.conf `
  -v ${PWD}/postgres-1/pg_hba.conf:/var/lib/postgresql/data/pgdata/pg_hba.conf `
  postgres:14.4 -c 'config_file=/etc/postgresql/postgresql.conf'
```