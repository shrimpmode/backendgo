package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/models"
)

type ServerHandler struct {
	Service ServerService
}

func (h *ServerHandler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	input := CreateServerRequest{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	server := &models.Server{
		Name:    input.Name,
		OwnerID: ctx.User.ID,
	}

	err = h.Service.Create(server)
	if err != nil {
		http.Error(w, "Failed to create server", http.StatusInternalServerError)
		return
	}
	log.Printf("User %v created sever %v", ctx.User.ID, server.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(server); err != nil {
		http.Error(w, "Failed to encode server", http.StatusInternalServerError)
		return
	}
}

func Create() http.Handler {
	h := &ServerHandler{
		Service: NewServerService(),
	}

	return routehandler.NewHandler(h)
}
