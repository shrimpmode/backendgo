package servers

import (
	"webserver/servers/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/server", handlers.NewCreateServerHandler(db)).Methods("POST")
}
