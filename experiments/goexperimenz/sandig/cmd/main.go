package main

import (
	"log"
	"sandig/pkg"
)

func main() {
	g, err := pkg.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	g.Run()
}
