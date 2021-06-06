package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"strconv"
	"gopkg.in/Shopify/sarama.v1"
)

var kafkaBrokers = os.Getenv("KAFKA_PEERS")
var kafkaTopic = os.Getenv("KAFKA_TOPIC")
var kafkaPartition = os.Getenv("KAFKA_PARTITION")
var partition int32 = -1
var globalProducer sarama.SyncProducer 

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
  p, err := strconv.Atoi(kafkaPartition)

	if err != nil {
	  fmt.Println("Failed to convert KAFKA_PARTITION to Int32")
		panic(err)
	}

	partition = int32(p)

	producer, err := sarama.NewSyncProducer(strings.Split(kafkaBrokers, ","), config)
	if err != nil {
	  fmt.Printf("Failed to open Kafka producer: %s", err)
		panic(err)
	}

	globalProducer = producer

	defer func() {
	  fmt.Println("Closing Kafka producer...")
		if err := globalProducer.Close(); err != nil {
		  fmt.Printf("Failed to close Kafka producer cleanly: %s", err)
		  panic(err)
		}
	}()

	router := httprouter.New()

	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		submit(w,r,p)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func submit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	
	messageValue := p.ByName("message")
	message := &sarama.ProducerMessage{Topic: kafkaTopic, Partition: partition }
	message.Value = sarama.StringEncoder(messageValue)

	fmt.Println("Received message: " + messageValue)

	partition, offset, err := globalProducer.SendMessage(message)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to Kafka", err)
	}

	fmt.Printf("publish success! topic=%s\tpartition=%d\toffset=%d\n", kafkaTopic, partition, offset)

}
