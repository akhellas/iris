package adapters

import (
	"net/http"
)

// Adapter yype
type Adapter func(http.Handler) http.Handler

// Adapt function
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
