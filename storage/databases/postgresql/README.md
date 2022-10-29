
# notes

/usr/share/postgresql/postgresql.conf.sample
SHOW hba_file
pg_hba.conf

wal_keep_size    or by configuring a replication slot for the standby

```

cd storage/databases/postgresqlql/docker
docker network create postgres

docker run -d --rm --name postgres-1 `
  --net postgres `
  -e PGDATA=/var/lib/postgresql/data/pgdata `
  -e POSTGRES_DB=postgresdb `
  -e POSTGRES_USER=postgresadmin `
  -e POSTGRES_PASSWORD=admin123 `
  -p 5000:5432 `
  -v ${PWD}/archive:/mnt/server/archive `
  -v ${PWD}/postgres-1/pgdata:/var/lib/postgresql/data/pgdata `
  -v ${PWD}/postgres-1/postgresql.conf:/etc/postgresql/postgresql.conf `
  -v ${PWD}/postgres-1/pg_hba.conf:/var/lib/postgresql/data/pgdata/pg_hba.conf `
  postgres:14.4 -c 'config_file=/etc/postgresql/postgresql.conf'


docker run -d --rm --name postgres-2 `
  --net postgres `
  -e PGDATA=/var/lib/postgresql/data/pgdata `
  -e POSTGRES_DB=postgresdb `
  -e POSTGRES_USER=postgresadmin `
  -e POSTGRES_PASSWORD=admin123 `
  -p 5001:5432 `
  -v ${PWD}/backup:/var/lib/backup `
  -v ${PWD}/postgres-2/pgdata:/var/lib/postgresql/data/pgdata `
  -v ${PWD}/postgres-2/postgresql.conf:/etc/postgresql/postgresql.conf `
  postgres:14.4 -c 'config_file=/etc/postgresql/postgresql.conf'
```  


# replication setup 

## user for replication
https://www.postgresql.org/docs/current/app-createuser.html

```
# enter the container 
docker exec -it postgres-1 bash
createuser -U postgresadmin -P -c 5 --replication replicationUser
psql --username=postgresadmin postgresdb
\du
```

## allow connection for that user in pg_hba.conf

```
# Allow replication connections
host     replication     replicationUser         0.0.0.0/0        md5
```

## setup replication

```
docker exec -it postgres-1 bash

wal_level = replica
archive_mode = on
archive_command = 'test ! -f /mnt/server/archive/%f && cp %p /mnt/server/archive/%f'
max_wal_senders = 3
```
```

docker exec -it postgres-2 bash
SHOW data_directory;
pg_basebackup -h postgres-1 -p 5432 -U replicationUser -D /var/lib/backup/ -Fp -Xs -R
```

# create sample data
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



docker run -d --rm --name postgres-2 `
  --net postgres `
  -e PGDATA=/var/lib/postgresql/data/pgdata `
  -e POSTGRES_DB=postgresdb `
  -e POSTGRES_USER=postgresadmin `
  -e POSTGRES_PASSWORD=admin123 `
  -p 5001:5432 `
  -v ${PWD}/backup:/var/lib/postgresql/data/pgdata `
  postgres:14.4 -c 'config_file=/var/lib/postgresql/data/pgdata/postgresql.conf'