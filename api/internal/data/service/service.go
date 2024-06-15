package service

import (
	"context"
	"math"
	"time"

	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	pb "github.com/EgorTarasov/lct-2024/api/internal/stubs"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// addressRegistry справочник с гео информацией.
type addressRegistry interface {
	GetByAdmArea(ctx context.Context, admArea string) ([]shared.Address, error)
	GetGeoDataByUnom(ctx context.Context, unom int64) (result shared.Address, err error)
	GetAllObjectsByUnom(ctx context.Context, unom int64) (shared.UnomResult, error)
}

type eventRepo interface {
	GetEmergencyEvents(ctx context.Context, unoms []int64) (events []shared.Event, err error)
	GetMeasurementsByUnom(ctx context.Context, unom int64, limit int64) (measurements []shared.Measurement, err error)
}

type incidentRepo interface {
	GetByID(ctx context.Context, id int64) (shared.Incident, error)
	Create(ctx context.Context, title, status string, priority int, unom int64) (int64, error)
	GetRecent(ctx context.Context, limit, offset int) ([]shared.Incident, error)
}

type service struct {
	ar     addressRegistry
	ev     eventRepo
	ir     incidentRepo
	client pb.InferenceClient
	tracer trace.Tracer
}

// NewService конструктор сервиса для работы с картой.
func NewService(ar addressRegistry, ev eventRepo, ir incidentRepo, client pb.InferenceClient, tracer trace.Tracer) *service {
	return &service{
		ar:     ar,
		ev:     ev,
		ir:     ir,
		client: client,
		tracer: tracer,
	}
}

// GetEmergencyPredictions получение предсказаний аварийных ситуаций.
func (s *service) GetEmergencyPredictions(ctx context.Context, admArea string, startDate, endDate time.Time, threshold float32) ([]shared.PredictionResult, error) {
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
	results := make([]shared.PredictionResult, len(res.Predictions))
	for idx, record := range res.Predictions {
		results[idx] = shared.PredictionResult{
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

// GetRecentIncidents получение списка недавних инцидентов.
func (s *service) GetRecentIncidents(ctx context.Context, limit, offset int) ([]shared.Incident, error) {
	ctx, span := s.tracer.Start(
		ctx,
		"data.GetRecentIncidents",
		trace.WithAttributes(
			attribute.Int("limit", limit),
			attribute.Int("offset", offset),
		),
	)
	defer span.End()

	recent, err := s.ir.GetRecent(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return recent, nil
}

func createGraph(hours int) shared.Graph {
	const a = 74
	f := func(x int) float64 {
		// y\ =\ \frac{a}{\sqrt{\left(x+25\right)}}\cdot3\ -\ 19
		return (a/math.Sqrt(float64(x)+25))*3 - 19
	}

	points := make([]shared.GraphDataPoint, hours)
	for i := 0; i < hours; i++ {
		points[i] = shared.GraphDataPoint{
			Temp:       f(i + 1),
			TimeString: i,
		}
	}
	return shared.Graph{
		Name:   "Температура",
		Points: points,
	}
}

// GetIncidentByID получение дополнительной информации об инциденте с данными о потребителе и производителя энергии.
func (s *service) GetIncidentByID(ctx context.Context, id int64) (shared.Incident, error) {
	ctx, span := s.tracer.Start(
		ctx,
		"data.GetIncidentByID",
		trace.WithAttributes(
			attribute.Int64("id", id),
		),
	)
	defer span.End()

	result, err := s.ir.GetByID(ctx, id)
	if err != nil {
		return shared.Incident{}, err
	}
	// TODO: add mongo requests for buildings data

	result.HeatingGraph = createGraph(20)

	measurements, err := s.ev.GetMeasurementsByUnom(ctx, result.Unom, 10)
	if err != nil {
		return shared.Incident{}, err
	}
	result.Measurements = measurements

	result.RelatedObjects, err = s.ar.GetAllObjectsByUnom(ctx, result.Unom)
	if err != nil {
		return shared.Incident{}, err
	}

	return result, nil
}

// CreateIncident создание нового инцидента.
func (s *service) CreateIncident(ctx context.Context, title, status string, priority int, unom int64) (int64, error) {
	ctx, span := s.tracer.Start(
		ctx,
		"data.CreateIncident",
		trace.WithAttributes(
			attribute.String("title", title),
			attribute.String("status", status),
			attribute.Int("priority", priority),
			attribute.Int64("unom", unom),
		),
	)
	defer span.End()

	id, err := s.ir.Create(ctx, title, status, priority, unom)
	if err != nil {
		return 0, err
	}

	return id, nil
}
