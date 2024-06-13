package service

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/config"
	"go.opentelemetry.io/otel/trace"
)

type service struct {
	tracer trace.Tracer
	cfg    *config.Config
	or     objectRepo
	moek   moekRepo
	ods    odsRepository
}

// New конструктор сервиса для работы с аккаунтами пользователей.
func New(_ context.Context, cfg *config.Config, objectRepo objectRepo, moek moekRepo, ods odsRepository, tracer trace.Tracer) *service {
	return &service{
		cfg:    cfg,
		tracer: tracer,
		or:     objectRepo,
		moek:   moek,
		ods:    ods,
	}
}
