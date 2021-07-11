# Consuming Data in Kafka

* [Start a Kafka environment](#Start-a-Kafka-environment)
* [Building a consumer in Go](#Building-a-consumer-in-Go)
* [Starting our Kafka Consumer](#Starting-our-Kafka-Consumer)
* [Consume Messages in Random Order](#Consume-Messages-in-Random-Order)
* [Consume Messages in Order](#Consume-Messages-in-Order)

## Start a Kafka environment

Let's start our Kafka components:

```
cd messaging/kafka

#only start the kafka containers, not everything!
docker compose up zookeeper-1 kafka-1 kafka-2 kafka-3

#ensure its running!
docker ps
```

### Create a Topic: Orders

To create a topic, we can `exec` into any container on our kafka network and create it. </br>
We'll need a Topic for our orders:

```
docker exec -it zookeeper-1 bash

# create
/kafka/bin/kafka-topics.sh \
--create \
--zookeeper zookeeper-1:2181 \
--replication-factor 1 \
--partitions 3 \
--topic Orders

# describe
/kafka/bin/kafka-topics.sh \
--describe \
--zookeeper zookeeper-1:2181 \
--topic Orders

exit
```

Now that we have our Kafka infrastructure, let's create our consumer.

## Building a consumer in Go
Example code: [Sarama Library](https://github.com/Shopify/sarama/blob/master/examples/consumergroup/main.go)


We start with `main.go`:

```
package main

import (
  "context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"gopkg.in/Shopify/sarama.v1"
)

var kafkaBrokers = os.Getenv("KAFKA_PEERS")
var kafkaTopic = os.Getenv("KAFKA_TOPIC")
var kafkaVersion = os.Getenv("KAFKA_VERSION")
var kafkaGroup = os.Getenv("KAFKA_GROUP")


func main() {
}
```

Above we define our dependencies, and connection details for kafka. <br/>

Let's define our client config:

```
  version, err := sarama.ParseKafkaVersion(kafkaVersion)
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	
  ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(kafkaBrokers, ","), kafkaGroup, config)

  if err != nil {
	  fmt.Printf("Failed to init Kafka consumer group: %s", err)
		panic(err)
	}

```

Now that we have a client and consumer group connection with Kafka,
we can proceed to define our consumer:

```
// in main.go
type Consumer struct {
	ready chan bool
}

//in main()
consumer := Consumer{
		ready: make(chan bool),
}
```

With the consumer we can now define a wait group which will consume messages
as they come in:

```
wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(kafkaTopic, ","), &consumer); err != nil {
				fmt.Printf("Error from consumer: %v", err)
				panic(err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()
  <-consumer.ready // Await till the consumer has been set up
  fmt.Println("Sarama consumer up and running!...")

```

In case we need to exit, let's handle exit signals

```
sigterm := make(chan os.Signal, 1)
signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
select {
case <-ctx.Done():
  fmt.Println("terminating: context cancelled")
case <-sigterm:
  fmt.Println("terminating: via signal")
}
cancel()
wg.Wait()

if err = client.Close(); err != nil {
  fmt.Printf("Error closing client: %v", err)
  panic(err)
}
```

In addition to the stuct we created, the Sarama library needs us to handle certain functions:


```
// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
```

And finally a ConsumeClaim() to handle the messages coming in:

```
// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for msg := range claim.Messages() {

		fmt.Printf("Partition:\t%d\n", msg.Partition)
		fmt.Printf("Offset:\t%d\n", msg.Offset)
		fmt.Printf("Key:\t%s\n", string(msg.Key))
		fmt.Printf("Value:\t%s\n", string(msg.Value))
		fmt.Printf("Topic:\t%s\n", msg.Topic)
		fmt.Println()

		session.MarkMessage(msg, "")
	}

	return nil
}
```

That's it for the code, we can build it by using Docker compose <br/>
Let's define it as a new service in our `docker-compose.yaml` file

```
kafka-consumer-go:
  container_name: kafka-consumer-go
  image: aimvector/kafka-consumer-go:1.0.0
  build: 
    context: ./applications/consumer
  environment:
  - "KAFKA_PEERS=kafka-1:9092,kafka-2:9092,kafka-3:9092"
  - "KAFKA_TOPIC=Orders"
  - "KAFKA_VERSION=2.7.0"
  - "KAFKA_GROUP=orders"
  networks:
  - kafka

```

To build our consumer, we'll need a `dockerfile`:

```
FROM golang:1.16-alpine as dev-env

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

```

Let's build our container so we can compile our `go` code:

```
cd .\messaging\kafka\applications\consumer\
docker build . -t consumer
docker run -it -v ${PWD}:/app consumer sh
ls -l

go mod init consumer
go mod tidy

go build
```

Now that we have our `go.mod` and `go.sum` files, and code is comiling, we can proceed to create a container that runs our app. Let's extend the `dockerfile`

```
FROM dev-env as build-env
COPY go.mod /go.sum /app/
RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 go build -o /consumer

FROM alpine:3.10 as runtime

COPY --from=build-env /consumer /usr/local/bin/consumer
RUN chmod +x /usr/local/bin/consumer

ENTRYPOINT ["consumer"]
```

Now we can proceed to build it:

```
cd .\messaging\kafka\
docker compose build kafka-consumer-go
```

## Starting our Kafka Consumer

Now with our Kafka environment and topic ready, our consumer can start and subscribe to the orders topic:

```
cd messaging/kafka
docker compose up kafka-consumer-go

```

## Consume Messages in Random Order

Let's produce messages in a loop

```
cd .\messaging\kafka\
docker compose up -d kafka-producer
docker exec -it kafka-producer bash
```

Produce messages in a loop:

```
upperlim=10
for ((i=0; i<=upperlim; i++)); do
   echo "{ 'id' : 'order-$i', 'data' : 'random-data'}" | \
    /kafka/bin/kafka-console-producer.sh \
    --broker-list kafka-1:9092,kafka-2:9092,kafka-3:9092 \
    --topic Orders > /dev/null
done
```

<p>
Notice messages go into partition in round robin.
Partition provides distribution of messages so all our messages
don't take up all the space on a single broker and distributed 
more evenly
</p>

<p>
Notice when we stop and restart out consumer, it does not 
re-consume message from the beginning.
</p>

<p>
This is because the brokers remember the offset where the consumer 
has read up to in our case as we use consumer groups.
</p>

Reading messages will occur in random order because they exist in different partitions:

```
cd .\messaging\kafka\
docker compose up -d kafka-consumer
docker exec -it kafka-consumer bash

/kafka/bin/kafka-console-consumer.sh \
--bootstrap-server kafka-1:9092,kafka-2:9092,kafka-3:9092 \
--topic Orders --from-beginning


```

Notice we receive messages out of order and also able to read
messages from the beginning too.

## Consume Messages in Order

To consume messages in given order, we need to use message keys.
When using keys, Kafka will place all messages with a given key in the same
partition.
This is useful for event based transactional systems where order is important.

Let's say we have `order-11` and we want to update it 10 times.
Order here is important so we'l structure our message with a key.

```
upperlim=10
for ((i=0; i<=upperlim; i++)); do
   echo "order-11: { 'id' : 'order-11', 'data' : '$i'}" | \
    /kafka/bin/kafka-console-producer.sh \
    --broker-list kafka-1:9092,kafka-2:9092,kafka-3:9092 \
    --topic Orders > /dev/null \
    --property "parse.key=true" --property "key.separator=:"
done

```