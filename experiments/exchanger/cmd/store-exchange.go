package main

import (
	"exchanger/pkg"
	"log"
)

func main() {
	db, err := pkg.NewDb()

	if err != nil {
		log.Fatal(err)
	}
	ex, err := pkg.GetExchanges()

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ex {
		db.Create(&pkg.ExchangeRates{
			Currency: v.Currency,
			Rate:     v.Rate,
		})
	}
}
