package auth

import (
	"net/http"
	"webserver/auth/handlers"
)

func DefineRoutes(r *http.ServeMux) {
	r.Handle("POST /login", handlers.NewLoginHandler())
	r.Handle("POST /signup", handlers.NewSignUpHandler())
}
