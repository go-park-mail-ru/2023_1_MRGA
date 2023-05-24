package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
)

func (repo Repository) GetDialogIfExists(ctx context.Context, initialDialogData []app.ChatUser) (outputChatData app.CreateChatResponse, found bool, err error) {
	if len(initialDialogData) != 2 {
		err = errors.New("Для инициализации диалога отправлены не два учатсника")
		return
	}

	err = repo.db.WithContext(ctx).
		Table("chat_users").
		Select("chat_id").
		Where("user_id = ? OR user_id = ?", initialDialogData[0].UserId, initialDialogData[1].UserId).
		Group("chat_id").
		Having("COUNT(user_id) = 2").
		First(&outputChatData).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		return
	}

	if err != nil {
		return
	}

	found = true
	return
}

func (repo Repository) CreateChat(ctx context.Context, initialChatData []app.ChatUser) (outputChatData app.CreateChatResponse, err error) {
	var newChat app.ChatUser
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

func (repo Repository) SendMessage(ctx context.Context, newMsg app.ChatMessage) (msgId uint, err error) {

	err = repo.db.WithContext(ctx).Create(&newMsg.Message).Error
	if err != nil {
		return
	}

	msgId = newMsg.Message.Id

	if newMsg.MessageType != app.TextMessage {
		media := app.Media{
			MessageId:   msgId,
			MessageType: string(newMsg.MessageType),
			Path:        newMsg.Path,
		}
		err = repo.db.WithContext(ctx).Create(&media).Error
		if err != nil {
			return
		}
	}
	return
}

func (repo Repository) GetChatsList(userData app.GetChatsListRequest) (recentMsgs []app.MessageWithChatUsers, err error) {
	var msgs []app.ChatMessage
	subQuery := repo.db.Select("MAX(sent_at) as max_sent_at, chat_id").Table("messages").Group("chat_id")
	err = repo.db.Table("messages").
		Joins("INNER JOIN chat_users ON chat_users.user_id = ? AND chat_users.chat_id = messages.chat_id", userData.UserId).
		Joins("INNER JOIN (?) AS m ON messages.chat_id = m.chat_id AND messages.sent_at = m.max_sent_at", subQuery).
		Joins("LEFT JOIN media ON messages.id = media.message_id").
		Select("messages.*, media.message_type, media.path").
		Order("m.max_sent_at DESC").
		Find(&msgs).Error
	if err != nil {
		return
	}

	for _, recentMsg := range msgs {
		chatId := recentMsg.ChatId
		var chatUsers []uint
		err = repo.db.Table("chat_users").
			Select("user_id").
			Where("chat_id = ? AND user_id <> ?", chatId, userData.UserId).
			Find(&chatUsers).Error
		if err != nil {
			return
		}

		recentMsgs = append(recentMsgs, app.MessageWithChatUsers{
			ChatMessage: recentMsg,
			ChatUserIds: chatUsers,
		})
	}

	return
}

func (repo Repository) GetChat(chatData app.GetChatRequest) (chatMsgs []app.ChatMessage, err error) {
	repo.db.Model(&app.Message{}).
		Where("chat_id = ? AND sender_id <> ? AND read_status = ?", chatData.ChatId, chatData.UserId, false).
		Update("read_status", true)

	err = repo.db.Model(&app.Message{}).
		Joins("LEFT JOIN media ON messages.id = media.message_id").
		Select("messages.*, media.message_type, media.path").
		Where("chat_id = ?", chatData.ChatId).
		Order("sent_at ASC, id DESC").
		Find(&chatMsgs).Error
	return
}

func (repo Repository) GetChatParticipants(ctx context.Context, initialChatData app.GetChatParticipantsRequest) (participants app.GetChatParticipantsResponse, err error) {
	err = repo.db.WithContext(ctx).Table("chat_users").
		Select("user_id").
		Where("chat_id = ? AND user_id <> ?", initialChatData.ChatId, initialChatData.UserId).
		Find(&participants.ChatUserIds).Error
	if err != nil {
		return
	}

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
