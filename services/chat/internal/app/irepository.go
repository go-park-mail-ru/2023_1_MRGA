package app

import (
	"context"
)

type IRepository interface {
	SendMessage(context.Context, Message) error
	GetRecentMessages(uint) ([]Message, error)
}
