package chat

import (
	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/repository"
)

// const (
// 	tableName = "chats"

// 	idColumn       = "id"
// 	chatNameColumn = "chat_name"
// )

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
