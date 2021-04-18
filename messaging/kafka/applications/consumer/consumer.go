package main


import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

var kafka_host = os.Getenv("KAFKA_HOSTS")

func main() {
	consume()
}

func consume() {

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to Kafka", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			
			d.Ack(false)
		}
	  }()
	  
	  fmt.Println("Running...")
	  <-forever
}
