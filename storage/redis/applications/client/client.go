package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/go-redis/redis/v8"
	"context"
	"strconv"
)

var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT") 
var redis_password = os.Getenv("REDIS_PASSWORD")

var ctx = context.Background()
var rdb *redis.Client

var counter = 0
func main() {

	r := redis.NewClient(&redis.Options{
        Addr:     redis_host + ":" + redis_port,
        Password: redis_password, // no password set
        DB:       0,  // use default DB
	})
	rdb = r

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		increment_redis_key(w,r,p)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func increment_redis_key(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	
    val, err := rdb.Get(ctx, "counter").Result()
	
	if err == redis.Nil {
		err := rdb.Set(ctx, "counter", 1, 0).Err()
		counter++
		if err != nil {
			panic(err)
		}

    } else if err != nil {
        panic(err)
    } else {
		counter,_ = strconv.Atoi(val)
		counter++
		err := rdb.Set(ctx, "counter", counter, 0).Err()

		if err != nil {
			panic(err)
		}
	}
	
	fmt.Fprint(writer, counter)
	fmt.Println("counter", counter)
}