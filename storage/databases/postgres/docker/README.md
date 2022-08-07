# notes

# docker image: https://hub.docker.com/_/postgres

```
cd storage\databases\postgres
```

```
docker run -it --rm --name postgres-1 `
  -e PGDATA=/var/lib/postgresql/data/pgdata `
  -e POSTGRES_DB=postgresdb `
  -e POSTGRES_USER=postgresadmin `
  -e POSTGRES_PASSWORD=admin123 `
  -p 5000:5432 `
  -v ${PWD}/postgres-1/pgdata/:/var/lib/postgresql/data/pgdata `
 postgres:14.4

```

```
# enter the container 
docker exec -it postgres-1 bash

# login to postgres
psql --username=postgresadmin postgresdb

#create a table
CREATE TABLE guestbook (visitor_email text, vistor_id serial, date timestamp, message text);

#add record
INSERT INTO guestbook (visitor_email, date, message) VALUES ( 'jim@gmail.com', current_date, 'This is a test.');

#show table
\dt

# quit 
\q

```