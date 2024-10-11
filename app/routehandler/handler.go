package routehandler

import (
	"net/http"
	"webserver/app"
	"webserver/auth/jwt"
	"webserver/middleware"

	"gorm.io/gorm"
)

func NewRouteHandler(handler http.Handler, db *gorm.DB) http.Handler {
	ctx := app.NewContext()
	authenticator := jwt.NewJWTAuthenticator(db)
	h := middleware.NewAuthMiddleware(handler, db, authenticator, ctx)
	return h
}

func NewHandler(h app.Handler, db *gorm.DB) http.Handler {
	ctx := app.NewContext()
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r, ctx)
	})

	middlewares := []middleware.MiddlewareFunc{
		middleware.NewLogger,
		middleware.NewJWTMiddleware(db, ctx),
	}

	return middleware.ApplyMiddlewares(f, middlewares...)
}
