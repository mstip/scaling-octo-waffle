package epg

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"regexp"
	"time"
)

type WebServer struct {
	Router *chi.Mux
	DB     *sqlx.DB
}

func NewWebServer() (*WebServer, error) {
	var err error
	w := WebServer{}

	// database
	w.DB, err = sqlx.Connect("postgres", "user=postgres dbname=dev password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	// router
	w.Router = chi.NewRouter()
	// middlewares
	w.Router.Use(middleware.RequestID)
	w.Router.Use(middleware.Logger)
	w.Router.Use(middleware.Recoverer)
	w.Router.Use(middleware.URLFormat)
	w.Router.Use(middleware.Timeout(60 * time.Second))
	w.Router.Use(w.AuthMiddleware)

	// routes
	w.Router.Route("/accounts", func(r chi.Router) {
		r.Get("/", w.AccountsIndexHandler)
		r.Post("/", w.AccountsStoreHandler)
	})

	return &w, nil
}

func (webServer *WebServer) CleanUp() error {
	err := webServer.DB.Close()
	return err
}

func (webServer *WebServer) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey != "secret" {
			respondWithError(w, 403, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (webServer *WebServer) AccountsIndexHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := QueryAccounts(webServer.DB)
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	respondWithJSON(w, 200, accounts)
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (webServer *WebServer) AccountsStoreHandler(w http.ResponseWriter, r *http.Request) {
	var account Account
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if len(account.Email) < 3 && len(account.Email) > 254 {
		respondWithError(w, http.StatusUnprocessableEntity, "Invalid email")
		return
	}
	if valid := emailRegex.MatchString(account.Email); !valid {
		respondWithError(w, http.StatusUnprocessableEntity, "Invalid email")
		return
	}

	newAccount, err := CreateAccount(webServer.DB, &account)
	if err != nil {
		respondWithError(w, 500, err.Error())
	}
	respondWithJSON(w, 200, newAccount)
}
