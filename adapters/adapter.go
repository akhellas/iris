package adapters

import (
	"net/http"
)

// Adapter type
type Adapter func(http.Handler) http.Handler

// ContextKey type
type ContextKey string

// Adapt function
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
