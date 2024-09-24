package auth

import (
	"net/http"
	"webserver/auth/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods(http.MethodPost)
	r.HandleFunc("/signup", handlers.SignUpHandler(db)).Methods(http.MethodPost)
}
