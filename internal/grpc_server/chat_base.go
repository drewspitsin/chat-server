package grpc_server

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

const (
	table    = "chat_server"
	id       = "id"
	username = "username"
	msg      = "msg"
)

func (s *ChatV1Server) CreatePg(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	builderInsert := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(username, msg).
		Values(req.Usernames, req.Msg).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatServerID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatServerID)
	if err != nil {
		log.Fatalf("failed to insert chat_server: %v", err)
	}

	log.Printf("inserted chat_server with id: %d", chatServerID)

	return &desc.CreateResponse{
		Id: chatServerID,
	}, nil
}

func (s *ChatV1Server) DeletePg(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	builderInsert := sq.Delete(table).
		Where(sq.Eq{id: req.GetId()}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
		return nil, err
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update note: %v tag: %v", err, res)
		return nil, err
	}

	return nil, nil
}

func (s *ChatV1Server) SendMessagePg(ctx context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	// Принцип работы
	return nil, nil
}
