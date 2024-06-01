package service

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/EgorTarasov/lct-2024/api/internal/auth/models"
	"github.com/EgorTarasov/lct-2024/api/internal/config"
)

type tokenRepo interface {
	Set(ctx context.Context, token string, data models.UserDao) error
	Get(ctx context.Context, token string) (models.UserDao, error)
}

type userRepo interface {
	vkUserRepo
	emailUserRepo
}

type service struct {
	tracer trace.Tracer
	cfg    *config.Config
	ur     userRepo
	tr     tokenRepo
}

// New конструктор сервиса для работы с аккаунтами пользователей
func New(_ context.Context, cfg *config.Config, userRepo userRepo, tokenRepo tokenRepo, tracer trace.Tracer) *service {
	return &service{
		cfg:    cfg,
		tracer: tracer,
		ur:     userRepo,
		tr:     tokenRepo,
	}
}
