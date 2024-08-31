package auth

import (
	"webserver/auth/handlers"
	"webserver/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/login", middleware.Chain(handlers.LoginHandler(db), middleware.Logging)).Methods("POST")
	r.HandleFunc("/signup", middleware.Chain(handlers.SignUpHandler(db), middleware.Logging)).Methods("POST")
}
