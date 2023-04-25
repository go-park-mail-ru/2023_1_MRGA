package repository

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/app"
)

func (repo Repository) SendMessage(ctx context.Context, msg app.Message) error {
	return repo.db.WithContext(ctx).Create(&msg).Error
}

func (repo Repository) GetRecentMessages(userId uint) (recentMessages []app.Message, err error) {
	err = repo.db.Table("messages").
		Where("sender_id = ? OR receiver_id = ?", userId, userId).
		Select("DISTINCT ON (LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id)) *").
		Order("LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id), sent_at DESC").
		Find(&recentMessages).Error
	return
}
