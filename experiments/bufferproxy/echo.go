package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("echo up")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		fmt.Println(string(body))
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	})
	log.Fatal(http.ListenAndServe(":31337", nil))
}
