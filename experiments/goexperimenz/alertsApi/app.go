package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

type App struct {
	Router  *mux.Router
	DB      *sql.DB
	ApiKeys []string
}

func (a *App) Initialize(user, password, dbname, dbhost string) {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dbhost, dbname)

	log.Println("DB Connection: " + connectionString)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := a.DB.PingContext(ctx); err != nil {
		log.Fatal("ping", err)
		return
	}

	a.Router = mux.NewRouter()
	a.Router.Use(a.checkAuth)

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/last-contact", a.getLastContact).Methods("GET")
	a.Router.HandleFunc("/low-disk", a.getLowDisk).Methods("GET")
	a.Router.HandleFunc("/low-memory", a.getLowMemory).Methods("GET")
	a.Router.HandleFunc("/low-swap", a.getLowSwap).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Println("Start Server at " + addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a App) checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authKey, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal", http.StatusInternalServerError)
		}

		for _, key := range a.ApiKeys {
			if key == string(authKey) {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (a *App) getLastContact(w http.ResponseWriter, r *http.Request) {
	systems, err := querySystemsWhereLastContactWasMinutesAgo(a.DB, 15)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, systems)
}

func (a *App) getLowDisk(w http.ResponseWriter, r *http.Request) {
	systems, err := querySystemsWithLowDisk(a.DB, 5)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, systems)
}

func (a *App) getLowMemory(w http.ResponseWriter, r *http.Request) {
	systems, err := querySystemsWithLowMemory(a.DB, 1)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, systems)
}

func (a *App) getLowSwap(w http.ResponseWriter, r *http.Request) {
	systems, err := querySystemsWithLowSwap(a.DB, 1)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, systems)
}
