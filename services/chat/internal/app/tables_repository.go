package app

import "time"

type Message struct {
	ID         uint      `gorm:"primary_key"`
	SenderId   uint      `gorm:"not null"`
	ReceiverId uint      `gorm:"not null"`
	Content    string    `gorm:"not null"`
	SentAt     time.Time `gorm:"not null; type:timestamp without time zone"`
}
