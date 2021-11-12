package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Text string
	Done bool
}

var todos []Todo
var tmpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	if tmpl == nil {
		tmpl = template.Must(template.ParseFiles("todo.html"))
	}
	tmpl.Execute(w, struct{ Todos []Todo }{Todos: todos})
}

func create(w http.ResponseWriter, r *http.Request) {
	todos = append(todos, Todo{Text: r.FormValue("text")})
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["index"])
	if r.FormValue("done") == "on" {
		todos[index].Done = true
	} else {
		todos[index].Done = false
	}
	todos[index].Text = r.FormValue("text")

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["index"])
	todos = append(todos[:index], todos[index+1:]...)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	tmpl = template.Must(template.ParseFiles("todo.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/", create).Methods(http.MethodPost)
	r.HandleFunc("/update/{index}", update).Methods(http.MethodPost)
	r.HandleFunc("/delete/{index}", delete).Methods(http.MethodPost)

	http.ListenAndServe(":3000", r)
}
