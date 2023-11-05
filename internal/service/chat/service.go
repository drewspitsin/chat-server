package chat

import (
	"github.com/drewspitsin/chat-server/internal/client/db"
	"github.com/drewspitsin/chat-server/internal/repository"
	"github.com/drewspitsin/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) service.ChatService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.ChatRepository:
			srv.chatRepository = s
		}
	}

	return &srv
}
