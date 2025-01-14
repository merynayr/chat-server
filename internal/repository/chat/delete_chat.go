package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/merynayr/chat-server/internal/client/db"
)

func (r *repo) DeleteChat(ctx context.Context, id int64) error {
	query, args, err := sq.Delete(tableNameChat).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build chat delete query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.DeleteChat",
		QueryRaw: query,
	}

	tag, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("failed to delete chat: %d not found", id)
	}
	return nil
}
