# RabbitMQ

<a href="https://youtu.be/hfUIWe1tK8E" title="rabbitmq-intro"><img src="https://i.ytimg.com/vi/hfUIWe1tK8E/hqdefault.jpg" width="20%" alt="rabbitmq-intro" /></a> 

Docker image over [here](https://hub.docker.com/_/rabbitmq)
```
# run a standalone instance
docker network create rabbits
docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 rabbitmq:3.8 

# how to grab existing erlang cookie
docker exec -it rabbit-1 cat /var/lib/rabbitmq/.erlang.cookie

# clean up
docker rm -f rabbit-1
```

# Management

```
docker run -d --rm --net rabbits -p 8080:15672 -e RABBITMQ_ERLANG_COOKIE=DSHEVCXBBETJJVJWTOWT --hostname rabbit-manager --name rabbit-manager rabbitmq:3.8-management

#join the manager

docker exec -it rabbit-manager rabbitmqctl stop_app
docker exec -it rabbit-manager rabbitmqctl reset
docker exec -it rabbit-manager rabbitmqctl join_cluster rabbit@rabbit-1
docker exec -it rabbit-manager rabbitmqctl start_app
docker exec -it rabbit-manager rabbitmqctl cluster_status
```

# Enable Statistics

```
docker exec -it rabbit-1 rabbitmq-plugins enable rabbitmq_management
docker exec -it rabbit-2 rabbitmq-plugins enable rabbitmq_management
docker exec -it rabbit-3 rabbitmq-plugins enable rabbitmq_management
```

# Message Publisher

```

cd messaging\rabbitmq\applications\publisher
docker build . -t aimvector/rabbitmq-publisher:v1.0.0

docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 aimvector/rabbitmq-publisher:v1.0.0
```

# Message Consumer

```

docker build . -t aimvector/rabbitmq-consumer:v1.0.0
docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest aimvector/rabbitmq-consumer:v1.0.0
```

# Clustering 

https://www.rabbitmq.com/cluster-formation.html

## Note

Remember we will need the Erlang Cookie to allow instances to authenticate with each other.

# Manual Clustering 

```

docker exec -it rabbit-1 rabbitmqctl cluster_status

#join node 2

docker exec -it rabbit-2 rabbitmqctl stop_app
docker exec -it rabbit-2 rabbitmqctl reset
docker exec -it rabbit-2 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec -it rabbit-2 rabbitmqctl start_app
docker exec -it rabbit-2 rabbitmqctl cluster_status

#join node 3
docker exec -it rabbit-3 rabbitmqctl stop_app
docker exec -it rabbit-3 rabbitmqctl reset
docker exec -it rabbit-3 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec -it rabbit-3 rabbitmqctl start_app
docker exec -it rabbit-3 rabbitmqctl cluster_status

```

# Automated Clustering

```
docker run -d --rm --net rabbits `
-v ${PWD}/config/rabbit-1/:/config/ `
-e RABBITMQ_CONFIG_FILE=/config/rabbitmq `
-e RABBITMQ_ERLANG_COOKIE=WIWVHCDTCIUAWANLMQAW `
--hostname rabbit-1 `
--name rabbit-1 `
-p 8081:15672 `
rabbitmq:3.8-management

docker run -d --rm --net rabbits `
-v ${PWD}/config/rabbit-2/:/config/ `
-e RABBITMQ_CONFIG_FILE=/config/rabbitmq `
-e RABBITMQ_ERLANG_COOKIE=WIWVHCDTCIUAWANLMQAW `
--hostname rabbit-2 `
--name rabbit-2 `
-p 8082:15672 `
rabbitmq:3.8-management

docker run -d --rm --net rabbits `
-v ${PWD}/config/rabbit-3/:/config/ `
-e RABBITMQ_CONFIG_FILE=/config/rabbitmq `
-e RABBITMQ_ERLANG_COOKIE=WIWVHCDTCIUAWANLMQAW `
--hostname rabbit-3 `
--name rabbit-3 `
-p 8083:15672 `
rabbitmq:3.8-management

#NODE 1 : MANAGEMENT http://localhost:8081
#NODE 2 : MANAGEMENT http://localhost:8082
#NODE 3 : MANAGEMENT http://localhost:8083

# enable federation plugin
docker exec -it rabbit-1 rabbitmq-plugins enable rabbitmq_federation 
docker exec -it rabbit-2 rabbitmq-plugins enable rabbitmq_federation
docker exec -it rabbit-3 rabbitmq-plugins enable rabbitmq_federation

```

# Basic Queue Mirroring 

```
docker exec -it rabbit-1 bash

# https://www.rabbitmq.com/ha.html#mirroring-arguments

rabbitmqctl set_policy ha-fed \
    ".*" '{"federation-upstream-set":"all", "ha-mode":"nodes", "ha-params":["rabbit@rabbit-1","rabbit@rabbit-2","rabbit@rabbit-3"]}' \
    --priority 1 \
    --apply-to queues
```

# Automatic Synchronization

https://www.rabbitmq.com/ha.html#unsynchronised-mirrors

```
rabbitmqctl set_policy ha-fed \
    ".*" '{"federation-upstream-set":"all", "ha-sync-mode":"automatic", "ha-mode":"nodes", "ha-params":["rabbit@rabbit-1","rabbit@rabbit-2","rabbit@rabbit-3"]}' \
    --priority 1 \
    --apply-to queues
```

# Further Reading

https://www.rabbitmq.com/ha.html


# Clean up

```
docker rm -f rabbit-1
docker rm -f rabbit-2
docker rm -f rabbit-3
```

# RabbitMQ on Kubernetes

Checkout the Kubernetes walkthrough [here](./kubernetes/readme.md)
