package redis

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/EgorTarasov/lct-2024/api/internal/users/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/redis"
)

type tokenRepo struct {
	r      *redis.Redis[models.UserDao]
	tracer trace.Tracer
}

// New создание репозитория для токенов авторизации.
func New(_ context.Context, redisClient *redis.Redis[models.UserDao], tracer trace.Tracer) *tokenRepo {
	return &tokenRepo{
		r:      redisClient,
		tracer: tracer,
	}
}

// Set сохраняет данные пользователя с токеном.
func (tr *tokenRepo) Set(ctx context.Context, token string, data models.UserDao) error {
	ctx, span := tr.tracer.Start(ctx, "webAuthSessionRepo.Set")
	defer span.End()
	return tr.r.Set(ctx, token, data)
}

// Get получения данных по токену.
func (tr *tokenRepo) Get(ctx context.Context, token string) (models.UserDao, error) {
	ctx, span := tr.tracer.Start(ctx, "webAuthSessionRepo.Get")
	defer span.End()
	return tr.r.Get(ctx, token)
}
