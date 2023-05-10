package app

import "time"

type Message struct {
	SenderId   uint      `structs:"senderId"`
	Content    string    `structs:"content"`
	SentAt     time.Time `structs:"sentAt"`
	ReadStatus bool      `structs:"readStatus"`
}

type MessageResponse struct {
	SenderId   uint   `structs:"senderId"`
	Content    string `structs:"content"`
	SentAt     string `structs:"sentAt"`
	ReadStatus bool   `structs:"readStatus"`
}

type ChatMessage struct {
	Msg         MessageResponse `structs:"msg"`
	ChatId      uint            `structs:"chatId"`
	ChatUserIds []uint          `structs:"chatUserIds"`
}

type MessageData struct {
	Msg   MessageResponse `structs:"msg"`
	MsgId uint            `structs:"msgId"`
}

type CreateChatRequest struct {
	UserIds []uint
}

type CreateChatResponse struct {
	ChatId uint `structs:"chatId"`
}

type SendMessageRequest struct {
	Content string
	UserIds []uint64
}

type SendMessageResponse struct {
	SentAt string `structs:"sentAt"`
}

type GetChatsListResponse struct {
	ChatsList []ChatMessage `structs:"chatsList"`
}

type GetChatResponse struct {
	Chat []MessageData `structs:"chat"`
}

type WSMessageRequest struct {
	SentAt  string   `json:"sentAt"`
	ChatId  uint64   `json:"chatId"`
	UserIds []uint64 `json:"userIds"`
	Msg     string   `json:"msg"`
}

type WSMessageResponse struct {
	SentAt   string `json:"sentAt"`
	ChatId   uint64 `json:"chatId"`
	SenderId uint64 `json:"senderId"`
	Msg      string `json:"msg"`
}

type WSSendResponse struct {
	Flag string            `json:"flag"`
	Body WSMessageResponse `json:"body"`
}
