# Persistent Volumes Demo

<a href="https://youtu.be/ZxC6FwEc9WQ" title="k8s-pv"><img src="https://i.ytimg.com/vi/ZxC6FwEc9WQ/hqdefault.jpg" width="20%" alt="k8s-pv" /></a> 

## Container Storage

By default containers store their data on the file system like any other process.
Container file system is temporary and not persistent during container restarts
When container is recreated, so is the file system



```
# run postgres
docker run -d --rm -e POSTGRES_DB=postgresdb -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 postgres:10.4

# enter the container 
docker exec -it <container-id> bash

# login to postgres
psql --username=admin postgresdb

#create a table
CREATE TABLE COMPANY(
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   AGE            INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL
);

#show table
\dt

# quit 
\q
```

Restarting the above container and going back in you will notice `\dt` commands returning no tables.
Since data is lost.

Same can be demonstrated using Kubernetes

```
cd .\kubernetes\persistentvolume\

kubectl create ns postgres
kubectl apply -n postgres -f ./postgres-no-pv.yaml
kubectl -n postgres get pods 
kubectl -n postgres exec -it postgres-0 bash

# run the same above mentioned commands to create and list the database table

kubectl delete po -n postgres postgres-0

# exec back in and confirm table does not exist.
```

# Persist data Docker

```
docker volume create postges
docker run -d --rm -v postges:/var/lib/postgresql/data -e POSTGRES_DB=postgresdb -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 postgres:10.4

# run the same tests as above and notice
```

# Persist data Kubernetes


```
kubectl apply -f persistentvolume.yaml
kubectl apply -n postgres -f persistentvolumeclaim.yaml

kubectl apply -n postgres -f postgres-with-pv.yaml

kubectl -n postgres get pods

```
