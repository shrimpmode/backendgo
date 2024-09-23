package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/auth/jwt"
	"webserver/middleware"
	"webserver/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetChat(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chats := []models.Chat{}
		vars := mux.Vars(r)

		where := map[string]interface{}{
			"server_id": vars["serverId"],
		}
		db.Where(where).Find(&chats)
		fmt.Println("Get chats", vars, chats)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chats)
	}
}

type GetChatHandler struct {
	db   *gorm.DB
	user *models.User
}

func (h *GetChatHandler) SetUser(user *models.User) {
	h.user = user
}

func (h *GetChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chats := []models.Chat{}
	vars := mux.Vars(r)

	where := map[string]interface{}{
		"server_id": vars["serverId"],
	}
	h.db.Where(where).Find(&chats)
	fmt.Println("Get chats", vars, chats)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chats)
}

func NewGetChatHandler(db *gorm.DB) http.Handler {
	authenticator := jwt.NewJWTAuthenticator(db)
	handler := GetChatHandler{db: db}

	return middleware.NewAuthUserMiddleware(
		&handler,
		db,
		authenticator,
	)
}
