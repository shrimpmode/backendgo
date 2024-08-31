package handlers

import (
	"encoding/json"
	"net/http"
	"webserver/middleware"
	"webserver/models"
	"webserver/servers/requests"

	"gorm.io/gorm"
)

func CreateServer(db *gorm.DB) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		createServerRequest := requests.CreateServerRequest{}

		err := json.NewDecoder(r.Body).Decode(&createServerRequest)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
		}

		server := models.Server{
			Name: createServerRequest.Name,
		}

		result := db.Create(&server)
		if result.Error != nil {
			http.Error(w, "Failed to create server", http.StatusInternalServerError)
		}
	}

	return middleware.Chain(handler, middleware.JwtAuthenticated(db))
}
