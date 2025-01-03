package app

import (
	"log"

	"github.com/merynayr/chat-server/internal/config"
	"github.com/merynayr/chat-server/internal/config/env"
)

// Структура приложения со всеми зависимости
type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
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
