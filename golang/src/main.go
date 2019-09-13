package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter" 
	"github.com/valyala/fasthttp"
	"os"
	"io/ioutil"
)

var configuration []byte
var secret []byte

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello") 
}

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "ok") 
}

func ReadConfig(){
	fmt.Println("reading config...")
	config, e := ioutil.ReadFile("/configs/config.json")
	if e != nil {
		fmt.Printf("Error reading config file: %v\n", e)
		os.Exit(1)
	}
	configuration = config
	fmt.Println("config loaded!")

}

func ReadSecret(){
	fmt.Println("reading secret...")
	s, e := ioutil.ReadFile("/secrets/secret.json")
	if e != nil {
		fmt.Printf("Error reading secret file: %v\n", e)
		os.Exit(1)
	}
	secret = s
	fmt.Println("secret loaded!")

}

func main() {
    
	fmt.Println("starting...")
	ReadConfig()
	ReadSecret()
	router := fasthttprouter.New()
	router.GET("/", Response)
	router.GET("/status", Status)
	
	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}
