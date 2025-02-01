package chat

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/merynayr/chat-server/internal/client/db"
)

func (r *repo) ChatExists(ctx context.Context, id int64) (bool, error) {
	query, args, err := sq.Select("1").
		From(tableNameChat).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: id}).
		Limit(1).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to build chat check query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.IsExistsByChat",
		QueryRaw: query,
	}

	var result int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	return true, nil
}
