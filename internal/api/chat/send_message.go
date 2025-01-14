package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/merynayr/chat-server/internal/converter"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage Метод отправки сообщения в чат.
func (a *API) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	convertedChat := converter.ToMessageFromDesc(req)
	if convertedChat == nil {
		return nil, fmt.Errorf("failed to send message: Request id bad")
	}
	log.Println(convertedChat)
	if err := a.chatService.SendMessage(ctx, convertedChat); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
