package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/merynayr/chat-server/internal/client/db"
	txMocks "github.com/merynayr/chat-server/internal/client/db/mocks"
	"github.com/merynayr/chat-server/internal/model"
	"github.com/merynayr/chat-server/internal/repository"
	repositoryMocks "github.com/merynayr/chat-server/internal/repository/mocks"

	"github.com/merynayr/chat-server/internal/service/chat"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID   = gofakeit.Int64()
		chatName = gofakeit.Name()
		userIDs  = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}

		repoErr   = fmt.Errorf("repository err")
		rosterErr = fmt.Errorf("roster err")

		req = &model.Chat{
			ChatName:  chatName,
			Usernames: userIDs,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success test",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: chatID,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(chatID, nil)
				mock.CreateRosterMock.Expect(ctx, chatID, req.Usernames).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "error in CreateChat function",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "error in CreateRoster function",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  rosterErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(chatID, nil)
				mock.CreateRosterMock.Expect(ctx, chatID, req.Usernames).Return(rosterErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatRepositoryMock := tt.chatRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := chat.NewService(chatRepositoryMock, txManagerMock)
			res, err := service.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
