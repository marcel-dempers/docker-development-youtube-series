package main

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type video struct {
	Id string  
	Title string  
	Description string 
	Imageurl string 
	Url string 
}

func getVideos()(videos []video){
	
	keys, err := redisClient.Keys(ctx,"*").Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		video := getVideo(key)
		videos = append(videos, video)
	}
	return videos
}

func getVideo(id string)(video video) {
	
	value, err := redisClient.Get(ctx, id).Result()

	if err != nil {
		panic(err)
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &video)
	}
	
	return video
}

func saveVideo(video video)(){

  videoBytes, err  := json.Marshal(video)
  if err != nil {
		panic(err)
	}

	err = redisClient.Set(ctx, video.Id, videoBytes, 0).Err()
  if err != nil {
		panic(err)
	}

}

func saveVideos(videos []video)(){
	for _, video := range videos {
		saveVideo(video)
	}
}