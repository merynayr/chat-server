package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/merynayr/chat-server/internal/api/chat"
	"github.com/merynayr/chat-server/internal/model"
	"github.com/merynayr/chat-server/internal/service"
	serviceMocks "github.com/merynayr/chat-server/internal/service/mocks"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID    = gofakeit.Int64()
		userID    = gofakeit.Int64()
		text      = gofakeit.BeerBlg()
		timestamp = timestamppb.New(time.Now())

		serviceErr = fmt.Errorf("service err")

		req = &desc.SendMessageRequest{
			ChatId:    wrapperspb.Int64(chatID),
			UserId:    wrapperspb.Int64(userID),
			Text:      wrapperspb.String(text),
			Timestamp: timestamp,
		}

		info = &model.MessageInfo{
			ChatID:    chatID,
			UserID:    userID,
			Text:      text,
			CreatedAt: timestamp.AsTime(),
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
				mock.SendMessageMock.Expect(ctx, info).Return(nil)
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
				mock.SendMessageMock.Expect(ctx, info).Return(serviceErr)
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
			err:  errors.New("failed to send message: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
		{
			name: "empty text of message",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:    wrapperspb.Int64(chatID),
					UserId:    wrapperspb.Int64(userID),
					Timestamp: timestamp,
				},
			},
			want: nil,
			err:  errors.New("failed to send message: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
		{
			name: "empty chatID",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					UserId:    wrapperspb.Int64(userID),
					Text:      wrapperspb.String(text),
					Timestamp: timestamp,
				},
			},
			want: nil,
			err:  errors.New("failed to send message: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
		{
			name: "empty userID",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId:    wrapperspb.Int64(chatID),
					Text:      wrapperspb.String(text),
					Timestamp: timestamp,
				},
			},
			want: nil,
			err:  errors.New("failed to send message: Request id bad"),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
		{
			name: "empty Time of create message",
			args: args{
				ctx: ctx,
				req: &desc.SendMessageRequest{
					ChatId: wrapperspb.Int64(chatID),
					UserId: wrapperspb.Int64(userID),
					Text:   wrapperspb.String(text),
				},
			},
			want: nil,
			err:  errors.New("failed to send message: Request id bad"),
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

			res, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
