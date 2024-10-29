package users

import (
	"webserver/users/handlers"

	"github.com/gorilla/mux"
)

func DefineRoutes(r *mux.Router) {
	r.Handle("/user", handlers.Create()).Methods("POST")
}
