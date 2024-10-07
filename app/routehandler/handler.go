package routehandler

import (
	"net/http"
	"webserver/app"
	"webserver/auth/jwt"
	"webserver/middleware"

	"gorm.io/gorm"
)

func NewRouteHandler(handler app.Handler, db *gorm.DB) http.Handler {
	ctx := app.NewContext()
	authenticator := jwt.NewJWTAuthenticator(db)
	h := middleware.NewAuthMiddleware(handler, db, authenticator, ctx)
	return h
}
