package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/merynayr/chat-server/internal/api/chat"
	"github.com/merynayr/chat-server/internal/service"
	serviceMocks "github.com/merynayr/chat-server/internal/service/mocks"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID = gofakeit.Int64()

		serviceErr  = fmt.Errorf("service err")
		reqIsNilErr = fmt.Errorf("Request is nil")

		req = &desc.DeleteChatRequest{
			Id: chatID,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock.DeleteChatMock.Expect(ctx, chatID).Return(nil)
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
				mock.DeleteChatMock.Expect(ctx, chatID).Return(serviceErr)
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
			err:  reqIsNilErr,
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

			res, err := api.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
