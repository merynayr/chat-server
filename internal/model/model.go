package model

import "database/sql"

// Chat модель чата в сервисном слое
type Chat struct {
	ChatID    int64
	Usernames []string
}

// MessageInfo модель сообщения в сервисном слое
type MessageInfo struct {
	From      string
	Text      string
	Timestamp sql.NullTime
}
