package repository

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/search/models"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/trace"
)

type searchFilterRepo struct {
	mongo mongoDB.Mongo
	tr    trace.Tracer
}

// NewSearchFilterRepo конструктор репозитория для работы с фильтрами поиска.
func NewSearchFilterRepo(mongo mongoDB.Mongo, tracer trace.Tracer) *searchFilterRepo {
	return &searchFilterRepo{
		mongo: mongo,
		tr:    tracer,
	}
}

var columnMapping = map[string]string{
	"balance_holder":              "Балансодержатель",
	"consumer":                    "Потребитель",
	"heating_point_src":           "Все подключенные к ТЭЦ",
	"municipal_district":          "Район",
	"county":                      "Округ", // TODO: rename in mongo
	"heating_point_location_type": "Тип ТП",
}

const consumerCollection = "consumers"

// GetSearchFilter возвращает фильтры для поиска по consumer collection.
func (repo *searchFilterRepo) GetSearchFilter(ctx context.Context) (filters []models.Filter, err error) {
	ctx, span := repo.tr.Start(ctx, "searchFilterRepo.GetSearchFilter")
	defer span.End()

	for key, value := range columnMapping {

		results, err := repo.mongo.Distinct(ctx, consumerCollection, key, bson.M{})
		if err != nil {
			return nil, err
		}

		values := make([]string, len(results))
		for i, v := range results {
			values[i] = v.(string)

		}

		filter := models.Filter{
			Name:   value,
			Values: values,
		}

		filters = append(filters, filter)
	}
	return filters, nil
}
