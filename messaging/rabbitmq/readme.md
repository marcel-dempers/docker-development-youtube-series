# RabbitMQ

Docker image over [here](https://hub.docker.com/_/rabbitmq)
```
docker network create rabbits
docker run -d --rm --net rabbits -e RABBITMQ_ERLANG_COOKIE=DSHEVCXBBETJJVJWTOWT --hostname rabbit-1 --name rabbit-1 rabbitmq:3.8
docker run -d --rm --net rabbits -e RABBITMQ_ERLANG_COOKIE=DSHEVCXBBETJJVJWTOWT --hostname rabbit-2 --name rabbit-2 rabbitmq:3.8
docker run -d --rm --net rabbits -e RABBITMQ_ERLANG_COOKIE=DSHEVCXBBETJJVJWTOWT --hostname rabbit-3 --name rabbit-3 rabbitmq:3.8

# how to grab existing erlang cookie
docker exec -it rabbit-1 cat /var/lib/rabbitmq/.erlang.cookie
DSHEVCXBBETJJVJWTOWT

docker rm -f rabbit-1
docker rm -f rabbit-2
docker rm -f rabbit-3
```

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

docker exec -it rabbit-1 rabbitmq-plugins enable rabbitmq_management
docker exec -it rabbit-2 rabbitmq-plugins enable rabbitmq_management
docker exec -it rabbit-3 rabbitmq-plugins enable rabbitmq_management

# Message Publisher

```
docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 publisher
```

# Message Consumer

```
docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest consumer
```
