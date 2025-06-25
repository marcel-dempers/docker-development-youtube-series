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

The following will start two shards of Prometheus, using the `hasmod` sharding technique. </br>
Basically two random services will be selected and scraped by each Prometheus shard. </br

```
cd sharding
docker compose up -d prometheus-00
docker compose up -d prometheus-01
```

# Start the Thanos Components 

The following will start two Thanos side cars, one for each Prometheus instance. 
It will also start a Minio S3 service
```
cd ../thanos
docker compose up
```