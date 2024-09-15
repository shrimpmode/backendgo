package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/chat/requests"
	"webserver/middleware"
	"webserver/models"

	"gorm.io/gorm"
)

type CreateChatHandler struct {
	db *gorm.DB
}

func (h *CreateChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	chatRequest := requests.CreateChatRequest{}
	server := models.Server{}

	if err := json.NewDecoder(r.Body).Decode(&chatRequest); err != nil {
		http.Error(w, CreateChatInvalidRequest.Msg, http.StatusBadRequest)
		log.Println(CreateChatInvalidRequest.Msg)
		return
	}

	if result := h.db.First(&server, chatRequest.ServerID); result.Error != nil {
		http.Error(w, CreateChatInternalError.Msg, http.StatusInternalServerError)
		log.Println(CreateChatInternalError.Msg)
		return
	}

	if !CanCreateChat(user, &server) {
		http.Error(w, CreateChatForbidden.Msg, http.StatusForbidden)
		log.Println(CreateChatForbidden.Msg)
		return
	}

	chat := models.Chat{
		Name:     chatRequest.Name,
		ServerID: server.ID,
	}

	if res := h.db.Create(&chat); res.Error != nil {
		http.Error(w, CreateChatInternalError.Msg, http.StatusInternalServerError)
		log.Println(CreateChatInternalError.Msg)
		return
	}

	log.Printf("Chat created. User: %d | Server: %d", user.ID, server.ID)
	json.NewEncoder(w).Encode(&chat)
}

func NewChatHandler(db *gorm.DB) *CreateChatHandler {
	handler := CreateChatHandler{db}
	return &handler
}
