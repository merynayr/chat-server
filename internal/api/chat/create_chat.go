package chat

import (
	"context"
	"fmt"

	"github.com/merynayr/chat-server/internal/converter"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

// CreateChat - отправляет запрос в сервисный слой на создание пользователя
func (a *API) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	convertedChat := converter.ToChatFromDesc(req)
	if convertedChat == nil {
		return nil, fmt.Errorf("failed to create chat: Request id bad")
	}

	chatID, err := a.chatService.CreateChat(ctx, convertedChat)
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{
		Id: chatID,
	}, nil
}
