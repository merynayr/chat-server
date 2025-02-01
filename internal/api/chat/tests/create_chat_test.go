package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/merynayr/chat-server/internal/api/chat"
	"github.com/merynayr/chat-server/internal/model"
	"github.com/merynayr/chat-server/internal/service"
	serviceMocks "github.com/merynayr/chat-server/internal/service/mocks"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID   = gofakeit.Int64()
		chatName = gofakeit.Name()
		userIDs  = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}

		serviceErr = fmt.Errorf("service err")

		req = &desc.CreateChatRequest{
			ChatName: chatName,
			UsersId:  userIDs,
		}

		info = &model.Chat{
			ChatName:  chatName,
			Usernames: userIDs,
		}

		res = &desc.CreateChatResponse{
			Id: chatID,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateChatResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success test",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, info).Return(chatID, nil)
				return mock
			},
		},
		{
			name: "error from service",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, info).Return(int64(0), serviceErr)
				return mock
			},
		},
		{
			name: "Request is nil",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: nil,
			err:  errors.New("failed to create chat: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
		{
			name: "empty users list",
			args: args{
				ctx: ctx,
				req: &desc.CreateChatRequest{
					ChatName: chatName,
					UsersId:  []int64{},
				},
			},
			want: nil,
			err:  errors.New("failed to create chat: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewAPI(chatServiceMock)

			res, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
