package chat

import (
	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/repository"
)

const (
	tableNameChat     = "chat"
	tableNameMessages = "messages"
	tableNameRoster   = "roster"

	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	messageIDColumn = "message_id"

	chatNameColumn  = "chat_name"
	contectColumn   = "contect"
	createdAtColumn = "created_at"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}
