package converter

import (
	"github.com/merynayr/chat-server/internal/model"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

// ToChatFromDesc конвертирует модель пользователя API слоя в
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
