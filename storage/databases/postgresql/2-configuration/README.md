# How to configure PostgreSQL

This is part 2 of our PostgreSQL series. </br>
In this chapter, we learn about fundamentals of the Postgres configuration. </br>

Many people make the mistakes of relying directly on Kubernetes PostgreSQL controllers
and Helm charts without having any understanding of Databases. </br>

Let's start where we left off, and review our simple PostgreSQL database:

## Run a simple PostgreSQL database (docker)

```
cd storage/databases/postgresql/2-configuration
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

The `{PWD}/pgdata` folder that we have mounted contains not only data, but some default configuration files that we can explore. </br>

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

### File Locations

Let's set our data directory locations as well as config file locations </br>
Our volume mount path in the container is also short and simple. </br>
Note that we also split config from data so we have separate paths :

```
data_directory = '/data'
hba_file = '/config/pg_hba.conf'
ident_file = '/config/pg_ident.conf'
```

### Connection and Authentication

The shared_buffers parameter determines how much memory is dedicated to the server for caching data. The value should be set to 15% to 25% of the machine's total RAM. For example: if your machine's RAM size is 32 GB, then the recommended value for shared_buffers is 8 GB </br>

We will take a look at `WAL` (Write Ahead Log), Archiving, Primary, and Standby configurations in a future chapter on replication </br>

```
port = 5432
listen_addresses = '*'
max_connections = 100
shared_buffers = 128MB
dynamic_shared_memory_type = posix
max_wal_size = 1GB
min_wal_size = 80MB
log_timezone = 'Etc/UTC'
datestyle = 'iso, mdy'
timezone = 'Etc/UTC'

#locale settings
lc_messages = 'en_US.utf8'			# locale for system error message
lc_monetary = 'en_US.utf8'			# locale for monetary formatting
lc_numeric = 'en_US.utf8'			# locale for number formatting
lc_time = 'en_US.utf8'				# locale for time formatting

default_text_search_config = 'pg_catalog.english'

```

We can also include other configurations from other locations with the `include_dir` and `include` options. </br>
We will skip these for the sake of keeping things simple. </br>
Nested configurations can over complicate a setup and makes it hard to troubleshoot when issues occur. </br>

### Specifying Custom Configuration

If we run on Linux, we need to ensure that the `postgres` user which has a user ID of `999` by default, should have access to the configuration files. </br>

```
sudo chown 999:999 config/postgresql.conf
sudo chown 999:999 config/pg_hba.conf
sudo chown 999:999 config/pg_ident.conf
```

There is another important gotcha here. </br>
The `PGDATA` variable tells PostgreSQL where our data directory is. </br>
Similarly, we've learnt that our configuration file also has `data_directory` which tells PostgreSQL the same. </br>

However, the latter is only read by PostgreSQL after initialization has occurred. </br>
PostgreSQL's initialization phase sets up directory permissions on the data directory. </br>
If we leave out `PGDATA`, then we will get errors that the data directory is invalid. </br>
Hence `PGDATA` is important here. </br>

## Running our PostgreSQL

Finally, we can run our database with our custom configuration files:

```
docker run -it --rm --name postgres `
-e POSTGRES_USER=postgresadmin `
-e POSTGRES_PASSWORD=admin123 `
-e POSTGRES_DB=postgresdb `
-e PGDATA="/data" `
-v ${PWD}/pgdata:/data `
-v ${PWD}/config:/config `
-p 5000:5432 `
postgres:15.0 -c 'config_file=/config/postgresql.conf'
```

That's it for chapter two! </br>
In [chapter 3](../3-replication/README.md), we will take a look at Replication and how to replicate our data to another PostgreSQL instance for better availability.