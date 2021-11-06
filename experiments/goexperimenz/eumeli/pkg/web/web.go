package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	router *mux.Router
}

func NewWebServer() (*WebServer, error) {
	w := &WebServer{}
	w.routes()
	return w, nil
}

func (ws *WebServer) routes() {
	ws.router = mux.NewRouter()
	ws.router.Handle("/", ws.handleIndex()).Methods("GET")
}

func (ws *WebServer) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "worx"); err != nil {
			log.Println("failed to write response", err)
		}
	}
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.router.ServeHTTP(w, r)
}
