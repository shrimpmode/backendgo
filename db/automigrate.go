package db

import (
	"webserver/models"

	"gorm.io/gorm"
)

func MigrateModles(db *gorm.DB) {
	models := []interface{}{
		&models.Message{},
		&models.User{},
		&models.Chat{},
	}

	for _, model := range models {
		db.AutoMigrate(
			model,
		)
	}
}
