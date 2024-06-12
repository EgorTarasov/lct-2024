package repository

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/search/models"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const statePropertyCollectionName = "state_property"

type statePropertyRepo struct {
	mongo mongoDB.Mongo
	tr    trace.Tracer
}

// NewStatePropertyRepo конструктор репозитория для работы с данными о состоянии объектов.
func NewStatePropertyRepo(mongo mongoDB.Mongo, tracer trace.Tracer) *statePropertyRepo {
	return &statePropertyRepo{
		mongo: mongo,
		tr:    tracer,
	}
}

// SearchSimilarObjects поиск по тексту социальных / промышленных объектов.
func (repo *statePropertyRepo) SearchSimilarObjects(ctx context.Context, query string) (result []models.StateProperty, err error) {
	ctx, span := repo.tr.Start(ctx, "statePropertyRepo.SearchSimilarObjects", trace.WithAttributes(attribute.String("query", query)))
	defer span.End()

	filter := bson.M{
		"$text": bson.M{
			"$search": "шоссе Энтузиастов",
		},
	}
	queryOptions := options.Find().
		SetSort(bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}).
		SetLimit(10).
		SetProjection(bson.M{"score": bson.M{"$meta": "textScore"}})

	err = repo.mongo.FindMany(ctx, statePropertyCollectionName, filter, &result, queryOptions)

	if err != nil {
		return nil, err
	}
	return result, nil
}
