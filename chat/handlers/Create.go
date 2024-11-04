package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/chat/inputs"
	"webserver/chat/service"
	"webserver/errs"

	"github.com/go-playground/validator/v10"
)

type CreateChatHandler struct {
	service *service.ChatService
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

func (h *CreateChatHandler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	input, err := GetInput(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, errs.AppErrBadRequest.Message, errs.AppErrBadRequest.Code)
		return
	}

	chat, err := h.service.CreateChat(input, ctx.User)
	if err != nil {
		fmt.Println(err)
		appErr, ok := err.(*errs.AppError)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, appErr.Message, int(appErr.Code))
		return
	}

	log.Printf(" User %d created Chat %d.", ctx.User.ID, chat.ID)
	json.NewEncoder(w).Encode(&chat)
}

func Create() http.Handler {
	handler := &CreateChatHandler{
		service: service.NewChatService(),
	}

	return routehandler.NewHandler(handler)
}
