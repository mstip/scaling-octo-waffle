package web

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	testifyAssert "github.com/stretchr/testify/assert"
	"github.com/ungefaehrlich/ppu_gaming/pkg/servercontrol"
	"github.com/ungefaehrlich/ppu_gaming/pkg/storage"
)

type mockRexec struct {
}

func (m mockRexec) Exec(host string, cmd string) (string, error) {
	return "", nil
}

type mockServercontrol struct {
}

func (m mockServercontrol) GetAllServers() ([]servercontrol.ServerInfo, error) {
	var info []servercontrol.ServerInfo
	return info, nil
}

func (m mockServercontrol) CreateServer(name string, serverType int) (*servercontrol.ServerInfo, error) {
	panic("implement me")
}

func (m mockServercontrol) DeleteServer(ID int) error {
	panic("implement me")
}

func (m mockServercontrol) GetServerById(ID int) (servercontrol.ServerInfo, error) {
	panic("implement me")
}

func setupServerTest() *server {
	s := &server{}
	s.tmplDir = "../../web/template/"
	s.gameServer = &mockRexec{}
	s.storeServer = &mockRexec{}
	s.sc = &mockServercontrol{}
	s.storeServerIP = "1.3.3.7"
	s.routes()
	s.cookieStore = sessions.NewCookieStore([]byte("test-secret"))
	s.storage = storage.NewInMemStorage()
	return s
}

func TestOverview(t *testing.T) {
	assert := testifyAssert.New(t)
	s := setupServerTest()
	assert.HTTPSuccess(s.handleOverview(), "GET", "/overview", nil)
	body := testifyAssert.HTTPBody(s.handleOverview(), "GET", "/overview", nil)
	assert.Contains(body, "Create Server")
	assert.Contains(body, "Servers")
	assert.Contains(body, "Saves")
}

func TestIndex(t *testing.T) {
	assert := testifyAssert.New(t)
	s := setupServerTest()
	assert.HTTPSuccess(s.handleIndex(), "GET", "/", nil)
	body := testifyAssert.HTTPBody(s.handleIndex(), "GET", "/", nil)
	assert.Contains(body, "PPU Gaming")
	assert.Contains(body, "Login")
	assert.Contains(body, "Register")

}

func TestShowRegister(t *testing.T) {
	assert := testifyAssert.New(t)
	s := setupServerTest()
	assert.HTTPSuccess(s.handleShowRegister(), "GET", "/register", nil)
	body := testifyAssert.HTTPBody(s.handleShowRegister(), "GET", "/register", nil)
	assert.Contains(body, "Register")
	assert.Contains(body, "Name")
	assert.Contains(body, "Email")
	assert.Contains(body, "Password")
}

func TestRegister(t *testing.T) {
	assert := testifyAssert.New(t)
	s := setupServerTest()
	params := url.Values{}
	params.Add("name", "test-name")
	params.Add("email", "test@test.de")
	params.Add("password", "qqqq")
	r := httptest.NewRequest("POST", "/register", strings.NewReader(params.Encode()))
	w := httptest.NewRecorder()

	s.handleRegister()(w, r)

	assert.Equal(301, w.Code)
	assert.Equal("/login", w.HeaderMap.Get("Location"))

	user, err := s.storage.User.FindByEmail("test@test.de")
	assert.NoError(err)
	assert.Equal("test-name", user.Name)
}
