package grpc_server

import (
	"context"

	desc "github.com/drewspitsin/chat-server/pkg/chat_api_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ChatV1Server
type ChatV1Server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

// NewChatV1Server returns a new ChatV1Server instance
func NewChatV1Server(p *pgxpool.Pool) *ChatV1Server {
	return &ChatV1Server{
		UnimplementedChatV1Server: desc.UnimplementedChatV1Server{},
		pool:                      p,
	}
}

// Create is a method that implements the Create method of the ChatV1Server
func (s *ChatV1Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{}, nil
}

// Delete is a method that implements the Delete method of the ChatV1Server
func (s *ChatV1Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// SendMessage is a method that implements the SendMessage method of the ChatV1Server
func (s *ChatV1Server) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
