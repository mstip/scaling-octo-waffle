package main

import (
	"eumeli/pkg/web"
	"log"
	"net/http"
)

func main() {
	s, _ := web.NewWebServer()
	if err := http.ListenAndServe("0.0.0.0:3000", s); err != nil {
		log.Fatal("listen and serve failed", err)
	}
}
