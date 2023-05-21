package app

import (
	"encoding/json"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
)

type Message struct {
	SenderId    uint                  `structs:"senderId"`
	Content     string                `structs:"content"`
	SentAt      time.Time             `structs:"sentAt"`
	ReadStatus  bool                  `structs:"readStatus"`
	MessageType constants.MessageType `structs:"messageType"`
	Path        string                `structs:"path"`
}

type MessageResponse struct {
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

type MessageResponseWithId struct {
	MsgId       uint                  `structs:"msgId"`
	SenderId    uint                  `structs:"senderId"`
	Content     string                `structs:"content"`
	SentAt      string                `structs:"sentAt"`
	ReadStatus  bool                  `structs:"readStatus"`
	MessageType constants.MessageType `structs:"messageType"`
	Path        string                `structs:"path"`
}

type MessageData struct {
	Msg MessageResponseWithId `structs:"msg"`
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
}

type SendMessageResponse struct {
	SentAt string `structs:"sentAt"`
	MsgId  uint64 `structs:"msgId"`
}

type GetChatsListResponse struct {
	ChatsList []ChatMessage `structs:"chatsList"`
}

type GetChatResponse struct {
	Chat []MessageData `structs:"chat"`
}

type WSMsgData struct {
	SenderId uint64
	UserIds  []uint64
	MsgData  interface{}
}

type WSMessageResponse struct {
	SentAt      string `json:"sentAt"`
	ChatId      uint64 `json:"chatId"`
	MsgId       uint64 `json:"msgId"`
	SenderId    uint64 `json:"senderId"`
	Msg         string `json:"msg"`
	MessageType string `json:"messageType"`
	Path        string `json:"path"`
}

type WSSendResponse struct {
	Flag string      `json:"flag"`
	Body interface{} `json:"body"`
}

type WSInput struct {
	Flag     string          `json:"flag"`
	ReadData json.RawMessage `json:"readData"`
}

type WSReadDataStruct struct {
	UserIds []uint64 `json:"userIds"`
	ChatId  uint64   `json:"chatId"`
}

type WSReadResponse struct {
	Flag string           `json:"flag"`
	Body WSReadDataStruct `json:"body"`
}
