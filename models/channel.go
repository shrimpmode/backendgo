package models

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	name string

	Users []*User `gorm:"many2many:user_channels;"`
}
