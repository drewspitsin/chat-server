package chat

import (
	"context"

	"github.com/drewspitsin/chat-server/internal/converter"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Send(ctx context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	err := i.chatService.Send(ctx, converter.ToChatFromDescSend(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
