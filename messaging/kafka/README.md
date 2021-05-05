# Introduction to Kafka

Official [Docs](https://kafka.apache.org/)

## Building a Docker file

As always, we start with a `dockerfile` </br>
We can build our `dockerfile`

```
cd .\messaging\kafka\
docker build . -t aimvector/kafka:2.7.0

```

## Exploring the Kafka Install

We can then run it to explore the contents:

```
docker run --rm --name kafka -it aimvector/kafka:2.7.0 bash

ls -l /kafka/
cat /kafka/config/server.properties
ls -l /kafka/bin
```

We can use the `docker cp` command to copy the file out of our container:

```
docker cp kafka:/kafka/config/server.properties ./server.properties
docker cp kafka:/kafka/config/zookeeper.properties ./zookeeper/zookeeper.properties
```

We'll need the Kafka configuration to tune our server and Kafka also requires
at least one Zookeeper instance in order to function. To achieve high availability, we'll run
multiple kafka as well as multiple zookeeper instances in the future

# Zookeeper

Let's build a Zookeeper image. The Apache folks have made it easy to start a Zookeeper instance the same way as the Kafka instance by simply running the `start-zookeeper.sh` script.

```
cd .\messaging\kafka\zookeeper
docker build . -t aimvector/zookeeper:2.7.0

```

Let's create a kafka network and run 1 zookeeper instance

```
docker network create kafka
docker run -d --rm --name zookeeper --net kafka zookeeper
```

# Kafka - 1

```
docker run -d `
--rm `
--name kafka-1 `
--net kafka `
-v ${PWD}/config/kafka-1/server.properties:/kafka/config/server.properties `
aimvector/kafka:2.7.0
```

# Kafka - 2

```
docker run -d `
--rm `
--name kafka-2 `
--net kafka `
-v ${PWD}/config/kafka-2/server.properties:/kafka/config/server.properties `
aimvector/kafka:2.7.0
```

# Kafka - 3

```
docker run -d `
--rm `
--name kafka-3 `
--net kafka `
-v ${PWD}/config/kafka-3/server.properties:/kafka/config/server.properties `
aimvector/kafka:2.7.0
```


# Topic

Let's create a Topic that allows us to store `Order` information. </br>
To create a topic, Kafka and Zookeeper have scripts with the installer that allows us to do so. </br>

Access the container:
```
docker exec -it zookeeper bash
```
Create the Topic:
```
/kafka/bin/kafka-topics.sh \
--create \
--zookeeper zookeeper:2181 \
--replication-factor 1 \
--partitions 3 \
--topic Orders
```

Describe our Topic:
```
/kafka/bin/kafka-topics.sh \
--describe \
--topic Orders \
--zookeeper zookeeper:2181
```

We can take a look at how Kafka stores data

```
apt install -y tree
tree /tmp/kafka-logs/
```

# Simple Producer & Consumer

The Kafka installation also ships with a script that allows us to produce
and consume messages to our Kafka network:

```
echo "New Order: 1" | \
/kafka/bin/kafka-console-producer.sh \
--broker-list kafka-1:9092 \
--topic Orders > /dev/null
```

We can then run the consumer that will receive that message on that Orders topic:

```
/kafka/bin/kafka-console-consumer.sh \
--bootstrap-server kafka-1:9092 \
--topic Orders --from-beginning

```

Once we have a message in Kafka, we can explore where it got stored in which partition:

```
ls -lh /tmp/kafka-logs/Orders-*

/tmp/kafka-logs/Orders-0:
total 4.0K
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.index    
-rw-r--r-- 1 root root   0 May  4 06:54 00000000000000000000.log      
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.timeindex
-rw-r--r-- 1 root root   8 May  4 06:54 leader-epoch-checkpoint       

/tmp/kafka-logs/Orders-1:
total 4.0K
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.index    
-rw-r--r-- 1 root root   0 May  4 06:54 00000000000000000000.log      
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.timeindex
-rw-r--r-- 1 root root   8 May  4 06:54 leader-epoch-checkpoint       

/tmp/kafka-logs/Orders-2:
total 8.0K
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.index    
-rw-r--r-- 1 root root  80 May  4 06:57 00000000000000000000.log      
-rw-r--r-- 1 root root 10M May  4 06:54 00000000000000000000.timeindex
-rw-r--r-- 1 root root   8 May  4 06:54 leader-epoch-checkpoint
```

By seeing 0 bytes in partition 0 and 1, we know the message is sitting in partition 2 as it has 80 bytes. </br>
We can check the message with :

```
cat /tmp/kafka-logs/Orders-2/*.log
```

## Building a Producer: Go

```
docker run -it `
--net kafka `
-e KAFKA_PEERS="kafka-1:9092,kafka-2:9092,kafka-3:9092" `
-e KAFKA_TOPIC="Orders" `
-e KAFKA_PARTITION=1 `
-p 80:80 `
kafka-producer
```

## Building a Consumer: Go

```
cd messaging\kafka\applications\consumer
docker build . -t kafka-consumer

docker run -it `
--net kafka `
-e KAFKA_PEERS="kafka-1:9092,kafka-2:9092,kafka-3:9092" `
-e KAFKA_TOPIC="Orders" `
kafka-consumer
```

# High Availability + Replication

Next up, we'll take a look at achieving high availability using replication techniques
and taking advantage of Kafka's distributed architecture.