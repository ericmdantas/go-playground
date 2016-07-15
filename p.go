package main

import (
	"net/http"
	"strings"
)

type httpCb func(w http.ResponseWriter, r *http.Request)

func newRouter() *router {
	return &router{
		Name: "rrr",
		Fns:  make(map[string]httpCb),
	}
}

type router struct {
	Name string
	Fns  map[string]httpCb
}

func (rou *router) on(url string, cb httpCb) {
	locURL := strings.Trim(url, " ")
	locURL = strings.ToLower(locURL)

	rou.Fns[locURL] = cb
}

func (rou *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Add("abc", "123")
	h.Add("Content-Type", "application/json")

	for k, v := range rou.Fns {
		if k == r.URL.Path {
			v(w, r)
		}
	}
}

func main() {
	r := newRouter()

	r.on("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("!"))
	})

	r.on("/wtf", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yo!"))
	})

	http.ListenAndServe(":1234", r)
}
