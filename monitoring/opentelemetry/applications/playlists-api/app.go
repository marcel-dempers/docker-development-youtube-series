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
			 
				v := videos{}
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

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

func getPlaylists()(response string){
	playlistData, err := rdb.Get(ctx, "playlists").Result()
	
	if err != nil {
		fmt.Println(err)
		fmt.Println("error occured retrieving playlists from Redis")
		return "[]"
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

func cors(writer http.ResponseWriter) () {
	if(environment == "DEBUG"){
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}