package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Chats   []Chat
	Name    string `gorm:"not null"`
	OwnerID uint   `gorm:"not null"`
}
