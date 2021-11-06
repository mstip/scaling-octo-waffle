package web

import (
	"net/http"
	"net/url"
)

func SetSessionCookie(sessionId string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:     "ndis_sess",
		Value:    url.QueryEscape(sessionId),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func GetSessionId(r *http.Request) (string, error) {
	c, err := r.Cookie("ndis_sess")
	if err != nil {
		return "", err
	}
	return url.QueryUnescape(c.Value)
}