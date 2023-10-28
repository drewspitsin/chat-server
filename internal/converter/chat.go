package converter

import (
	"github.com/drewspitsin/chat-server/internal/model"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
)

func ToChatFromDescCreate(chat *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		ID:   0,
		Name: chat.Username,
		Msg:  chat.Msg,
	}
}

func ToChatFromDescDelete(chat *desc.DeleteRequest) int64 {
	return chat.Id

}

func ToChatFromDescSend(chat *desc.SendRequest) *model.Chat {
	return &model.Chat{
		ID: 0,
	}
}
