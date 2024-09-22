package middleware

import (
	"net/http"
	"webserver/app"
	"webserver/app/auth"
	"webserver/errs"

	"gorm.io/gorm"
)

type AuthUserMiddleware struct {
	handler       app.AppHandler
	db            *gorm.DB
	authenticator auth.Authenticator
}

func (m *AuthUserMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := m.authenticator.GetAuthenticatedUser(r)
	if !ok {
		http.Error(w, errs.AuthUserError.Message, errs.AuthUserError.Code)
	}

	m.handler.SetUser(user)
	m.handler.ServeHTTP(w, r)
}

func NewAuthUserMiddleware(handler app.AppHandler, db *gorm.DB, authenticator auth.Authenticator) http.Handler {
	return &AuthUserMiddleware{handler, db, authenticator}
}
