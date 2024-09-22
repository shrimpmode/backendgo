package middleware

import (
	"context"
	"log"
	"net/http"
	"webserver/app"
	errorsAuth "webserver/auth/errors"
	"webserver/auth/jwt"
	"webserver/errs"

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

type AuthUserMiddleware struct {
	handler app.AppHandler
	db      *gorm.DB
}

func (m *AuthUserMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.GetAuthenticatedUser(m.db, r)
	if !ok {
		http.Error(w, errs.AuthUserError.Message, errs.AuthUserError.Code)
	}

	m.handler.SetUser(user)
	m.handler.ServeHTTP(w, r)
}

func NewAuthUserMiddleware(handler app.AppHandler, db *gorm.DB) http.Handler {
	return &AuthUserMiddleware{handler, db}
}
