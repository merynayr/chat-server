package service

import (
	"context"

	"github.com/merynayr/chat-server/internal/model"
)

// ChatService интерфейс сервисного слоя chat
type ChatService interface {
	CreateChat(ctx context.Context, chat *model.Chat) (int64, error)
}
