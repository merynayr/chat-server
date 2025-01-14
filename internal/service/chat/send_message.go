package chat

import (
	"context"

	"github.com/merynayr/chat-server/internal/model"
)

func (s *srv) SendMessage(ctx context.Context, message *model.MessageInfo) error {
	if err := s.chatRepo.CreateMessage(ctx, message); err != nil {
		return err
	}

	return nil
}
