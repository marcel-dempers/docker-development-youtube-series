# Prometheus Application Monitoring
## a Video reference guide

To run any of the commands, please ensure you open a terminal and navigate to the path where this readme is located.

## Start Prometheus

```
docker-compose up -d prometheus
docker-compose up -d grafana
```

Wait for Grafana to start up
Import the dashboards
```
TODO
```

You should see all application targets un `UNKNOWN` or  `DOWN` status.
```http://localhost:9090/targets```

## Start the example app you prefer

```
docker-compose up -d golang-application
docker-compose up -d python-application
docker-compose up -d dotnet-application
docker-compose up -d nodejs-application
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