package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type GetChatHandler struct {
	db *gorm.DB
}

func (h *GetChatHandler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	chats := []models.Chat{}
	vars := mux.Vars(r)

	where := map[string]interface{}{
		"server_id": vars["serverId"],
	}
	h.db.Where(where).Find(&chats)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chats)
}

func NewGetChatHandler(db *gorm.DB) http.Handler {
	h := &GetChatHandler{db: db}

	return routehandler.NewHandler(h)
}
