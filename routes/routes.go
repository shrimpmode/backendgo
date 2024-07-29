package routes

import (
	"webserver/messages"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	messages.DefineRoutes(r)
	return r
}
