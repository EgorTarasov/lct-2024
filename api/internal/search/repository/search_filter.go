package repository

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/search/models"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/trace"
)

type consumersSearchRepo struct {
	mongo mongoDB.Mongo
	tr    trace.Tracer
}

// NewSearchFilterRepo конструктор репозитория для работы с фильтрами поиска.
func NewSearchFilterRepo(mongo mongoDB.Mongo, tracer trace.Tracer) *consumersSearchRepo {
	return &consumersSearchRepo{
		mongo: mongo,
		tr:    tracer,
	}
}

var columnMappingFromDBToUser = map[string]string{
	"balance_holder":              "Балансодержатель",
	"consumer":                    "Потребитель",
	"heating_point_src":           "Все подключенные к ТЭЦ",
	"municipal_district":          "Район",
	"county":                      "Округ", // TODO: rename in mongo
	"heating_point_location_type": "Тип ТП",
}

var columnMappingFromUserToDB = map[string]string{
	"Балансодержатель":       "balance_holder",
	"Потребитель":            "consumer",
	"Все подключенные к ТЭЦ": "heating_point_src",
	"Тип ТП": "heating_point_location_type",
	"Район":  "municipal_district",
	"Округ":  "county",
}

const consumerCollection = "consumers"

// GetSearchFilter возвращает фильтры для поиска по consumer collection.
func (repo *consumersSearchRepo) GetSearchFilter(ctx context.Context) (filters []models.Filter, err error) {
	ctx, span := repo.tr.Start(ctx, "consumersSearchRepo.GetSearchFilter")
	defer span.End()

	for key, value := range columnMappingFromDBToUser {
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

// SearchWithFilters поиск по consumer collection с учетом фильтров.
func (repo *consumersSearchRepo) SearchWithFilters(ctx context.Context, filters []models.Filter) (values []models.HeatingPoint, err error) {
	ctx, span := repo.tr.Start(ctx, "consumersSearchRepo.SearchWithFilters")
	defer span.End()

	// TODO: add testing for query building
	var filter bson.M
	switch len(filters) {
	case 0:
		filter = bson.M{}
	case 1:
		name, ok := columnMappingFromUserToDB[filters[0].Name]
		if !ok {
			log.Error().Interface("filter", filters[0]).Msg("unknown filter")
			return nil, nil
		}
		filter = bson.M{name: bson.M{"$in": filters[0].Values}}
	default:
		processedFilters := make([]bson.M, 0, len(filters))

		for _, f := range filters {
			if len(f.Values) > 0 {
				name, ok := columnMappingFromUserToDB[f.Name]
				if !ok {
					log.Error().Interface("filter", f).Msg("unknown filter")
					continue
				}
				processedFilters = append(processedFilters, bson.M{name: bson.M{"$in": f.Values}})
			}
		}
		filter = bson.M{"$and": processedFilters}
	}

	log.Debug().Interface("filter", filter).Msg("search filter")

	if err = repo.mongo.FindMany(ctx, consumerCollection, filter, &values); err != nil {
		return nil, err
	}

	return values, nil
}
