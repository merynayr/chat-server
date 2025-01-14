package converter

import (
	"github.com/merynayr/chat-server/internal/model"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

// ToChatFromDesc конвертирует модель чата API слоя в
// модель сервисного слоя
func ToChatFromDesc(chat *desc.CreateChatRequest) *model.Chat {
	if chat == nil {
		return nil
	}

	return &model.Chat{
		ChatName:  chat.ChatName,
		Usernames: chat.UsersId,
	}
}

// ToMessageFromDesc конвертирует модель сообщения API слоя в
// модель сервисного слоя
func ToMessageFromDesc(message *desc.SendMessageRequest) *model.MessageInfo {
	if message == nil {
		return nil
	}

	m := &model.MessageInfo{}

	if message.ChatId != nil {
		m.ChatID = message.ChatId.GetValue()
	} else {
		return nil
	}

	if message.UserId != nil {
		m.UserID = message.UserId.GetValue()
	} else {
		return nil
	}

	if message.Timestamp != nil {
		m.CreatedAt = message.Timestamp.AsTime()
	} else {
		return nil
	}

	if message.Text != nil {
		m.Text = message.Text.GetValue()
	} else {
		return nil
	}

	return m
}
