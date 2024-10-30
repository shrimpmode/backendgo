package auth

import (
	"net/http"
	"webserver/auth/handlers"

	"github.com/gorilla/mux"
)

func DefineRoutes(r *mux.Router) {
	r.Handle("/login", handlers.Login()).Methods(http.MethodPost)
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods(http.MethodPost)
}
