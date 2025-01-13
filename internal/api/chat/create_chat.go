package chat

import (
	"context"

	"github.com/merynayr/chat-server/internal/converter"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

// CreateChat - отправляет запрос в сервисный слой на создание пользователя
func (a *API) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatID, err := a.chatService.CreateChat(ctx, converter.ToChatFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{
		Id: chatID,
	}, nil
}
