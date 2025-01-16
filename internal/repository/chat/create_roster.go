package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/chat-server/internal/client/db"
)

func (r *repo) CreateRoster(ctx context.Context, chatID int64, UserIDs []int64) error {
	for _, userID := range UserIDs {
		query, args, err := sq.Insert(tableNameRoster).
			PlaceholderFormat(sq.Dollar).
			Columns(chatIDColumn, userIDColumn).
			Values(chatID, userID).
			ToSql()

		if err != nil {
			return fmt.Errorf("failed to build roster insertion query: %w", err)
		}

		q := db.Query{
			Name:     "chat_repository.CreateRoster",
			QueryRaw: query,
		}

		_, err = r.db.DB().ExecContext(ctx, q, args...)
		if err != nil {
			return fmt.Errorf("failed to insert roster: %w", err)
		}
	}

	return nil
}
