package main

import (
	epg "epg/pkg"
	"log"
	"net/http"
)

func main() {
	webServer, err := epg.NewWebServer()
	defer func() {
		err := webServer.CleanUp()
		if err != nil {
			log.Fatal()
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen to :3000")
	err = http.ListenAndServe(":3000", webServer.Router)
	if err != nil {
		log.Fatal(err)
	}
}
