package model

import "time"

// Chat модель чата в репо слое
type Chat struct {
	ChatID    int64     `db:"chat_id"`
	ChatName  string    `db:"chat_name"`
	CreatedAt time.Time `db:"created_at"`
}

// Roster модель, отражающая связь чата и пользователя в репо слое
type Roster struct {
	ChatID int64  `db:"chat_id"`
	UserID string `db:"user_id"`
}

// Messages модель сообщений в репо слое
type Messages struct {
	MessageID int64     `db:"message_id"`
	ChatID    int64     `db:"chat_id"`
	UserID    string    `db:"user_id"`
	Contect   string    `db:"contect"`
	CreatedAt time.Time `db:"created_at"`
}
