package main

import (
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

var globalProducer sarama.SyncProducer 

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	consumer, err := sarama.NewConsumer(strings.Split(kafkaBrokers, ","), config)
	if err != nil {
	  fmt.Printf("Failed to open Kafka consumer: %s", err)
		panic(err)
	}

	partitionList, err := consumer.Partitions(kafkaTopic)

	if err != nil {
		fmt.Printf("Failed to get the list of partitions: %s", err)
		panic(err)
	}

	var bufferSize = 256
	var (
		messages = make(chan *sarama.ConsumerMessage, bufferSize)
		closing  = make(chan struct{})
		wg       sync.WaitGroup
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
		<-signals
		fmt.Println("Initiating shutdown of consumer...")
		close(closing)
	}()
  
	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(kafkaTopic, partition, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			panic(err)
		}

		go func(pc sarama.PartitionConsumer) {
			<-closing
			pc.AsyncClose()
		}(pc)

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				messages <- message
			}
		}(pc)
	}

	go func() {
		for msg := range messages {
			fmt.Printf("Partition:\t%d\n", msg.Partition)
			fmt.Printf("Offset:\t%d\n", msg.Offset)
			fmt.Printf("Key:\t%s\n", string(msg.Key))
			fmt.Printf("Value:\t%s\n", string(msg.Value))
			fmt.Println()
		}
	}()

	wg.Wait()
	fmt.Println("Done consuming topic", kafkaTopic)
	close(messages)

	if err := consumer.Close(); err != nil {
		fmt.Printf("Failed to close consumer: %s", err)
		panic(err)
	}
}

