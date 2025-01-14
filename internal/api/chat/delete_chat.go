package chat

import (
	"context"

	desc "github.com/merynayr/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat - отправляет запрос в сервисный слой на создание пользователя
func (a *API) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	err := a.chatService.DeleteChat(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
