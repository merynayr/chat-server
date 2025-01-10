package chat

import (
	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/repository"
	"github.com/merynayr/chat-server/internal/service"
)

// Структура сервисного слоя с объектами репо слоя
// и транзакционного менеджера
type srv struct {
	chatRepo  repository.ChatRepository
	txManager db.TxManager
}

// NewService возвращает объект сервисного слоя
func NewService(
	chatRepo repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &srv{
		chatRepo:  chatRepo,
		txManager: txManager,
	}
}
