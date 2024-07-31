package routes

import (
	"webserver/messages"
	"webserver/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	messages.DefineRoutes(r, db)
	users.DefineRoutes(r, db)

	return r
}
