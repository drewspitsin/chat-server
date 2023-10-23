package chat

import (
	"context"

	"github.com/drewspitsin/chat-server/internal/model"
)

func (s *serv) Delete(ctx context.Context, info *model.Chat) error {
	err := s.chatRepository.Delete(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
