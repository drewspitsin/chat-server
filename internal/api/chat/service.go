package chat

import (
	"github.com/drewspitsin/chat-server/internal/service"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
