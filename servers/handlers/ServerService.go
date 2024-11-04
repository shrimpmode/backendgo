package handlers

import (
	"webserver/models"

	"gorm.io/gorm"
)

type ServerService struct {
	DB *gorm.DB
}

func (s *ServerService) Create(name string, owner *models.User) (*models.Server, error) {
	server := &models.Server{
		Name:    name,
		OwnerID: owner.ID,
	}
	return server, s.DB.Create(&server).Error
}
