package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"webserver/db"
	"webserver/models"

	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{db}
}

func (repo *ChatRepo) GetUserServer(user *models.User, serverID string) (*models.Server, error) {
	server := models.Server{
		OwnerID: user.ID,
	}

	err := repo.db.First(&server, serverID).Error

	if err != nil {
		log.Println("Error getting user server", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, AppError{
				fmt.Sprintf("Server %v not found for user %d", serverID, user.ID),
				http.StatusNotFound,
			}
		}
		log.Println(err)
		return nil, AppError{
			fmt.Sprintf("Error getting server %v not found for user %d", serverID, user.ID),
			http.StatusInternalServerError,
		}
	}

	return &server, err

}

func (repo *ChatRepo) CreateChat(chat *models.Chat) error {
	err := repo.db.Create(&chat).Error
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), db.ErrDuplicatedKeyCode) {
			return AppError{
				fmt.Sprintf("A chat with the name %s already exists in server %d.", chat.Name, chat.ServerID),
				http.StatusConflict,
			}
		}
		return AppError{
			"Error trying to create a chat",
			http.StatusInternalServerError,
		}
	}
	return err
}
