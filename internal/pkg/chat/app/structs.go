package app

import (
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
)

type Message struct {
	SenderId   uint      `structs:"senderId"`
	Content    string    `structs:"content"`
	SentAt     time.Time `structs:"sentAt"`
	ReadStatus bool      `structs:"readStatus"`
}

type MessageResponse struct {
	MsgId       uint                  `structs:"msgId"`
	SenderId    uint                  `structs:"senderId"`
	Content     string                `structs:"content"`
	SentAt      string                `structs:"sentAt"`
	ReadStatus  bool                  `structs:"readStatus"`
	MessageType constants.MessageType `structs:"messageType"`
	Path        string                `structs:"path"`
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
	Content     string
	UserIds     []uint64
	MessageType constants.MessageType
	Path        string
}

type InitialMessageData struct {
	Message
	MessageType constants.MessageType
	Path        string
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

type WSMsgData struct {
	UserIds []uint64
	MsgData WSMessageResponse
}

type WSMessageResponse struct {
	SentAt      string `json:"sentAt"`
	ChatId      uint64 `json:"chatId"`
	MsgId       uint64 `json:"msgId"`
	SenderId    uint64 `json:"senderId"`
	Msg         string `json:"msg"`
	MessageType string `json:"msgType"`
	Path        string `json:"path"`
}

type WSSendResponse struct {
	Flag string            `json:"flag"`
	Body WSMessageResponse `json:"body"`
}
