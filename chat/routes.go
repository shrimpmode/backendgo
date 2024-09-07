package chat

import (
	"net/http"
	"webserver/chat/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/chat", handlers.CreateChat(db)).Methods("POST")
	r.Handle("/servers/{serverId}/chats", handlers.GetChat(db)).Methods(http.MethodGet)

	// r.Handle("/servers/{serverId}/chats/{chatId}", handlers.GetChat(db)).Methods("GET")
}
