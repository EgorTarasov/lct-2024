package mongo

import (
	"context"

	models2 "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// eventRepo структура для работы с state_properties.
type eventRepo struct {
	m      mongoDB.Mongo
	tracer trace.Tracer
}

// NewEventRepo конструктор репозитория для работы с коллекцией state_properties.
func NewEventRepo(m mongoDB.Mongo, tracer trace.Tracer) *eventRepo {
	return &eventRepo{
		m:      m,
		tracer: tracer,
	}
}

// GetEmergencyEvents получение аварийных ситуаций и информации об объекте
// отображение только не завершенных событий.
func (repo *eventRepo) GetEmergencyEvents(ctx context.Context, unoms []int64) (events []models2.Event, err error) {
	ctx, span := repo.tracer.Start(ctx, "eventRepo.GetEmergencyEvents")
	defer span.End()

	filter := bson.D{
		{
			Key:   "external_closed_at",
			Value: nil,
		},
		{
			Key: "unom",
			Value: bson.D{
				{
					Key:   "$in",
					Value: unoms,
				},
			},
		},
	}

	err = repo.m.FindMany(ctx, "events", filter,
		&events,
	)
	log.Info().Int("len", len(events)).Interface("err", err).Msg("events")
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *eventRepo) GetMeasurementsByUnom(ctx context.Context, unom int64, limit int64) (measurements []models2.Measurement, err error) {
	ctx, span := repo.tracer.Start(ctx,
		"eventRepo.GetMeasurementsByUnom",
		trace.WithAttributes(
			attribute.Int64("unom", unom),
		),
	)

	defer span.End()

	filter := bson.M{
		"unom": bson.M{
			"$eq": unom,
		},
	}
	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSort(bson.D{{Key: "datetime", Value: -1}})

	err = repo.m.FindMany(ctx, "property_metering_devices", filter, &measurements, opts)

	if err != nil {
		return nil, err
	}

	return measurements, nil
}
