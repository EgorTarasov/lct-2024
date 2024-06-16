package service

import (
	"context"
	"encoding/json"
	"math"
	"time"

	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/kafka"
	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// addressRegistry справочник с гео информацией.
type addressRegistry interface {
	GetByAdmArea(ctx context.Context, admArea string) ([]shared.Address, error)
	GetGeoDataByUnom(ctx context.Context, unom int64) (result shared.Address, err error)
	GetAllObjectsByUnom(ctx context.Context, unom int64) (shared.UnomResult, error)
	GetMkdConsumersByHeatingPoint(ctx context.Context, unom int64) ([]shared.MKDConsumer, error)
	GetConsumersUnomsByHeatingPoint(ctx context.Context, unom int64) ([]int64, error)
}

type eventRepo interface {
	GetEmergencyEvents(ctx context.Context, unoms []int64) (events []shared.Event, err error)
	GetMeasurementsByUnom(ctx context.Context, unom int64, limit int64) (measurements []shared.Measurement, err error)
}

type incidentRepo interface {
	GetByID(ctx context.Context, id int64) (shared.Incident, error)
	Create(ctx context.Context, title, status string, priority int, unom int64) (int64, error)
	GetRecent(ctx context.Context, limit, offset int) ([]shared.Incident, error)
	GetByConsumerUnom(ctx context.Context, unom int64) ([]shared.Incident, error)
	CreateCalculationRecord(ctx context.Context, userID int64, admArea string) (int64, error)
}

type service struct {
	ar       addressRegistry
	ev       eventRepo
	ir       incidentRepo
	producer *kafka.Producer
	tracer   trace.Tracer
	topic    string
}

// NewService конструктор сервиса для работы с картой.
func NewService(ar addressRegistry, ev eventRepo, ir incidentRepo, producer *kafka.Producer, kafkaTopic string, tracer trace.Tracer) *service {
	return &service{
		ar:       ar,
		ev:       ev,
		ir:       ir,
		producer: producer,
		topic:    kafkaTopic,
		tracer:   tracer,
	}
}

func buildSaramaMessage(predictionID int64, unoms []int64, startDate, endDate time.Time, threshold float32, regionName string) (*sarama.ProducerMessage, error) {
	jsonMessage := shared.PredictionMessage{
		Unoms:        unoms,
		StartDate:    startDate,
		EndDate:      endDate,
		Threshold:    threshold,
		RegionName:   regionName,
		PredictionID: predictionID,
	}

	rawBytes, err := json.Marshal(jsonMessage)
	if err != nil {
		return nil, err
	}
	messageHeader := sarama.RecordHeader{
		Key:   []byte("type"),
		Value: []byte("predict"),
	}

	return &sarama.ProducerMessage{
		Topic:     "",
		Value:     sarama.ByteEncoder(rawBytes),
		Partition: -1,
		Headers: []sarama.RecordHeader{
			messageHeader,
		},
	}, nil
}

func (s *service) GetEmergencyPredictions(ctx context.Context, userID int64, admArea string, startDate, endDate time.Time, threshold float32) error {
	_, span := s.tracer.Start(
		ctx,
		"data.GetEmergencyPredictions",
		trace.WithAttributes(
			attribute.String("admArea", admArea),
			attribute.String("startDate", startDate.String()),
			attribute.String("endDate", endDate.String()),
			attribute.Float64("threshold", float64(threshold)),
		),
	)
	defer span.End()

	addresses, err := s.ar.GetByAdmArea(ctx, admArea)
	if err != nil {
		return err
	}
	unoms := make([]int64, len(addresses))
	for idx, value := range addresses {
		unoms[idx] = value.Unom
	}

	// create message for prediction service
	// send message to kafka

	newID, err := s.ir.CreateCalculationRecord(ctx, userID, admArea)
	if err != nil {
		return err
	}

	msg, err := buildSaramaMessage(newID, unoms, startDate, endDate, threshold, admArea)
	if err != nil {
		return err
	}
	msg.Topic = s.topic

	s.producer.SendAsyncMessage(msg)
	// create record about calculations

	return nil
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
	for i, incident := range recent {

		result, er := s.ar.GetAllObjectsByUnom(ctx, incident.Unom)
		if er != nil {
			log.Error().Err(er).Msg("can't get objects by unom")
		}
		if result.Consumers != nil {
			recent[i].Consumer = result.Consumers
		} else {
			recent[i].Consumer = nil
		}
		if result.HeatingPoint != nil {
			recent[i].HeatingPoint = result.HeatingPoint
		} else {
			recent[i].HeatingPoint = nil
		}

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

func (s *service) GetIncedentsByHeatingPoint(ctx context.Context, unom int64) ([]shared.Incident, error) {
	ctx, span := s.tracer.Start(
		ctx,
		"data.GetIncedentsByHeatingPoint",
		trace.WithAttributes(
			attribute.Int64("unom", unom),
		),
	)
	defer span.End()

	consumerUnoms, err := s.ar.GetConsumersUnomsByHeatingPoint(ctx, unom)
	if err != nil {
		return nil, err
	}

	incidents := make([]shared.Incident, 0)
	for _, consumerUnom := range consumerUnoms {
		consumerIncidents, err := s.ir.GetByConsumerUnom(ctx, consumerUnom)
		if err != nil || len(consumerIncidents) == 0 {
			continue
		}
		incidents = append(incidents, consumerIncidents...)
	}

	return incidents, nil
}

// GetPredictions получение статусов текущих вычислений.
func (s *service) GetPredictions(ctx context.Context, limit, offset int) ([]shared.PredictionCalculation, error) {
	ctx, span := s.tracer.Start(
		ctx,
		"data.GetPredictions",
		trace.WithAttributes(
			attribute.Int("limit", limit),
			attribute.Int("offset", offset),
		),
	)
	defer span.End()
	return nil, nil
}
