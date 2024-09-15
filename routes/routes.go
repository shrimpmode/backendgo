package routes

import (
	"webserver/auth"
	"webserver/chat"
	"webserver/messages"
	"webserver/servers"
	"webserver/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	messages.DefineRoutes(r, db)
	users.DefineRoutes(r, db)
	auth.DefineRoutes(r, db)
	servers.DefineRoutes(r, db)
	chat.DefineRoutes(r, db)

	return r
}
