package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

type app struct {
	session *mgo.Session
	auth    *auth
	api     *api
}

func (h *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, tail := shiftPath(r.URL.Path)
	switch head {
	case "auth":
		r.URL.Path = tail
		s := adapters.Adapt(h.auth, adapters.WithDB(h.session), adapters.WithLog(), adapters.WithCors())
		s.ServeHTTP(w, r)
	case "api":
		r.URL.Path = tail
		s := adapters.Adapt(h.api, adapters.WithAuth(), adapters.WithDB(h.session), adapters.WithLog(), adapters.WithCors())
		s.ServeHTTP(w, r)
	default:
		//staticHandler(w, r)
	}
}

func (h *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "buildings":
		buildingsHandler(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
