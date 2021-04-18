package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"os"
)

var kafka_host = os.Getenv("KAFKA_HOSTS")

func main() {

	router := httprouter.New()

	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		submit(w,r,p)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func submit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	message := p.ByName("message")
	
	fmt.Println("Received message: " + message)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to Kafka", err)
	}

	defer conn.Close()
	fmt.Println("publish success!")
}