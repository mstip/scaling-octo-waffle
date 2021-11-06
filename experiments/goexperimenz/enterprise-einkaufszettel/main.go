package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ShopItem struct {
	Name string
	Done bool
}

var items []ShopItem

var tmpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, struct {
		Items []ShopItem
	}{Items: items})
}

func newItem(w http.ResponseWriter, r *http.Request) {
	newItem := r.FormValue("newItem")
	items = append(items, ShopItem{Name: newItem})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func toggleItemDone(w http.ResponseWriter, r *http.Request) {
	rawItemIndex := r.FormValue("itemIndex")
	itemIndex, err := strconv.Atoi(rawItemIndex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	items[itemIndex].Done = !items[itemIndex].Done
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	rawItemIndex := r.FormValue("itemIndex")
	itemIndex, err := strconv.Atoi(rawItemIndex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	items = append(items[:itemIndex], items[itemIndex+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func main() {

	items = append(items, ShopItem{Name: "Käse"})
	items = append(items, ShopItem{Name: "Wurst", Done: true})
	items = append(items, ShopItem{Name: "Klopapier"})
	items = append(items, ShopItem{Name: "Keckse"})

	tmpl = template.Must(template.ParseFiles("einkaufszettel.html"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", index)
	r.Post("/new-item", newItem)
	r.Post("/toggle-item-done", toggleItemDone)
	r.Post("/delete-item", deleteItem)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)

	http.ListenAndServe(":3000", r)
}
