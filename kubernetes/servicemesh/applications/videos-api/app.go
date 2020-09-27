package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"fmt"
	"time"
	"context"
	"encoding/json"
	"os"
	"strconv"
)

var environment = os.Getenv("ENVIRONMENT")
var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")

var ctx = context.Background()
var rdb *redis.Client

func main() {

	router := httprouter.New()

	router.GET("/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		video := video(w,r,p)

		cors(w)
		fmt.Fprintf(w, "%s", video)
	})

	r := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		DB:       0,
	})
	rdb = r

	initDummyData()

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

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}
 
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

func cors(writer http.ResponseWriter) () {
	if(environment == "DEBUG"){
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

func initDummyData(){

	//init dummy data into redis
	duration, _ := time.ParseDuration("4s")
 
	v1 := `{
		"id" : "jM36M39MA3I",
		"title" : "Kubernetes cluster autoscaling",
		"imageurl" : "https://i.ytimg.com/vi/jM36M39MA3I/sddefault.jpg",
		"url" : "https://youtu.be/jM36M39MA3I",
		"description" : ""
	}
	`
	v2 := `{
		"id" : "FfDI08sgrYY",
		"title" : "Kubernetes pod autoscaling",
		"imageurl" : "https://i.ytimg.com/vi/FfDI08sgrYY/sddefault.jpg",
		"url" : "https://youtu.be/FfDI08sgrYY",
		"description" : ""
	}
	`
	v3 := `{
		"id" : "JmCn7k0PlV4",
		"title" : "Redis on Kubernetes",
		"imageurl" : "https://i.ytimg.com/vi/JmCn7k0PlV4/sddefault.jpg",
		"url" : "https://youtu.be/JmCn7k0PlV4",
		"description" : ""
	}
	`
	v4 := `{
		"id" : "_lpDfMkxccc",
		"title" : "RabbitMQ on Kubernetes",
		"imageurl" : "https://i.ytimg.com/vi/_lpDfMkxccc/sddefault.jpg",
		"url" : "https://youtu.be/_lpDfMkxccc",
		"description" : ""
	}
	`
	v5 := `{
		"id" : "OFgziggbCOg",
		"title" : "Flux CD",
		"imageurl" : "https://i.ytimg.com/vi/OFgziggbCOg/sddefault.jpg",
		"url" : "https://youtu.be/OFgziggbCOg",
		"description" : ""
	}
	`
	v6 := `{
		"id" : "myCcJJ_Fk10",
		"title" : "Drone CI",
		"imageurl" : "https://i.ytimg.com/vi/myCcJJ_Fk10/sddefault.jpg",
		"url" : "https://youtu.be/myCcJJ_Fk10",
		"description" : ""
	}
	`
	v7 := `{
		"id" : "2WSJF7d8dUg",
		"title" : "Argo CD",
		"imageurl" : "https://i.ytimg.com/vi/2WSJF7d8dUg/sddefault.jpg",
		"url" : "https://youtu.be/2WSJF7d8dUg",
		"description" : ""
	}
	`
	v8 := `{
		"id" : "QThadS3Soig",
		"title" : "Kubernetes on Amazon",
		"imageurl" : "https://i.ytimg.com/vi/QThadS3Soig/sddefault.jpg",
		"url" : "https://youtu.be/QThadS3Soig",
		"description" : ""
	}
	`
	v9 := `{
		"id" : "eyvLwK5C2dw",
		"title" : "Kubernetes on Azure",
		"imageurl" : "https://i.ytimg.com/vi/eyvLwK5C2dw/mqdefault.jpg?sqp=CISC_PoF&rs=AOn4CLDo7kizrJozB0pxBhxL9JbyiW_EPw",
		"url" : "https://youtu.be/eyvLwK5C2dw",
		"description" : ""
	}`


	dummy := []string{ v1, v2, v3, v4,v5, v6, v7, v8,v9 }
  

	retry(10, duration, func() (err error){
	
		for i := range dummy {
			dummyItem := dummy[i]
			v := videos{}
			iStr := strconv.Itoa(i)

			fmt.Println("checking vid: " +  iStr)
			err = json.Unmarshal([]byte(dummyItem), &v)
			if err != nil {
				break
			}

			fmt.Println("adding vid: " +  v.Id)
			err = rdb.Set(ctx, v.Id,dummyItem , 0).Err()

			if err != nil {
				break
			}
		}
		

		if err != nil {
			fmt.Println("error occured connecting to Redis, retrying...")
			return err
		}

	

		fmt.Println("Redis dummy data initialised")
		return nil
	})

}

type videos struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageurl string `json:"imageurl"`
	Url string `json:"url"`

}
