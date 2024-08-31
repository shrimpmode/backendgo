package db

import (
	"log"
	"webserver/models"

	"gorm.io/gorm"
)

func MigrateModles(db *gorm.DB) {
	appModels := []interface{}{
		&models.Message{},
		&models.User{},
		&models.Channel{},
		&models.Server{},
		&models.Chat{},
	}

	if err := db.AutoMigrate(appModels...); err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}
}
