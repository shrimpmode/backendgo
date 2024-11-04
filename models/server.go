package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Chats   []Chat `json:"chats"`
	Name    string `json:"name" gorm:"not null;uniqueIndex:idx_server_name_owner"`
	OwnerID uint   `json:"owner_id" gorm:"not null;uniqueIndex:idx_server_name_owner"`
}
