package chat

import (
	"context"

	"github.com/drewspitsin/chat-server/internal/model"
)

func (s *serv) Send(ctx context.Context, info *model.Chat) error {
	err := s.chatRepository.Send(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
