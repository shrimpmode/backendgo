package handlers

import (
	"webserver/db"
	"webserver/models"

	"gorm.io/gorm"
)

type ServerService interface {
	Create(*models.Server) error
}

type serverService struct {
	DB *gorm.DB
}

func (s *serverService) Create(server *models.Server) error {
	return s.DB.Create(&server).Error
}

func NewServerService() ServerService {
	return &serverService{DB: db.GetDB()}
}
