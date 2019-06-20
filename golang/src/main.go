package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter" 
	"github.com/valyala/fasthttp"
)

var searchMock []byte

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello") 
}
func main() {
    
	fmt.Println("starting.")

	router := fasthttprouter.New()
	router.GET("/", Response)
	
	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}