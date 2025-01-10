package chat

import (
	"github.com/merynayr/chat-server/internal/service"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

// API user структура с заглушками gRPC-методов (при их отсутствии) и
// объект сервисного слоя (его интерфейса)
type API struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewAPI возвращает новый объект имплементации API-слоя
func NewAPI(chatService service.ChatService) *API {
	return &API{
		chatService: chatService,
	}
}
