package repository

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
)

func (repo Repository) CreateChat(ctx context.Context, initialChatData []app.ChatUser) (outputChatData app.CreateChatResponse, err error) {
	newChat := app.Chat{}
	err = repo.db.WithContext(ctx).Create(&newChat).Error
	if err != nil {
		return
	}
	chatId := newChat.Id
	outputChatData.ChatId = chatId

	for idx := range initialChatData {
		initialChatData[idx].ChatId = chatId
	}

	err = repo.db.WithContext(ctx).Create(&initialChatData).Error
	return
}

func (repo Repository) SendMessage(ctx context.Context, newMsg app.Message) (err error) {
	err = repo.db.WithContext(ctx).Create(&newMsg).Error
	return
}

func (repo Repository) GetChatsList(userData app.GetChatsListRequest) (recentMsgs []app.Message, err error) {
	subQuery := repo.db.Select("MAX(sent_at) as max_sent_at, chat_id").Table("messages").Group("chat_id")
	err = repo.db.Joins("INNER JOIN chat_users ON chat_users.user_id = ? AND chat_users.chat_id = messages.chat_id", userData.UserId).
		Joins("INNER JOIN (?) AS m ON messages.chat_id = m.chat_id AND messages.sent_at = m.max_sent_at", subQuery).
		Order("m.max_sent_at DESC").
		Find(&recentMsgs).Error
	return
}

func (repo Repository) GetChat(chatData app.GetChatRequest) (chatMsgs []app.Message, err error) {
	repo.db.Model(&app.Message{}).
		Where("chat_id = ? AND sender_id <> ? AND read_status = ?", chatData.ChatId, chatData.UserId, false).
		Update("read_status", true)

	err = repo.db.Where("chat_id = ?", chatData.ChatId).Order("sent_at DESC").Find(&chatMsgs).Error
	return
}
