package service

import (
	"testing"
	"webserver/chat/inputs"
	"webserver/models"
)

type RepositoryMock struct{}

func (r *RepositoryMock) GetUserServer(user *models.User, serverID string) (*models.Server, error) {
	return &models.Server{}, nil
}

func (r *RepositoryMock) CreateChat(chat *models.Chat) error {
	return nil
}

func TestCreateChat(t *testing.T) {
	service := &ChatService{
		repo: &RepositoryMock{},
	}
	input := &inputs.CreateChatInput{
		Name:     "test",
		ServerID: "1",
	}
	got, err := service.CreateChat(input, &models.User{})
	if err != nil {
		t.Errorf("error creating chat: %v", err)
	}
	t.Logf("got: %v", got)
}
