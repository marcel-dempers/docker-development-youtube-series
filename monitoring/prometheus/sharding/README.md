# Prometheus Basic Sharding
## a Video reference guide

To run any of the commands, please ensure you open a terminal and navigate to the path where this readme is located.

## Start example applications 

To start our example applications, we'll need to navigate to the folder where we have our example services in our docker-compose file 

```
cd monitoring\prometheus\

docker compose up -d --build go-application
docker compose up -d --build python-application
docker compose up -d --build dotnet-application
docker compose up -d --build nodejs-application
```

## Start Prometheus 

```
cd sharding
docker compose up -d prometheus-00
docker compose up -d prometheus-01
```