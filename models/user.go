package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm: "unique"`
	DisplayName string
	Password    string `json:"-"`
	Messages    []Message
	Email       string `gorm:"unique"`
}
