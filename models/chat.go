package models

type Chat struct {
	Model
	Name     string `gorm:"not null;uniqueIndex:idx_server_chat_name" json:"name"`
	ServerID uint   `gorm:"not null;uniqueIndex:idx_server_chat_name" json:"server_id"`
	Messages []Message
	Members  []*User `gorm:"many2many:chat_members;"`
}
