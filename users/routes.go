package users

import (
	"webserver/middleware"
	"webserver/users/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc(
		"/user",
		middleware.Chain(
			handlers.CreateUser(db),
			middleware.Logging,
		),
	).Methods("POST")
}
