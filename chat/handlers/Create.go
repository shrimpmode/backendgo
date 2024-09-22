package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webserver/auth/jwt"
	"webserver/chat/inputs"
	"webserver/middleware"
	"webserver/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateChatHandler struct {
	db      *gorm.DB
	user    *models.User
	service *ChatService
}

func (h *CreateChatHandler) SetUser(user *models.User) {
	h.user = user
}

func GetInput(r io.Reader) (*inputs.CreateChatInput, error) {
	input := inputs.CreateChatInput{}
	err := json.NewDecoder(r).Decode(&input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = validator.New().Struct(input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}

func (h *CreateChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	input, err := GetInput(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, AppErrBadRequest.Message, AppErrBadRequest.Code)
		return
	}

	chat, err := h.service.CreateChat(input, h.user)
	if err != nil {
		fmt.Println(err)
		appErr, ok := err.(AppError)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, appErr.Message, int(appErr.Code))
		return
	}

	log.Printf(" User %d created Chat %d.", h.user.ID, chat.ID)
	json.NewEncoder(w).Encode(&chat)
}

func NewChatHandler(db *gorm.DB) http.Handler {
	handler := CreateChatHandler{
		db:      db,
		service: NewChatService(db),
	}
	authenticator := jwt.NewJWTAuthenticator(db)
	return middleware.NewAuthUserMiddleware(&handler, db, authenticator)
}
