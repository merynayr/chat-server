package chat

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/model"
)

func (r *repo) CreateChat(ctx context.Context, chat *model.Chat) (int64, error) {
	query, args, err := sq.Insert(tableNameChat).
		PlaceholderFormat(sq.Dollar).
		Columns(chatNameColumn, createdAtColumn).
		Values(chat.ChatName, chat.CreatedAt).
		Suffix("RETURNING chat_id").
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to build chat insertion query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repository.CreateChat",
		QueryRaw: query,
	}

	var chatID int64
	err = r.db.DB().ScanOneContext(ctx, &chatID, q, args...)
	if err != nil {
		log.Printf("failed to insert chat: %s", err)
		return 0, fmt.Errorf("failed to insert chat: %w", err)
	}
	return chatID, nil
}
