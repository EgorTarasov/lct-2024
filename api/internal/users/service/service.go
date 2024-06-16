package service

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/pkg/kafka"
	pkgs3 "github.com/EgorTarasov/lct-2024/api/pkg/s3"
	"go.opentelemetry.io/otel/trace"

	"github.com/EgorTarasov/lct-2024/api/internal/config"
	"github.com/EgorTarasov/lct-2024/api/internal/users/models"
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
	tracer   trace.Tracer
	cfg      *config.Config
	s3       *pkgs3.S3
	ur       userRepo
	tr       tokenRepo
	data     fileRepo
	producer *kafka.Producer
	topic    string
}

// New конструктор сервиса для работы с аккаунтами пользователей.
func New(_ context.Context, cfg *config.Config, userRepo userRepo, data fileRepo, tokenRepo tokenRepo, s3 *pkgs3.S3, producer *kafka.Producer, topic string, tracer trace.Tracer) *service {
	return &service{
		cfg:      cfg,
		tracer:   tracer,
		s3:       s3,
		ur:       userRepo,
		tr:       tokenRepo,
		data:     data,
		producer: producer,
		topic:    topic,
	}
}
