package converter

import (
	"github.com/drewspitsin/chat-server/internal/model"
	modelRepo "github.com/drewspitsin/chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:   chat.ID,
		Name: chat.Name,
		Msg:  chat.Msg,
	}
}
