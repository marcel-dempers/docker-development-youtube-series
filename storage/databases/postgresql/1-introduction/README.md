# Introduction to PostgreSQL

This is part 1 of my series on learning PostgreSQL. <br/>
The primary focus is getting PostgreSQL up and running and </br>
taking a first look at the database. </br>

PostgreSQL [Docker Image](https://hub.docker.com/_/postgres)

## Run a simple PostgreSQL database (docker)

```
 docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

## Run a simple PostgreSQL database (compose)

```
cd storage/databases/postgresql/1-introduction

docker compose up
```

We can access our database from the adminer web page on `http://localhost8080` </br>


When running containers, its always important to pull the image by tag </br>
to ensure you always get the same container image, else it will pull latest. </br>
We will do that in the next step. </br>

## Persisting Data

To persist data to PostgreSQL, we simply mount a docker volume. </br>
This is the way to persist container data. </br>
PostgreSQL stores its data by default under `/var/lib/postgresql/data` 
Also take note we are running a specific version of PostgreSQL now:

```
docker run -it --rm --name postgres `
  -e POSTGRES_PASSWORD=admin123 `
  -v ${PWD}/pgdata:/var/lib/postgresql/data `
  postgres:15.0
```

We can enter the container to connect to SQL:

```
# enter the container
docker exec -it postgres bash

# login to postgres
psql -h localhost -U postgres

#create a table
CREATE TABLE customers (firstname text,lastname text, customer_id serial);

#add record
INSERT INTO customers (firstname, lastname) VALUES ( 'Bob', 'Smith');

#show table
\dt

# get records
SELECT * FROM customers;

# quit 
\q

```

Now we can see our data persisted by killing and removing the container:

```
docker rm -f postgres
```

Run it again with the above `docker run` command and list our record with the above commands we've learnt </br>

## Networking

PostgreSQL by default uses port `5432`. </br>
Since we are running in Docker, we can bind a different port if we wish with Docker's `-p` flag. </br>
For example, we can expose port `5000` outside the container :

```
docker run -it --rm --name postgres `
  -e POSTGRES_PASSWORD=admin123 `
  -v ${PWD}/pgdata:/var/lib/postgresql/data `
  -p 5000:5432 `
  postgres:15.0
```
Note that this does not change the port which PostgreSQL runs on. </br>
To change that, we need to explore the configuration.

## Configuration 

PostgreSQL can be configured using environment variables as well as a config file. </br>

PostgreSQL has a ton of configuration options. </br>
In the next chapter, we will explore the configuration of PostgreSQL. </br>

## Docker Compose

Let's update our compose file to reflect our latest changes. </br>

We need to update the docker image, the port we want to expose outside of the container as well as a volume mount for persistence. </br>

```
version: '3.1'
services:
  db:
    image: postgres:15.0
    restart: always
    environment:
      POSTGRES_PASSWORD: admin123
    ports:
    - 5000:5432
    volumes:
    - ./pgdata:/var/lib/postgresql/data
  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```

That's it for chapter one! </br>
In [chapter 2](../2-configuration/README.md), we will take a look at Configuration and how to start our PostgreSQL instance with a custom configuration file. </br>
We will also explore the customization options available. </br>