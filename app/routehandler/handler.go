package routehandler

import (
	"net/http"
	"webserver/app"
	"webserver/middleware"

	"gorm.io/gorm"
)

func NewHandler(h app.Handler, db *gorm.DB) http.Handler {
	ctx := app.NewContext()
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r, ctx)
	})

	middlewares := []middleware.MiddlewareFunc{
		middleware.NewLogger,
		middleware.NewJWTMiddleware(ctx),
	}

	return middleware.ApplyMiddlewares(f, middlewares...)
}
