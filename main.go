package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	// os.Setenv("CHANNEL_ID", "")
	// os.Setenv("SLACK_BOT_TOKEN", "")
	err := godotenv.Load(".env")
	checkError(err)

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Scorecard.sql"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		checkError(err)
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
