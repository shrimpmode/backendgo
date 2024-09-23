package chat

import (
	"net/http"
	"webserver/chat/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/servers/{serverId}/chats", handlers.NewGetChatHandler(db)).Methods(http.MethodGet)
	r.Handle("/v2/chat", handlers.NewChatHandler(db)).Methods(http.MethodPost)
}
