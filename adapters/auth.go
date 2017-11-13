package adapters

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// AuthContext interface
type AuthContext interface {
	signingKey() string
	userKey() ContextKey
}

// WithAuth adapter
func WithAuth(authContext AuthContext) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return authContext.signingKey(), nil
			})

			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), authContext.userKey(), claims["sub"])
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprint(w, "Token is not valid")
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Unauthorised access to this resource")
			}
		})
	}
}
