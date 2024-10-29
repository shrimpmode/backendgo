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
	api := r.PathPrefix("/api").Subrouter()
	messages.DefineRoutes(api, db)
	users.DefineRoutes(api)
	auth.DefineRoutes(api)
	servers.DefineRoutes(api, db)
	chat.DefineRoutes(api, db)

	return r
}
