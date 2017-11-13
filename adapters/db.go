package adapters

import (
	"context"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// DbContext interface
type DbContext interface {
	session() mgo.Session
	key() ContextKey
}

// WithDB adapter
func WithDB(db DbContext) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbsession := db.session().Copy()
			defer dbsession.Close()
			ctx := context.WithValue(r.Context(), DbContextKey, dbsession)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
