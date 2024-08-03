package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	godotenv.Load(".env")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelId := os.Getenv("CHANNEL_ID")
	fileSlice := []string{"pic.jpg"}

	for i := 0; i < len(fileSlice); i++ {
		params := slack.UploadFileV2Parameters{
			Channel: channelId,
			File: fileSlice[i],
			Filename: "pic",
			FileSize: 10000,
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %v\n", file)
	}
}