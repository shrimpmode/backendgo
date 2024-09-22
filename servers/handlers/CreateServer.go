package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/auth/jwt"
	"webserver/middleware"
	"webserver/models"
	"webserver/servers/requests"

	"gorm.io/gorm"
)

type CreateServerHandler struct {
	db   *gorm.DB
	user *models.User
}

func (h *CreateServerHandler) SetUser(user *models.User) {
	h.user = user
}

func (h *CreateServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	createServerRequest := requests.CreateServerRequest{}

	err := json.NewDecoder(r.Body).Decode(&createServerRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	server := models.Server{
		Name:    createServerRequest.Name,
		OwnerID: h.user.ID,
	}

	result := h.db.Create(&server)
	if result.Error != nil {
		http.Error(w, "Failed to create server", http.StatusInternalServerError)
	}
}

func NewCreateServerHandler(db *gorm.DB) http.Handler {

	handler := CreateServerHandler{db: db}
	authenticator := jwt.NewJWTAuthenticator(db)

	return middleware.NewAuthUserMiddleware(
		&handler,
		db,
		authenticator,
	)
}
