package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/merynayr/chat-server/internal/client/db"
	txMocks "github.com/merynayr/chat-server/internal/client/db/mocks"
	"github.com/merynayr/chat-server/internal/model"
	"github.com/merynayr/chat-server/internal/repository"
	repositoryMocks "github.com/merynayr/chat-server/internal/repository/mocks"

	"github.com/merynayr/chat-server/internal/service/chat"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.MessageInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID    = gofakeit.Int64()
		userID    = gofakeit.Int64()
		text      = gofakeit.BeerBlg()
		timestamp = timestamppb.New(time.Now())

		repoErr   = fmt.Errorf("repository err")
		existsErr = fmt.Errorf("failed to send message: chat %d does not exist", chatID)

		req = &model.MessageInfo{
			ChatID:    chatID,
			UserID:    userID,
			Text:      text,
			CreatedAt: timestamp.AsTime(),
		}
	)

	tests := []struct {
		name               string
		args               args
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
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.ChatExistsMock.Expect(ctx, req.ChatID).Return(true, nil)
				mock.CreateMessageMock.Expect(ctx, req).Return(nil)
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
			name: "error in CheckExists function",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.ChatExistsMock.Expect(ctx, req.ChatID).Return(false, repoErr)
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
			name: "error in CreateMessage function",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.ChatExistsMock.Expect(ctx, req.ChatID).Return(true, nil)
				mock.CreateMessageMock.Expect(ctx, req).Return(repoErr)
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
			name: "chatID does not exists: error in CheckExists function",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: existsErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.ChatExistsMock.Expect(ctx, req.ChatID).Return(false, nil)
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
			err := service.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
