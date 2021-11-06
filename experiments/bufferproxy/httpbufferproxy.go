package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type BufferProxy struct {
	Addr       string
	TargetHost string
	Tries      int
	WaitMs     int
}

func (b *BufferProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s - %s - %s", r.RemoteAddr, r.Method, r.URL.Path, r.URL.RawQuery)

	url := b.TargetHost + r.URL.Path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}

	client := &http.Client{}
	var resp *http.Response
	var err error
	for n := 0; n < b.Tries; n++ {
		req, err := http.NewRequest(http.MethodGet, url, r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err = client.Do(req)

		if err == nil {
			break
		}
		log.Printf("retry - %s", err.Error())
		time.Sleep(time.Duration(b.WaitMs) * time.Millisecond)
		if n == b.Tries-1 {
			log.Println("give up")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if resp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, headerValues := range resp.Header {
		for _, v := range headerValues {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (b *BufferProxy) Run() {
	log.Printf("Forward %s to %s - wait %dms - tries %d", b.Addr, b.TargetHost, b.WaitMs, b.Tries)
	log.Println(http.ListenAndServe(":3000", b))
}

func main() {
	// TODO: settings via env
	// TODO: better logging when giveup
	bp := BufferProxy{
		Addr:       ":3000",
		TargetHost: "http://localhost:31337",
		Tries:      10,
		WaitMs:     100,
	}
	bp.Run()
}
