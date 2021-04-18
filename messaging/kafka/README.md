# Notes 
https://hub.docker.com/_/openjdk?tab=description&page=1&ordering=last_updated&name=alpine
https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-debian-10

# Building a Docker file
docker run --rm --name kafka -it kafka bash

docker run --rm -it kafka bash -c "ls -l /kafka/"
docker run --rm -it kafka bash -c "cat ~/kafka/config/server.properties"
docker run --rm -it kafka bash -c "ls -l ~/kafka/bin"

docker cp kafka:/kafka/config/server.properties ./server.properties
docker cp kafka:/kafka/config/zookeeper.properties ./zookeeper/zookeeper.properties

# Kafka

docker network create kafka
docker run -it --rm --name kafka --net kafka -v ${PWD}/server.properties:/kafka/config/server.properties kafka

# Zookeeper

docker run -it --rm --name zookeeper --net kafka zookeeper

# Topic
docker exec -it kafka bash

/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic TutorialTopic

# Producer

echo "Hello, World" | /kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic TutorialTopic > /dev/null

# Consumer

/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic TutorialTopic --from-beginning

# Build an Application: Producer

https://docs.confluent.io/clients-confluent-kafka-go/current/overview.html#go-installation

```
cd messaging/kafka/applications/producer

docker run -it --rm -v ${PWD}:/app -w /app golang:1.15-alpine

apk -U add ca-certificates && \
apk update && apk upgrade && apk add pkgconf git bash build-base && \
cd /tmp && \
git clone https://github.com/edenhill/librdkafka.git && \
  cd librdkafka && \
  git checkout v1.6.1 && \ 
  ./configure --prefix /usr && make && make install

#apk add --no-cache git make librdkafka-dev gcc musl-dev librdkafka

go mod init producer
go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
```

