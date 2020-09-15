package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"fmt"
	"os"
	"bytes"
	"io/ioutil"
	"context"
	"time"
	"github.com/go-redis/redis/v8"
)

var environment = os.Getenv("ENVIRONMENT")
var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")

var ctx = context.Background()
var rdb *redis.Client

func main() {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		cors(w)
		playlistsJson := getPlaylists()
		
		playlists := []playlist{}
		err := json.Unmarshal([]byte(playlistsJson), &playlists)
		if err != nil {
			panic(err)
		}

		//get videos for each playlist from videos api
		for pi := range playlists {

			vs := []videos{}
			for vi := range playlists[pi].Videos {
			 
				videoResp, err := http.Get("http://videos-api:10010/" + playlists[pi].Videos[vi].Id)
				
				if err != nil {
					fmt.Println(err)
					break
				}

				defer videoResp.Body.Close()
				video, err := ioutil.ReadAll(videoResp.Body)

				if err != nil {
					panic(err)
				}

				v := videos{}
				err = json.Unmarshal(video, &v)

				if err != nil {
					panic(err)
				}
				
				vs = append(vs, v)

			}

			playlists[pi].Videos = vs
		}

		playlistsBytes, err := json.Marshal(playlists)
		if err != nil {
			panic(err)
		}

		reader := bytes.NewReader(playlistsBytes)
		if b, err := ioutil.ReadAll(reader); err == nil {
			fmt.Fprintf(w, "%s", string(b))
		}

	})

	r := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		DB:       0,
	})
	rdb = r

	//init dummy data into redis
	duration, _ := time.ParseDuration("4s")
	dummyData := `
	[
		{
			"id" : "1",
			"name" : "CI/CD",
			"videos": [ { "id" : "OFgziggbCOg"}, { "id" : "myCcJJ_Fk10"}, { "id" : "2WSJF7d8dUg"}]
		},
		{
			"id" : "2",
			"name" : "K8s in the Cloud",
			"videos": [ { "id" : "QThadS3Soig"}, { "id" : "eyvLwK5C2dw"}]
		},
		{
			"id" : "3",
			"name" : "Storage and MessageBrokers",
			"videos": [ { "id" : "JmCn7k0PlV4"}, { "id" : "_lpDfMkxccc"}]

		},
		{
			"id" : "4",
			"name" : "K8s Autoscaling",
			"videos": [ { "id" : "jM36M39MA3I"}, { "id" : "FfDI08sgrYY"}]
		}
	]
	`
	retry(10, duration, func() (err error){
	  err = rdb.Set(ctx, "playlists",dummyData , 0).Err()
		
		if err != nil {
			fmt.Println("error occured connecting to Redis, retrying...")
			return err
		}

		fmt.Println("Redis dummy data initialised")
		return nil
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

func getPlaylists()(response string){
	playlistData, err := rdb.Get(ctx, "playlists").Result()
	
	if err != nil {
		panic(err)
	}

	return playlistData
}

type playlist struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Videos []videos `json:"videos"`
}

type videos struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageurl string `json:"imageurl"`
	Url string `json:"url"`

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