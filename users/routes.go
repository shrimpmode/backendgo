package users

import (
	"webserver/users/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/user", handlers.NewCreateUserHandler(db)).Methods("POST")
}
