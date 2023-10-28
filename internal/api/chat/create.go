package chat

import (
	"context"
	"log"

	"github.com/drewspitsin/chat-server/internal/converter"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatFromDescCreate(req))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted auth with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
