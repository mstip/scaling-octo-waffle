package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"platform/pkg/pushover"
)

func main() {
	// this is magic and populate all env calls by magic :)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var text string

	text += "wusar\n"
	text += "moep \n"
	text += "waarrr"

	p := pushover.NewPushover(os.Getenv("PUSHOVER_TOKEN"), os.Getenv("PUSHOVER_USER"))
	if err := p.SendMessage(text); err != nil {
		log.Fatal(err)
	}
}
