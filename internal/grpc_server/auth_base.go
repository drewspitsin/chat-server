package grpc_server

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

// id serial primary key,
// UserName text not null,
// Email text not null,
// Pswd text not null,
// created_at timestamp not null default now(),
// updated_at timestamp

func (s *ChatV1Server) CreatePg(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	builderInsert := sq.Insert("chat_server").
		PlaceholderFormat(sq.Dollar).
		Columns("msg").
		Values(gofakeit.City()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chat_serverID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chat_serverID)
	if err != nil {
		log.Fatalf("failed to insert chat_server: %v", err)
	}

	log.Printf("inserted chat_server with id: %d", chat_serverID)

	return &desc.CreateResponse{
		Id: chat_serverID,
	}, nil
}

func (s *ChatV1Server) DeletePg(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *ChatV1Server) SendMessagePg(ctx context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	return nil, nil
}
