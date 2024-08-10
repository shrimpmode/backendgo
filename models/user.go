package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"unique"`
	DisplayName string
	Password    string `gorm:"column:password" json:"-"`
	Email       string `gorm:"unique"`
	Messages    []Message
	Channels    []*Channel   `gorm:"many2many:user_channels;"`
	Active      sql.NullBool `gorm:"default:false"`
}
