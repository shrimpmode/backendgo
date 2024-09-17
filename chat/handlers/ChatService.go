package handlers

import (
	"webserver/chat/inputs"
	"webserver/models"

	"gorm.io/gorm"
)

type ChatService struct {
	repo *ChatRepo
}

func (s *ChatService) CreateChat(input *inputs.CreateChatInput, user *models.User) (*models.Chat, error) {
	server, err := s.repo.GetUserServer(user, input.ServerID)
	if err != nil {
		return nil, err
	}
	chat := models.Chat{
		Name:     input.Name,
		ServerID: server.ID,
	}
	return &chat, s.repo.CreateChat(&chat)
}

func NewChatService(db *gorm.DB) *ChatService {
	return &ChatService{
		&ChatRepo{db},
	}
}
