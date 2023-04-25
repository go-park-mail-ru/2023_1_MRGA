package app

import "time"

type Message struct {
	SenderId   uint
	ReceiverId uint
	Content    string
	SentAt     time.Time
}

type Messages struct {
	Messages []Message
}
