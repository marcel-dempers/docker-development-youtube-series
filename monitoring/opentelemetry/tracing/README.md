# Tracing with OpenTelemetry


<a href="https://youtu.be/bIxt1b0GOU4" title="fluentd-intro"><img src="https://i.ytimg.com/vi/bIxt1b0GOU4/hqdefault.jpg" width="20%" alt="opentel-intro" /></a> 

<a href="https://youtu.be/LQOeaxfiAt8" title="fluentd-intro"><img src="https://i.ytimg.com/vi/LQOeaxfiAt8/hqdefault.jpg" width="20%" alt="opentel-intro" /></a> 


## Build the applications 

To build the applications navigate your terminal to the location of the compose file

```
cd monitoring/opentelemetry/
```

Build: 

```
docker compose build
```

## Docker

Start up the stack using local docker compose, by running:

```
docker compose up
```

## Kubernetes

For Kubernetes, we will use the OpenTelemetry Operator. </br>
See that guide over [here](../kubernetes/README.md)


## Generate from Traffic

You can generate some traffic by visiting the application on http://localhost

<i>Note: You will need port 80 and 81 open for the applications to function</i>

## Access Grafana

You can access Grafana on http://localhost:3000 to test the Tempo datasource.

