package models

type Chat struct {
	Model
	Name     string `gorm:"not null" json:"name"`
	ServerID uint   `gorm:"not null" json:"serverId`
	Messages []Message
	Members  []*User `gorm:"many2many:chat_members;"`
}
