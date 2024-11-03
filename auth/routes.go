package auth

import (
	"net/http"
	"webserver/auth/handlers"
)

func DefineRoutes(r *http.ServeMux) {
	r.Handle("POST /login", handlers.Login())
	r.HandleFunc("POST /signup", handlers.SignUpHandler)
}
