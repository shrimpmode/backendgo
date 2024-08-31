package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Name     string `gorm:"not null"`
	ServerID uint   `gorm:"not null"`
	Messages []Message
	Members  []*User `gorm:"many2many:chat_members;"`
}
