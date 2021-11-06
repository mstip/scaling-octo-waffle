package main

import (
	"exchanger/pkg"
	"fmt"
	"log"
)

func main() {
	ex, err := pkg.GetExchanges()

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ex {
		fmt.Printf("%s %f\n", v.Currency, v.Rate)
	}
}
