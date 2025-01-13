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
	Chat      *Chat
	User      *User
	Text      string
	CreatedAt time.Time
}

// User модель пользователя в сервисном слое
type User struct {
	ID   int64
	Name string
}
