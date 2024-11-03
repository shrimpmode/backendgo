package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"webserver/db"
	"webserver/errs"
	"webserver/models"

	"gorm.io/gorm"
)

type ChatRepository interface {
	GetUserServer(user *models.User, serverID string) (*models.Server, error)
	CreateChat(chat *models.Chat) error
}

type Repository struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) GetUserServer(user *models.User, serverID string) (*models.Server, error) {
	server := models.Server{
		OwnerID: user.ID,
	}

	err := repo.db.First(&server, serverID).Error
	if err != nil {
		log.Println("Error getting user server", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewAppErr(
				fmt.Sprintf("Server %v not found for user %d", serverID, user.ID),
				http.StatusNotFound,
			)
		}
		log.Println(err)
		return nil, errs.NewAppErr(
			fmt.Sprintf("Error getting server %v not found for user %d", serverID, user.ID),
			http.StatusInternalServerError,
		)
	}

	return &server, err
}

func (repo *Repository) CreateChat(chat *models.Chat) error {
	err := repo.db.Create(&chat).Error
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), db.ErrDuplicatedKeyCode) {
			return errs.NewAppErr(
				fmt.Sprintf("A chat with the name %s already exists in server %d.", chat.Name, chat.ServerID),
				http.StatusConflict,
			)
		}
		return errs.NewAppErr(
			"Error trying to create a chat",
			http.StatusInternalServerError,
		)
	}
	return err
}
