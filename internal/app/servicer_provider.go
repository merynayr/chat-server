package app

import (
	"context"
	"log"

	"github.com/merynayr/chat-server/internal/client/db"
	"github.com/merynayr/chat-server/internal/client/db/pg"
	"github.com/merynayr/chat-server/internal/client/db/transaction"
	"github.com/merynayr/chat-server/internal/closer"
	"github.com/merynayr/chat-server/internal/config"
	"github.com/merynayr/chat-server/internal/config/env"
	"github.com/merynayr/chat-server/internal/repository"
	"github.com/merynayr/chat-server/internal/service"

	chatAPI "github.com/merynayr/chat-server/internal/api/chat"
	chatRepository "github.com/merynayr/chat-server/internal/repository/chat"
	chatService "github.com/merynayr/chat-server/internal/service/chat"
)

// Структура приложения со всеми зависимости
type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	chatRepository repository.ChatRepository
	chatService    service.ChatService
	chatAPI        *chatAPI.API
}

// NewServiceProvider возвращает новый объект API слоя
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get gprc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatAPI(ctx context.Context) *chatAPI.API {
	if s.chatAPI == nil {
		s.chatAPI = chatAPI.NewAPI(s.ChatService(ctx))
	}

	return s.chatAPI
}
