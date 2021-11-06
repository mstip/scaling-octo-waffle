package pkg

import (
	"html/template"
	"log"
	"ndis/pkg/vault"
	"ndis/pkg/web"
	"net/http"
	"strings"
)

type Game struct {
	views *template.Template
	vault *vault.Vault
}

func NewGame() *Game {

	return &Game{
		views: template.Must(template.ParseFiles("tmpl/index.html", "tmpl/not_found.html", "tmpl/game.html")),
		vault: vault.NewVault(),
	}
}

func (g *Game) Run() error {
	log.Fatal(http.ListenAndServe(":3000", g))
	return nil
}

func (g *Game) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/static") {
		http.StripPrefix("/static", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
		return
	}

	switch r.URL.Path {
	case "/":
		if err := g.views.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Fatal(err)
		}
		break
	case "/login":
		if r.Method != "POST" {
			if err := g.notFound(w, r); err != nil {
				log.Fatal(err)
			}
		}
		if err := g.login(w, r); err != nil {
			log.Fatal(err)
		}
		break
	case "/game":
		userId, _ := g.GetUserIdFromSession(r)

		if userId == 0 {
			log.Println("userid in session is 0")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

		userName := g.vault.GetUserById(userId)

		if err := g.views.ExecuteTemplate(w, "game.html", struct {
			User string
		}{User: userName}); err != nil {
			log.Fatal(err)
		}
		break
	default:
		if err := g.notFound(w, r); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) notFound(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s - %s - 404 - not found", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
	if err := g.views.ExecuteTemplate(w, "not_found.html", nil); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (g *Game) login(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	user := r.Form.Get("user")
	password := r.Form.Get("password")

	userId, err := g.vault.CheckLogin(user, password)
	if err != nil {
		return err
	}

	if userId != 0 {
		sessId, _ := g.vault.StartSessionForUser(userId)
		web.SetSessionCookie(sessId, w)
		http.Redirect(w, r, "/game", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	return nil
}

func (g *Game) GetUserIdFromSession(r *http.Request) (int, error) {
	sessionId, err := web.GetSessionId(r)
	if err != nil {
		return 0, err
	}
	return g.vault.GetSessionUserId(sessionId)
}

func RunGame() error {
	game := NewGame()
	err := game.Run()
	if err != nil {
		return err
	}
	return nil
}
