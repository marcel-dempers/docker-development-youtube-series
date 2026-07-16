# Prometheus Application Monitoring
## a Video reference guide

To run any of the commands, please ensure you open a terminal and navigate to the path where this readme is located.

## Start Prometheus, Grafana & Dashboards

```
docker compose up -d prometheus
docker compose up -d grafana
docker compose up -d grafana-dashboards
```


## Start the example app you prefer

```
docker compose up -d --build go-application
docker compose up -d --build python-application
docker compose up -d --build dotnet-application
docker compose up -d --build nodejs-application
```

## Generate some requests by opening the application in the browser

```
http://localhost:80 #Golang
http://localhost:81 #Python
http://localhost:82 #Dotnet
http://localhost:83 #NodeJS
```

## Check Dashboards
```
http://localhost:3000

```
## Prometheus Queries
### Golang Examples

Requests per Second over 2minutes
```
irate(go_request_operations_total[2m])
```
Request duration
```
rate(go_request_duration_seconds_sum[2m]) / rate(go_request_duration_seconds_total[2m])
```

# Prometheus Guide on Kubernetes

Checkout the prometheus guide [here](./kubernetes/readme.md)