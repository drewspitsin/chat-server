package chat

import (
	"context"

	"github.com/drewspitsin/chat-server/internal/converter"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := i.chatService.Delete(ctx, converter.ToChatFromDescDelete(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
