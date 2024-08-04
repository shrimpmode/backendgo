package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"unique"`
	DisplayName string
	Password    string `gorm:"column:password" json:"-"`
	Messages    []Message
	Email       string `gorm:"unique"`

	Channels []*Channel `gorm:"many2many:user_channels;"`
}
