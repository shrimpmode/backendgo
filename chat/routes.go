package chat

import (
	"net/http"
	"webserver/chat/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/server/{serverId}/chat", handlers.NewGetChatHandler(db)).Methods(http.MethodGet)
	r.Handle("/chat", handlers.NewChatHandler(db)).Methods(http.MethodPost)
}
