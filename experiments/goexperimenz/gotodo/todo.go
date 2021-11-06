package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Todo struct {
	Text string
	Done bool
}

var todos []Todo

var tmpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		todos = append(todos, Todo{Text: r.FormValue("text"), Done: false})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.Execute(
		w,
		struct{ Todos []Todo }{
			Todos: todos,
		},
	)
}

func toggleTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	rawTodoIndex := r.FormValue("todoIndex")
	todoIndex, err := strconv.Atoi(rawTodoIndex)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	todos[todoIndex].Done = !todos[todoIndex].Done
	fmt.Fprintf(w, "OK")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	rawTodoIndex := r.FormValue("todoIndex")
	todoIndex, err := strconv.Atoi(rawTodoIndex)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	todos = append(todos[:todoIndex], todos[todoIndex+1:]...)
	fmt.Fprintf(w, "OK")
}

func main() {
	tmpl = template.Must(template.ParseFiles("todo.html"))
	todos = append(todos, Todo{Text: "test", Done: false})
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/toggleTodo", toggleTodo)
	http.HandleFunc("/deleteTodo", deleteTodo)
	fmt.Println("listen on :3000")
	http.ListenAndServe(":3000", nil)
}
