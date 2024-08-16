package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"unique;not null;default:null"`
	DisplayName string `gorm:"not null"`
	Password    string `gorm:"column:password;not null;" json:"-"`
	Email       string `gorm:"unique;not null"`
	Messages    []Message
	Channels    []*Channel   `gorm:"many2many:user_channels;"`
	Active      sql.NullBool `gorm:"default:false"`
}
