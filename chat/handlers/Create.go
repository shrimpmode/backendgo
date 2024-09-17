package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webserver/chat/inputs"
	"webserver/middleware"
	"webserver/models"

	"gorm.io/gorm"
)

type CreateChatHandler struct {
	db      *gorm.DB
	service *ChatService
}

func GetInput(r io.Reader) (*inputs.CreateChatInput, error) {
	input := inputs.CreateChatInput{}
	err := json.NewDecoder(r).Decode(&input)
	return &input, err
}

func (h *CreateChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)

	input, err := GetInput(r.Body)
	if err != nil {
		// http.Error(w, CreateChatInvalidRequest.Msg, http.StatusBadRequest)
		// log.Println(CreateChatInvalidRequest.Msg)
		return
	}

	chat, err := h.service.CreateChat(input, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf(" User %d created Chat %d.", user.ID, chat.ID)
	json.NewEncoder(w).Encode(&chat)
}

func NewChatHandler(db *gorm.DB) *CreateChatHandler {
	handler := CreateChatHandler{db, NewChatService(db)}
	return &handler
}
