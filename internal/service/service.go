package service

import (
	"context"

	"github.com/drewspitsin/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, info *model.Chat) (int64, error)
	Delete(ctx context.Context, info *model.Chat) error
	Send(ctx context.Context, info *model.Chat) error
}
