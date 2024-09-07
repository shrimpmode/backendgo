package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/auth/jwt"
	"webserver/chat/requests"
	"webserver/middleware"
	"webserver/models"
	"webserver/servers/authorization"

	"gorm.io/gorm"
)

func CreateChat(db *gorm.DB) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		chatRequest := requests.CreateChatRequest{}
		server := models.Server{}

		err := json.NewDecoder(r.Body).Decode(&chatRequest)
		if err != nil {
			http.Error(w, "Error creating chat: Invalid Request", http.StatusBadRequest)
			return
		}

		user, ok := jwt.GetAuthenticatedUser(db, r)
		if !ok {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if result := db.Where("ID = ?", chatRequest.ServerID).First(&server); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Server not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Error creating server", http.StatusInternalServerError)
			return
		}

		if !authorization.CanCreateServer(user, &server, db) {
			http.Error(w, "Invalid authorization", http.StatusForbidden)
			return
		}

		chat := models.Chat{
			Name:     chatRequest.Name,
			ServerID: server.ID,
		}
		if result := db.Create(&chat); result.Error != nil {
			http.Error(w, "Failed to created server", http.StatusInternalServerError)
			return
		}

		log.Printf("Chat created successfully by user %v", user.ID)

		json.NewEncoder(w).Encode(chat)
	}

	return middleware.Chain(f, middleware.JwtAuthenticated(db))
}
