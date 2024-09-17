package chat

import (
	"fmt"
	"net/http"
	"webserver/chat/handlers"
	"webserver/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ChatHandler struct {
	Value int
}

func (c *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c.Value)
}

func DefineRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/servers/{serverId}/chats", handlers.GetChat(db)).Methods(http.MethodGet)
	r.Handle("/v2/chat", middleware.WithUser(db)(handlers.NewChatHandler(db))).Methods("POST")
}
