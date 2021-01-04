package main

import "fmt"

func main() {
	
	videos := getVideos()
	
	for i, _ := range videos {
	  videos[i].Description = videos[i].Description + "\n Remember to Like & Subscribe!"
  }

	fmt.Println(videos)

	saveVideos(videos)

}