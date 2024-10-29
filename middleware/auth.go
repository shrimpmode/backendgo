package middleware

import (
	"net/http"
	"webserver/app"
	"webserver/auth/jwt"
	"webserver/errs"
)

type AuthMiddleware struct {
	handler http.Handler
	ctx     *app.Context
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, ok := jwt.GetAuthenticatedUser(r)

	if !ok {
		http.Error(w, errs.AuthUserError.Message, errs.AuthUserError.Code)
		return
	}
	m.ctx.SetUser(user)
	m.handler.ServeHTTP(w, r)
}

func NewAuthMiddleware(handler http.Handler, ctx *app.Context) http.Handler {
	return &AuthMiddleware{handler, ctx}
}

func NewJWTMiddleware(ctx *app.Context) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return &AuthMiddleware{h, ctx}
	}
}
