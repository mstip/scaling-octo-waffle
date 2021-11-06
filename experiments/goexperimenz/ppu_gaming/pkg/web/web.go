package web

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/ungefaehrlich/ppu_gaming/pkg/rexec"
	"github.com/ungefaehrlich/ppu_gaming/pkg/servercontrol"
	"github.com/ungefaehrlich/ppu_gaming/pkg/storage"
)

type server struct {
	router        *mux.Router
	tmplDir       string
	gameServer    rexec.Rexecer
	storeServer   rexec.Rexecer
	sc            servercontrol.Servercontroler
	cookieStore   *sessions.CookieStore
	storeServerIP string
	storage       *storage.Storage
}

func (s *server) routes() {
	s.router = mux.NewRouter()
	s.router.Handle("/", s.handleIndex()).Methods("GET")
	s.router.Handle("/register", s.handleShowRegister()).Methods("GET")
	s.router.Handle("/overview", s.handleOverview()).Methods("GET")
	s.router.Handle("/create-new-server", s.createNewServer()).Methods("POST")
	s.router.Handle("/delete-server/{ID}", s.deleteServer()).Methods("POST")
	s.router.Handle("/save-and-delete-server/{ID}", s.saveAndDeleteServer()).Methods("POST")
	s.router.Handle("/load-save-in-new-server/{saveFileName}", s.loadSaveInNewServer()).Methods("POST")
	s.router.Handle("/delete-save/{saveFileName}", s.deleteSave()).Methods("POST")
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) getFlashes(w http.ResponseWriter, r *http.Request) ([]string, error) {
	session, err := s.cookieStore.Get(r, "ppu-keks")

	if err != nil {
		return nil, err
	}

	rawFlashes := session.Flashes()
	if err = session.Save(r, w); err != nil {
		return nil, err
	}

	var flashes []string
	for _, v := range rawFlashes {
		flashes = append(flashes, v.(string))
	}

	return flashes, nil
}

func (s *server) addFlash(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := s.cookieStore.Get(r, "ppu-keks")

	if err != nil {
		return err
	}

	session.AddFlash(message)
	if err = session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func (s *server) woops(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("wooops: " + err.Error()))
}

func (s *server) validatorError(w http.ResponseWriter, validatorMessage string) {
	log.Println("Validate Error: " + validatorMessage)
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte("Validator Error: " + validatorMessage))
}

func (s *server) CleanUp() error {
	err := s.storage.Close()
	return err
}

func NewServer(pathToKeyFile string, hcloudToken string, storeServerIP string, storeServerUser string) *server {
	s := &server{}
	s.tmplDir = "./web/template/"
	s.gameServer = rexec.NewRexec(pathToKeyFile, "root")
	if s.gameServer == nil {
		log.Fatal("could not create rexec for gameserver")
	}

	s.storeServer = rexec.NewRexec(pathToKeyFile, storeServerUser)
	if s.gameServer == nil {
		log.Fatal("could not create rexec for store server")
	}

	s.sc = servercontrol.NewServercontrol(hcloudToken)
	if s.sc == nil {
		log.Fatal("could not create server control")
	}

	s.storeServerIP = storeServerIP

	s.routes()

	s.cookieStore = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))

	var err error
	s.storage, err = storage.NewStorage(os.Getenv("DATABASE"), os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
