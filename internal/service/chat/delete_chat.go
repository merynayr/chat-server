package chat

import (
	"context"
)

func (s *srv) DeleteChat(ctx context.Context, id int64) error {
	if err := s.chatRepo.DeleteChat(ctx, id); err != nil {
		return err
	}
	return nil
}
