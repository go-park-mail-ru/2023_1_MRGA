package app

import "time"

type MessageData struct {
	SenderId   uint
	Content    string
	SentAt     time.Time
	ReadStatus bool
}

type ChatMessage struct {
	Message
	MessageType MessageType
	Path        string
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

type MessageWithChatUsers struct {
	ChatMessage
	ChatUserIds []uint
}
