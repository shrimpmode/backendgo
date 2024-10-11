package middleware

import (
	"net/http"
	"webserver/app"
	"webserver/app/auth"
	"webserver/auth/jwt"
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
		return
	}
	m.handler.SetUser(user)
	m.handler.ServeHTTP(w, r)
}

func NewAuthUserMiddleware(handler app.AppHandler, db *gorm.DB, authenticator auth.Authenticator) http.Handler {
	return &AuthUserMiddleware{handler, db, authenticator}
}

type AuthMiddleware struct {
	handler       http.Handler
	db            *gorm.DB
	authenticator auth.Authenticator
	ctx           *app.Context
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := m.authenticator.GetAuthenticatedUser(r)

	if !ok {
		http.Error(w, errs.AuthUserError.Message, errs.AuthUserError.Code)
		return
	}
	m.ctx.SetUser(user)
	m.handler.ServeHTTP(w, r)
}

func NewAuthMiddleware(handler http.Handler, db *gorm.DB, authenticator auth.Authenticator, ctx *app.Context) http.Handler {
	return &AuthMiddleware{handler, db, authenticator, ctx}
}

func NewJWTMiddleware(db *gorm.DB, ctx *app.Context) MiddlewareFunc {
	authenticator := jwt.NewJWTAuthenticator(db)
	return func(h http.Handler) http.Handler {
		return &AuthMiddleware{h, db, authenticator, ctx}
	}
}
