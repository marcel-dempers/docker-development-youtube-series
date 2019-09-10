package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter" 
	"github.com/valyala/fasthttp"
)

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello") 
}

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "ok") 
}
func main() {
    
	fmt.Println("starting...")

	router := fasthttprouter.New()
	router.GET("/", Response)
	router.GET("/status", Status)
	
	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}