package main

import (
	"encoding/json"
	"exchanger/pkg"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	db, err := pkg.NewDb()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var ex []pkg.ExchangeRates
		db.Find(&[]pkg.ExchangeRates{}).Scan(&ex)
		respondWithJSON(w, http.StatusOK, ex)
	})
	http.ListenAndServe(":3000", r)
}
