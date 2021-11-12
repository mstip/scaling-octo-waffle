package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	index(wr, req)
	if wr.Code != http.StatusOK {
		t.Error("status not ok")
	}

	if !strings.Contains(wr.Body.String(), "<h1>Todo</h1>") {
		t.Error("response dont contain todo header")
	}
}

func TestCreateAndDelete(t *testing.T) {
	wr := httptest.NewRecorder()
	form := url.Values{}
	form.Add("text", "this is my new test todo")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	create(wr, req)
	if wr.Code != http.StatusMovedPermanently {
		t.Error("status not ok")
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	index(wr, req)
	if wr.Code != http.StatusOK {
		t.Error("status not ok")
	}

	if !strings.Contains(wr.Body.String(), "this is my new test todo") {
		t.Error("response dont contain new todo")
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/delete/0", nil)

	delete(wr, req)
	if wr.Code != http.StatusMovedPermanently {
		t.Error("status not ok")
	}

}

func TestUpdate(t *testing.T) {
	wr := httptest.NewRecorder()
	form := url.Values{}
	form.Add("text", "this is my new test todo")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	create(wr, req)
	if wr.Code != http.StatusMovedPermanently {
		t.Error("status not ok")
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	index(wr, req)
	if wr.Code != http.StatusOK {
		t.Error("status not ok")
	}

	if !strings.Contains(wr.Body.String(), "this is my new test todo") {
		t.Error("response dont contain new todo")
	}

	wr = httptest.NewRecorder()
	form = url.Values{}
	form.Add("text", "this is the updated text")
	form.Add("done", "on")

	req = httptest.NewRequest(http.MethodPost, "/update/0", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	update(wr, req)
	if wr.Code != http.StatusMovedPermanently {
		t.Error("status not ok")
	}

	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	index(wr, req)
	if wr.Code != http.StatusOK {
		t.Error("status not ok")
	}

	if strings.Contains(wr.Body.String(), "this is my new test todo") {
		t.Error("response still have the old todo text")
	}

	if !strings.Contains(wr.Body.String(), "this is the updated text") {
		t.Error("response dont contain the new todo text")
	}
}
