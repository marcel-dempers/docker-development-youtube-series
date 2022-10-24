# How to configure PostgreSQL

This is part 2 of our PostgreSQL series. </br>
In this chapter, we learn about fundamentals of the Postgres configuration. </br>

Many people make the mistakes of relying directly on Kubernetes PostgreSQL controllers
and Helm charts without having any understanding of Databases. </br>

Let's start where we left off, and review our simple PostgreSQL database:

## Run a simple PostgreSQL database (docker)

```
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