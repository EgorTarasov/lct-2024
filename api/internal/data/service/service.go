package service

import (
	"context"
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/data/models"
	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	pb "github.com/EgorTarasov/lct-2024/api/internal/stubs"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// addressRegistry справочник с гео информацией.
type addressRegistry interface {
	GetByAdmArea(ctx context.Context, admArea string) ([]shared.Address, error)
}

type eventRepo interface {
	GetEmergencyEvents(ctx context.Context, unoms []int64) (events []models.Event, err error)
}

type service struct {
	ar     addressRegistry
	ev     eventRepo
	client pb.InferenceClient
	tracer trace.Tracer
}

// NewService конструктор сервиса для работы с картой.
func NewService(ar addressRegistry, ev eventRepo, client pb.InferenceClient, tracer trace.Tracer) *service {
	return &service{
		ar:     ar,
		ev:     ev,
		client: client,
		tracer: tracer,
	}
}

// GetEmergencyPredictions получение предсказаний аварийных ситуаций.
func (s *service) GetEmergencyPredictions(ctx context.Context, admArea string, startDate, endDate time.Time, threshold float32) ([]models.PredictionResult, error) {
	ctx, span := s.tracer.Start(ctx, "data.GetEmergencyPredictions",
		trace.WithAttributes(
			attribute.String("start", startDate.String()),
			attribute.String("end", endDate.String()),
			attribute.Float64("threshold", float64(threshold)),
		),
	)
	defer span.End()

	addresses, err := s.ar.GetByAdmArea(ctx, admArea)
	if err != nil {
		return nil, err
	}
	unoms := make([]int64, len(addresses))
	for idx, value := range addresses {
		unoms[idx] = value.Unom
	}
	res, err := s.client.Inference(ctx, &pb.Query{
		Unoms:     unoms,
		StartDate: timestamppb.New(startDate),
		EndDate:   timestamppb.New(endDate),
		Threshold: threshold,
	})
	if err != nil {
		return nil, err
	}
	results := make([]models.PredictionResult, len(res.Predictions))
	for idx, record := range res.Predictions {
		results[idx] = models.PredictionResult{
			Unom:          record.Unom,
			Date:          record.Date.AsTime(),
			P1:            record.P1,
			P2:            record.P2,
			T1:            record.T1,
			T2:            record.T2,
			No:            record.No,
			NoHeating:     record.NoHeating,
			Leak:          record.Leak,
			StrongLeak:    record.StrongLeak,
			TempLow:       record.TempLow,
			TempLowCommon: record.TempLowCommon,
			LeakSystem:    record.LeakSystem,
		}
	}
	return results, nil
}
