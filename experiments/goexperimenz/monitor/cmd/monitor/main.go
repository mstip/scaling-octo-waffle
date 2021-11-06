package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"platform/pkg/monitor"
)

func main() {
	// this is magic and populate all env calls by magic :)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m, err := monitor.NewMonitor()
	if err != nil {
		log.Fatal("failed to create monitor", err)
	}

	err = m.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
