package chat

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/drewspitsin/chat-server/internal/client/db"
	"github.com/drewspitsin/chat-server/internal/model"
	"github.com/drewspitsin/chat-server/internal/repository"
)

const (
	table    = "chat_server"
	id       = "id"
	username = "username"
	msg      = "msg"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.ChatRepository {
	return &repo{db: dbClient}
}

func (s *repo) Create(ctx context.Context, info *model.Chat) (int64, error) {
	builderInsert := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(username, msg).
		Values(info.Name, info.Msg).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatServerID int64
	err = s.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatServerID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert chat_server: %v", err)
	}

	log.Printf("inserted chat_server with id: %d", chatServerID)

	return chatServerID, nil
}

func (s *repo) Delete(ctx context.Context, id_d int64) error {
	builderInsert := sq.Delete(table).
		Where(sq.Eq{id: id_d}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %v", err)
	}
	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	res, err := s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update note: %v tag: %v", err, res)
	}

	return nil
}

func (s *repo) Send(ctx context.Context, info *model.Chat) error {
	return nil
}
