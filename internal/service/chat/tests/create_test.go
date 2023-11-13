package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/drewspitsin/chat-server/internal/client/db"
	txMocks "github.com/drewspitsin/chat-server/internal/client/db/mocks"
	"github.com/drewspitsin/chat-server/internal/client/db/transaction"
	"github.com/drewspitsin/chat-server/internal/model"
	"github.com/drewspitsin/chat-server/internal/repository"
	repoMocks "github.com/drewspitsin/chat-server/internal/repository/mocks"
	"github.com/drewspitsin/chat-server/internal/service/chat"
)

type TxMock struct {
	pgxpool.Tx
}

func (t *TxMock) Commit(_ context.Context) error {
	return nil
}

func (t *TxMock) Rollback(_ context.Context) error {
	return nil
}

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txTransactorMockFunc func(mc *minimock.Controller) db.Transactor

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Animal()
		msg  = gofakeit.Animal()

		repoErr = fmt.Errorf("repo error")

		txM TxMock

		req = &model.Chat{
			ID:   id,
			Name: []string{name, name},
			Msg:  msg,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Return(id, nil)
				return mock
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Return(&txM, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				return mock
			},

			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Return(nil, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			authRepoMock := tt.chatRepositoryMock(mc)
			txTransact := transaction.NewTransactionManager(tt.txTransactorMock(mc))
			service := chat.NewService(authRepoMock, txTransact)
			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.want, newID)

		})
	}
}
