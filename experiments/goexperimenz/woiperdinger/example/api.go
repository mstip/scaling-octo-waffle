package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Todo struct {
	gorm.Model
	Name string
	Done bool
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

type WebServer struct {
	Router  *mux.Router
	Storage *gorm.DB
}

func NewWebServer(storage *gorm.DB) (*WebServer, error) {
	w := &WebServer{}
	w.routes()
	w.Storage = storage
	return w, nil
}

func (ws *WebServer) routes() {
	ws.Router = mux.NewRouter()
	ws.Router.Handle("/todos", ws.handleTodosIndex()).Methods("GET")
	ws.Router.Handle("/todos", ws.handleTodosStore()).Methods("POST")
	ws.Router.Handle("/todos/{ID}", ws.handleTodosShow()).Methods("GET")
	ws.Router.Handle("/todos/{ID}", ws.handleTodosUpdate()).Methods("PUT")
	ws.Router.Handle("/todos/{ID}", ws.handleTodosDelete()).Methods("DELETE")
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.Router.ServeHTTP(w, r)
}

func (ws *WebServer) handleTodosIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []Todo
		ws.Storage.Find(&todos)
		respondWithJSON(w, http.StatusOK, todos)
	}
}

func (ws *WebServer) handleTodosStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&todo); err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, "Invalid request payload")
			return
		}
		defer func() {
			err := r.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		result := ws.Storage.Create(&todo) // pass pointer of data to Create

		if result.Error != nil {
			log.Println(result.Error)
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		respondWithJSON(w, http.StatusOK, todo)
	}
}

func (ws *WebServer) handleTodosShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["ID"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		var todo Todo
		result := ws.Storage.First(&todo, ID)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Not found")
			return
		}
		respondWithJSON(w, http.StatusOK, todo)
	}
}

func (ws *WebServer) handleTodosUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["ID"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		var todo Todo
		result := ws.Storage.First(&todo, ID)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Not found")
			return
		}

		var newTodo Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&newTodo); err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, "Invalid request payload")
			return
		}

		todo.Name = newTodo.Name
		todo.Done = newTodo.Done

		ws.Storage.Save(todo)

		respondWithJSON(w, http.StatusOK, todo)
	}
}

func (ws *WebServer) handleTodosDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["ID"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		var todo Todo
		result := ws.Storage.First(&todo, ID)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Not found")
			return
		}
		ws.Storage.Delete(todo)
		respondWithJSON(w, http.StatusOK, nil)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal(err)
	}

	webServer, err := NewWebServer(db)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3000"
	}

	log.Println("listen to :" + serverPort)
	err = http.ListenAndServe(":"+serverPort, webServer.Router)
	if err != nil {
		log.Fatal(err)
	}
}
