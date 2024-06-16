package mongo

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/shared/models"
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
const stateConsumerCollection = "state_property"
const mkdConsumerCollection = "mkd"

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

	for idx := 0; idx < len(values); idx += 1 {
		unomFilter := bson.M{"unom": values[idx].ConsumerAddress.Unom}
		err = repo.mongo.FindMany(ctx, stateConsumerCollection, unomFilter, &values[idx].StateConsumers)
		if err != nil {
			log.Info().Err(err).Int64("unom", values[idx].Unom).Msg("no state consumers found")
		}
		err = repo.mongo.FindMany(ctx, "mkd", unomFilter, &values[idx].MKDConsumers)
		if err != nil {
			log.Info().Err(err).Int64("unom", values[idx].Unom).Msg("no mkd consumers found")
		}

	}

	return values, nil
}

// GetStateConsumersByUnoms возвращает данные о состоянии объектов по их уникальным номерам.
func (repo *consumersSearchRepo) GetStateConsumersByUnoms(ctx context.Context, unoms []int64) (consumers []models.StateConsumer, err error) {
	ctx, span := repo.tr.Start(ctx, "consumersSearchRepo.GetStateConsumersByUnoms")
	defer span.End()

	filter := bson.M{"unom": bson.M{"$in": unoms}}
	if err = repo.mongo.FindMany(ctx, stateConsumerCollection, filter, &consumers); err != nil {
		return nil, err
	}

	for idx := 0; idx < len(consumers); idx += 1 {
		filter = bson.M{"unom": consumers[idx].Unom}
		err = repo.mongo.FindMany(ctx, "events", filter, &consumers[idx].Events)
		if err != nil {
			err = nil
			log.Info().Err(err).Msg("err during retrieving events")
		}
	}

	return consumers, nil
}

// GetMKDConsumersByUnoms возвращает данные о МКД по их уникальным номерам.
func (repo *consumersSearchRepo) GetMKDConsumersByUnoms(ctx context.Context, unoms []int64) (consumers []models.MKDConsumer, err error) {
	ctx, span := repo.tr.Start(ctx, "consumersSearchRepo.GetMKDConsumersByUnoms")
	defer span.End()

	filter := bson.M{"unom": bson.M{"$in": unoms}}
	if err = repo.mongo.FindMany(ctx, mkdConsumerCollection, filter, &consumers); err != nil {
		return nil, err
	}

	for idx := 0; idx < len(consumers); idx += 1 {
		filter = bson.M{"unom": consumers[idx].Unom}
		err = repo.mongo.FindMany(ctx, "events", filter, &consumers[idx].Events)
		if err != nil {
			err = nil
			log.Info().Err(err).Msg("err during retrieving events")
		}
	}

	return consumers, nil
}

// GetDispatcherInfoByUnoms возвращает данные о диспетчерах по уникальным номерам объектов.
func (repo *consumersSearchRepo) GetDispatcherInfoByUnoms(ctx context.Context, unoms []int64) (consumers []models.DispatchServices, err error) {
	ctx, span := repo.tr.Start(ctx, "consumersSearchRepo.GetDispatcherInfoByUnoms")
	defer span.End()

	filter := bson.M{"unom": bson.M{"$in": unoms}}
	if err = repo.mongo.FindMany(ctx, consumerCollection, filter, &consumers); err != nil {
		return nil, err
	}

	for idx := 0; idx < len(consumers); idx += 1 {
		filter = bson.M{"unom": consumers[idx].Unom}
		err = repo.mongo.FindMany(ctx, "events", filter, &consumers[idx].Events)
		if err != nil {
			err = nil
			log.Info().Err(err).Msg("err during retrieving events")
		}
	}

	return consumers, nil
}
