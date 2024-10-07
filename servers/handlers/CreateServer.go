package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/app"
	"webserver/app/routehandler"
	"webserver/models"
	"webserver/servers/requests"

	"gorm.io/gorm"
)

type ServerService struct {
	DB *gorm.DB
}

func (s *ServerService) Create(name string, owner *models.User) error {
	server := &models.Server{
		Name:    name,
		OwnerID: owner.ID,
	}
	return s.DB.Create(&server).Error
}

type ServerHandler struct {
	Service *ServerService
}

func (h *ServerHandler) Handle(w http.ResponseWriter, r *http.Request, ctx *app.Context) {
	createServerRequest := requests.CreateServerRequest{}

	err := json.NewDecoder(r.Body).Decode(&createServerRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	error := h.Service.Create(createServerRequest.Name, ctx.User)
	if error != nil {
		http.Error(w, "Failed to create server", http.StatusInternalServerError)
	}
}

func NewCreateServerHandler(db *gorm.DB) http.Handler {
	h := &ServerHandler{
		Service: &ServerService{
			DB: db,
		},
	}

	return routehandler.NewRouteHandler(h, db)
}
