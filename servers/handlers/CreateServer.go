package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/servers/requests"
	"webserver/servers/services"

	"gorm.io/gorm"
)

type ServerHandler struct {
	Service *services.ServerService
}

func (h *ServerHandler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	createServerRequest := requests.CreateServerRequest{}

	err := json.NewDecoder(r.Body).Decode(&createServerRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	server, error := h.Service.Create(createServerRequest.Name, ctx.User)
	if error != nil {
		http.Error(w, "Failed to create server", http.StatusInternalServerError)
	}
	log.Printf("User %v created sever %v", ctx.User.ID, server.ID)
}

func NewCreateServerHandler(db *gorm.DB) http.Handler {
	h := &ServerHandler{
		Service: &services.ServerService{
			DB: db,
		},
	}

	return routehandler.NewHandler(h, db)
}
