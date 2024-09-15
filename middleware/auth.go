package middleware

import (
	"context"
	"log"
	"net/http"
	errorsAuth "webserver/auth/errors"
	"webserver/auth/jwt"

	"gorm.io/gorm"
)

type key int

const (
	UserKey key = iota
	DBKey
)

func WithUser(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := jwt.GetAuthenticatedUser(db, r)
			if !ok {
				http.Error(w, errorsAuth.AuthenticationError, http.StatusUnauthorized)
				log.Println(errorsAuth.AuthenticationError)
				return
			}
			ctx := context.WithValue(r.Context(), UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
