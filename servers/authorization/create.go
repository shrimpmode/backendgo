package authorization

import (
	"webserver/models"

	"gorm.io/gorm"
)

func CanCreateServer(user *models.User, server *models.Server, db *gorm.DB) bool {
	return server.OwnerID == user.ID
}
