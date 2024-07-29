package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string
	DisplayName string
	Messages    []Message
}
