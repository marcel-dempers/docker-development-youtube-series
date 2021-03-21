package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	
	//'videos get' subcommand
  getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `videos get` command
  getAll := getCmd.Bool("all", false, "Get all videos")
  getID := getCmd.String("id", "", "YouTube video ID")

	//'videos add' subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for `videos add` command
	addID := addCmd.String("id", "", "YouTube video ID")
  addTitle := addCmd.String("title", "", "YouTube video Title")
  addUrl := addCmd.String("url", "", "YouTube video URL")
  addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
  addDesc := addCmd.String("desc", "", "YouTube video description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
  }

	//look at the 2nd argument's value
	switch os.Args[1] {
		case "get": // if its the 'get' command
			HandleGet(getCmd, getAll, getID)
		case "add": // if its the 'add' command
			HandleAdd(addCmd, addID,addTitle,addUrl, addImageUrl, addDesc)
		default: // if we don't understand the input
	}


}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string){

	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
    fmt.Print("id is required or specify --all for all videos")
    getCmd.PrintDefaults()
    os.Exit(1)
  }

	if *all {
		//return all videos
		videos := getVideos()
		
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n",video.Id, video.Title, video.Url, video.Imageurl,video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n",video.Id, video.Title, video.Url, video.Imageurl,video.Description)
			}
		}
	} 



}

func ValidateVideo(addCmd *flag.FlagSet,id *string, title *string, url *string, imageUrl *string, description *string ){

	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet,id *string, title *string, url *string, imageUrl *string, description *string ){

	ValidateVideo(addCmd, id,title,url, imageUrl, description)

	video := video{
		Id: *id,
		Title: *title,
		Description: *description,
		Imageurl: *imageUrl, 
		Url: *url, 
	}

	videos := getVideos()
	videos = append(videos,video)

	saveVideos(videos)

}