package service

import (
	"webserver/chat/inputs"
	"webserver/db"
	"webserver/models"
)

type ChatService struct {
	repo ChatRepository
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

func NewChatService() *ChatService {
	return &ChatService{
		repo: &Repository{db.GetDB()},
	}
}
