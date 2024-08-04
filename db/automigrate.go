package db

import (
	"log"
	"webserver/models"

	"gorm.io/gorm"
)

func MigrateModles(db *gorm.DB) {
	models := []interface{}{
		&models.Message{},
		&models.User{},
		&models.Channel{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}
}
