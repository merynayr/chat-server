package model

import (
	"time"
)

// Chat модель чата в сервисном слое
type Chat struct {
	ChatID    int64
	ChatName  string
	Usernames []int64
	CreatedAt time.Time
}

// MessageInfo модель сообщения в сервисном слое
type MessageInfo struct {
	ID        int64
	ChatID    int64
	UserID    int64
	Text      string
	CreatedAt time.Time
}
