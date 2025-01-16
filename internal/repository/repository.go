package repository

import (
	"context"

	"github.com/merynayr/chat-server/internal/model"
)

// ChatRepository - интерфейс репо слоя chat
type ChatRepository interface {
	CreateChat(ctx context.Context, chat *model.Chat) (int64, error)
	CreateRoster(ctx context.Context, chatID int64, UserIDs []int64) error
	DeleteChat(ctx context.Context, id int64) error
	CreateMessage(ctx context.Context, message *model.MessageInfo) error
	ChatExists(ctx context.Context, id int64) (bool, error)
}
