package app

import "time"

type MessageData struct {
	SenderId   uint
	Content    string
	SentAt     time.Time
	ReadStatus bool
}

type CreateChatResponse struct {
	ChatId uint
}

type GetChatsListRequest struct {
	UserId uint
}

type GetChatRequest struct {
	ChatId uint
	UserId uint
}
