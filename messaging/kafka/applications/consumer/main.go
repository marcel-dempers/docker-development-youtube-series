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

type Consumer struct {
	ready chan bool
}

func main() {

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

	consumer := Consumer{
		ready: make(chan bool),
	}

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

}

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