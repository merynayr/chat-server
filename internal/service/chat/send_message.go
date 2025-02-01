package chat

import (
	"context"
	"fmt"

	"github.com/merynayr/chat-server/internal/model"
)

func (s *srv) SendMessage(ctx context.Context, message *model.MessageInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		var exist bool

		exist, errTx = s.chatRepo.ChatExists(ctx, message.ChatID)
		if !exist && errTx == nil {
			return fmt.Errorf("failed to send message: chat %d does not exist", message.ChatID)
		}
		if errTx != nil {
			return errTx
		}

		errTx = s.chatRepo.CreateMessage(ctx, message)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
