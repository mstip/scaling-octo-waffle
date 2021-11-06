package web

import (
	"errors"
	"github.com/gorilla/sessions"
	testifyAssert "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWoops(t *testing.T) {
	assert := testifyAssert.New(t)
	s := server{}
	w := httptest.NewRecorder()
	s.woops(w, errors.New("test error"))
	assert.Equal(w.Code, http.StatusInternalServerError)
	assert.Contains(w.Body.String(), "test error")
}

func TestValidateError(t *testing.T) {
	assert := testifyAssert.New(t)
	s := server{}
	w := httptest.NewRecorder()
	s.validatorError(w, "test validation error")
	assert.Equal(w.Code, http.StatusUnprocessableEntity)
	assert.Contains(w.Body.String(), "test validation error")
}

func TestGetFlashesWithoutEntry(t *testing.T) {
	assert := testifyAssert.New(t)
	s := server{}
	s.cookieStore = sessions.NewCookieStore([]byte("geheim"))
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	flashes, err := s.getFlashes(w, r)
	if assert.NoError(err) {
		var expectedFlashes []string
		assert.Equal(expectedFlashes, flashes)
	}
}

func TestAddFlashes(t *testing.T) {
	assert := testifyAssert.New(t)
	s := server{}
	s.cookieStore = sessions.NewCookieStore([]byte("geheim"))
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	err := s.addFlash(w, r, "test flash")
	if assert.NoError(err) {
		flashes, err := s.getFlashes(w, r)
		if assert.NoError(err) {
			assert.Equal([]string{"test flash"}, flashes)
		}
	}
}
