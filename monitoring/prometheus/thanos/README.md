# Prometheus Thanos Basics

This guide relies on our Prometheus Application Monitoring Guide. </br>

## Start example applications 

To start our example applications, we'll need to navigate to the folder where we have our example services in our docker-compose file 

```
cd monitoring\prometheus\

docker compose up -d --build go-application
docker compose up -d --build python-application
docker compose up -d --build dotnet-application
docker compose up -d --build nodejs-application
```

## Start Sharded Prometheus Instances

```
cd sharding
docker compose up -d prometheus-00
docker compose up -d prometheus-01
```