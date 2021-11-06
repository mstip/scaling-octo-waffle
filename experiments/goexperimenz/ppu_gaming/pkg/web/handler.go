package web

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ungefaehrlich/ppu_gaming/pkg/servercontrol"
)

func (s *server) handleIndex() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(s.tmplDir+"layout.html", s.tmplDir+"index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (s *server) handleShowRegister() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(s.tmplDir+"layout.html", s.tmplDir+"register.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (s *server) handleRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.FormValue("email"))
		log.Print(r.FormValue("name"))
		log.Print(r.FormValue("password"))
		http.Redirect(w, r, "/login", 301)
	}
}

func (s *server) handleOverview() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(s.tmplDir+"layout.html", s.tmplDir+"overview.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		servers, err := s.sc.GetAllServers()
		if err != nil {
			s.woops(w, err)
			return
		}

		customerID := "1"
		saves, err := s.listSavesByCustomerID(customerID)
		if err != nil {
			s.woops(w, err)
			return
		}

		flashes, err := s.getFlashes(w, r)
		if err != nil {
			s.woops(w, err)
			return
		}

		tmpl.Execute(w, struct {
			Servers []servercontrol.ServerInfo
			Flashes []string
			Saves   []string
		}{Servers: servers, Flashes: flashes, Saves: saves})
	}
}

func (s *server) createNewServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serverType, err := strconv.Atoi(r.FormValue("serverType"))
		if err != nil {
			s.woops(w, err)
			return
		}

		if serverType < 1 || serverType > 9 {
			s.validatorError(w, "invalid server type")
			return
		}

		go s.createServerAction(serverType)

		if err = s.addFlash(w, r, "New server will be created, this can take some minutes"); err != nil {
			s.woops(w, err)
			return
		}

		http.Redirect(w, r, "/", 301)
	}
}

func (s *server) saveAndDeleteServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		go s.saveAndDeleteServerAction(params["ID"])

		if err := s.addFlash(w, r, "server will be saved and deleted, this can take some minutes"); err != nil {
			s.woops(w, err)
			return
		}

		http.Redirect(w, r, "/", 301)
	}
}

func (s *server) deleteServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		go s.deleteServerAction(params["ID"])

		if err := s.addFlash(w, r, "Server will be deleted"); err != nil {
			s.woops(w, err)
			return
		}

		http.Redirect(w, r, "/", 301)
	}
}

func (s *server) loadSaveInNewServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		serverType, err := strconv.Atoi(r.FormValue("serverType"))
		if err != nil {
			s.woops(w, err)
			return
		}

		if serverType < 1 || serverType > 9 {
			s.validatorError(w, "invalid server type")
			return
		}

		go s.loadSaveInNewServerAction(serverType, params["saveFileName"])

		if err = s.addFlash(w, r, "save game will be loaded in a new server, this can take some minutes"); err != nil {
			s.woops(w, err)
			return
		}

		http.Redirect(w, r, "/", 301)
	}
}

func (s *server) deleteSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		go s.deleteSaveAction(params["saveFileName"])

		if err := s.addFlash(w, r, "save will be deleted, this can take some minutes"); err != nil {
			s.woops(w, err)
			return
		}

		http.Redirect(w, r, "/", 301)
	}
}
