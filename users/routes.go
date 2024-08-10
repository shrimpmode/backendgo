package users

import (
	"webserver/middleware"
	"webserver/users/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	r.HandleFunc(
		"/user",
		middleware.Chain(
			handlers.CreateUser(db),
			middleware.IsAuthenticated(db, store),
			middleware.Logging,
		),
	).Methods("POST")
}
