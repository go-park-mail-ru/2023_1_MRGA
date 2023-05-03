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

func (repo Repository) GetChatsList(userData app.GetChatsListRequest) (recentMsgs []app.MessageWithChatUsers, err error) {
	var msgs []app.Message
	subQuery := repo.db.Select("MAX(sent_at) as max_sent_at, chat_id").Table("messages").Group("chat_id")
	err = repo.db.
		Joins("INNER JOIN chat_users ON chat_users.user_id = ? AND chat_users.chat_id = messages.chat_id", userData.UserId).
		Joins("INNER JOIN (?) AS m ON messages.chat_id = m.chat_id AND messages.sent_at = m.max_sent_at", subQuery).
		Select("messages.*").
		Order("m.max_sent_at DESC").
		Find(&msgs).Error
	if err != nil {
		return
	}

	for _, recentMsg := range msgs {
		chatId := recentMsg.ChatId
		var chatUsers []uint
		repo.db.Table("chat_users").
			Select("user_id").
			Where("chat_id = ? AND user_id <> ?", chatId, userData.UserId).
			Find(&chatUsers)

		recentMsgs = append(recentMsgs, app.MessageWithChatUsers{
			Message:     recentMsg,
			ChatUserIds: chatUsers,
		})
	}

	return
}

func (repo Repository) GetChat(chatData app.GetChatRequest) (chatMsgs []app.Message, err error) {
	repo.db.Model(&app.Message{}).
		Where("chat_id = ? AND sender_id <> ? AND read_status = ?", chatData.ChatId, chatData.UserId, false).
		Update("read_status", true)

	err = repo.db.Where("chat_id = ?", chatData.ChatId).Order("sent_at ASC, id DESC").Find(&chatMsgs).Error
	return
}

func (repo Repository) IsMemberOfChat(userId uint, chatId uint) (result bool, err error) {
	var count int64

	err = repo.db.Model(&app.ChatUser{}).
		Where("user_id = ? AND chat_id = ?", userId, chatId).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return uint(count) > 0, nil
}
