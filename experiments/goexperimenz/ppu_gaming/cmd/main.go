package main

import (
	"github.com/joho/godotenv"
	"github.com/ungefaehrlich/ppu_gaming/pkg/web"
	"log"
	"net/http"
	"os"
)

func main() {
	// this is magic and populate all env calls by magic :)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := web.NewServer(
		os.Getenv("PATH_TO_SSH_SECRET"),
		os.Getenv("HCLOUD_TOKEN"),
		os.Getenv("STORE_SERVER_IP"),
		os.Getenv("STORE_SERVER_USER"),
	)
	defer s.CleanUp()
	log.Println("ppu_gaming running on 0.0.0.0:" + os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe("0.0.0.0:"+os.Getenv("SERVER_PORT"), s); err != nil {
		log.Fatal("listen and serve failed", err)
	}
}
