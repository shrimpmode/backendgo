package auth

import (
	"webserver/auth/handlers"
	"webserver/middleware"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB, store *sessions.CookieStore) {
	r.HandleFunc("/login", middleware.Chain(handlers.LoginHandler(db, store))).Methods("POST")
}
