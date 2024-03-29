package app

import "time"

type MessageType string

const (
	TextMessage  MessageType = "text"
	AudioMessage MessageType = "audio"
)

type Message struct {
	Id         uint      `gorm:"primary_key"`
	ChatId     uint      `gorm:"not null"`
	SenderId   uint      `gorm:"not null"`
	Content    string    `gorm:"not null"`
	SentAt     time.Time `gorm:"not null; type:timestamp without time zone"`
	ReadStatus bool      `gorm:"not null"`
}

type Chat struct {
	Id uint `gorm:"primary_key"`
}

type ChatUser struct {
	Id     uint `gorm:"primary_key"`
	ChatId uint `gorm:"not null"`
	UserId uint `gorm:"not null"`
}

type Media struct {
	Id          uint   `gorm:"primary_key"`
	MessageId   uint   `gorm:"not null"`
	MessageType string `gorm:"not null"`
	Path        string `gorm:"not null"`
}
