package routes

import (
	"webserver/auth"
	"webserver/messages"
	"webserver/users"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, store *sessions.CookieStore) *mux.Router {
	r := mux.NewRouter()
	messages.DefineRoutes(r, db, store)
	users.DefineRoutes(r, db, store)
	auth.DefineRoutes(r, db, store)

	return r
}
