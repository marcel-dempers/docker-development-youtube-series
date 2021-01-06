package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func main() {
	
	http.HandleFunc("/", HandleGetVideos)
	http.HandleFunc("/update", HandleUpdateVideos)

	http.ListenAndServe(":8080", nil)
}


func HandleGetVideos(w http.ResponseWriter, r *http.Request){
	
	videos := getVideos()

	videoBytes, err  := json.Marshal(videos)

	if err != nil {
  	panic(err)
	}

	w.Write(videoBytes)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
			
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var videos []video
			err = json.Unmarshal(body, &videos)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Bad request")
			}

			saveVideos(videos)

		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method not Supported!")
		}
}