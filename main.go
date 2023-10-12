package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	err := godotenv.Load("../env/file.env")
	if err != nil {
		log.Fatalf("Error occured : %s", err)
	}

	slack_bot_token := os.Getenv("SLACK_BOT_TOKEN")
	channel_id := os.Getenv("CHANNEL_ID")

	api := slack.New(slack_bot_token)
	channelArr := []string{channel_id}

	fileArr := []string{"../myfile.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Printf("File Name: %s\n", file.Name)
	}

}
