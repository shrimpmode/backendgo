package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Chats   []Chat
	Name    string `gorm:"not null;uniqueIndex:idx_server_name_owner"`
	OwnerID uint   `gorm:"not null;uniqueIndex:idx_server_name_owner"`
}
