package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string
	DisplayName string
	Password    string
	Messages    []Message
	Email       string `gorm:"unique"`
}
