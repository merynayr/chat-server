package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/model"
)

func (r *repo) CreateMessage(ctx context.Context, message *model.MessageInfo) error {
	query, args, err := sq.Insert(tableNameMessages).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn, contectColumn, createdAtColumn).
		Values(message.ChatID, message.UserID, message.Text, message.CreatedAt).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build message insertion query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.CreateMessage",
		QueryRaw: query,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}
	return nil
}
