package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"fmt"
	"context"
	"os"
	"math/rand"
)

var environment = os.Getenv("ENVIRONMENT")
var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")
var flaky = os.Getenv("FLAKY")

var ctx = context.Background()
var rdb *redis.Client

func main() {

	router := httprouter.New()

	router.GET("/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		
		if flaky == "true"{
			if rand.Intn(90) < 30 {
				panic("flaky error occurred ")
		  } 
		}
		
		video := video(w,r,p)

		cors(w)
		fmt.Fprintf(w, "%s", video)
	})

	r := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		DB:       0,
	})
	rdb = r

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

func video(writer http.ResponseWriter, request *http.Request, p httprouter.Params)(response string){
	
	id := p.ByName("id")
	fmt.Print(id)

	videoData, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return "{}"
	} else if err != nil {
		panic(err)
} else {
	return videoData
}
}

type stop struct {
	error
}

func cors(writer http.ResponseWriter) () {
	if(environment == "DEBUG"){
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

type videos struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageurl string `json:"imageurl"`
	Url string `json:"url"`

}
