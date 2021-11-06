package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	tmplWelcome := template.Must(template.ParseFiles("base.html", "welcome.html"))
	tmplWoop := template.Must(template.ParseFiles("base.html", "button.html", "woop.html", "alert.html"))
	tmplBtn := template.Must(template.ParseFiles("button.html"))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmplWelcome.Execute(rw, nil)
	})
	http.HandleFunc("/woop", func(rw http.ResponseWriter, r *http.Request) {
		log.Print(tmplWoop.Execute(rw, "knopf"))
	})
	http.HandleFunc("/btn", func(rw http.ResponseWriter, r *http.Request) {
		tmplBtn.Execute(rw, "knopf")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
