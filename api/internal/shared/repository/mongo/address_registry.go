package mongo

import (
	"context"
	"strings"

	"github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type addressRegistryRepository struct {
	mongo  mongo.Mongo
	tracer trace.Tracer
}

const addressCollectionName = "address_registry"

var addressProjection = bson.D{
	{Key: "unom", Value: 1},
	{Key: "geo_center", Value: 1},
	{Key: "geo_data", Value: 1},
	{Key: "address", Value: 1},
}

// NewAddressRegistryRepository справочная информация по адресам.
func NewAddressRegistryRepository(mongo mongo.Mongo, tracer trace.Tracer) *addressRegistryRepository {
	return &addressRegistryRepository{
		mongo:  mongo,
		tracer: tracer,
	}
}

// GetByAdmArea возвращает список объектов недвижимости по району.
func (r *addressRegistryRepository) GetByAdmArea(ctx context.Context, admArea string) (result []models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetByDistrict", trace.WithAttributes(attribute.String("admArea", admArea)))
	defer span.End()

	filter := bson.M{
		"adm_area": bson.M{
			"$eq": admArea,
		},
	}

	findOptions := options.Find()
	findOptions.SetLimit(10)

	err = r.mongo.FindMany(ctx, addressCollectionName, filter, &result, findOptions)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetGeoDataByUnom возвращает геоданные по уникальному номеру объекта недвижимости.
func (r *addressRegistryRepository) GetGeoDataByUnom(ctx context.Context, unom int64) (result models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetGeoDataByUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	filter := bson.M{"unom": unom}

	err = r.mongo.FindOne(ctx, addressCollectionName,
		filter,
		&result,
	)
	if err != nil {
		return models.Address{}, err
	}

	return result, nil
}

// GetGeoDataByUnoms возвращает геоданные по списку уникальных номеров объектов недвижимости.
func (r *addressRegistryRepository) GetGeoDataByUnoms(ctx context.Context, unoms []int64) (result []models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetGeoDataByUnoms", trace.WithAttributes(attribute.Int64Slice("unoms", unoms)))
	defer span.End()

	err = r.mongo.FindMany(ctx, addressCollectionName, bson.D{{Key: "unom", Value: bson.D{{Key: "$in", Value: unoms}}}}, &result, options.
		Find().
		SetProjection(addressProjection))
	if err != nil {
		return nil, err
	}

	return result, err
}

// GetUnomsInRadius поиск объектов в радиусе возвращает список уникальных номеров недвижимости
// Latitude and longitude coordinates are: 55.751244, 37.618423
// latitude - долгота в нашем случае около 55.5
// longitude - широта в нашем случае около 37.5
// distance - расстояние в пределах которого производить.
//
// Результат - уникальные номера недвижимости которые попадают в указанную область.
func (r *addressRegistryRepository) GetUnomsInRadius(ctx context.Context, latitude, longitude float64, distance int) (unoms []int64, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetUnomsInRadius", trace.WithAttributes(attribute.Float64("latitude", latitude), attribute.Float64("longitude", longitude)))
	defer span.End()

	err = r.mongo.FindMany(ctx, addressCollectionName, bson.D{
		{
			Key: "geo_center",
			Value: bson.D{
				{
					Key: "$near",
					Value: bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": bson.A{latitude, longitude},
						},
						"$maxDistance": distance,
					},
				},
			},
		},
	}, &unoms, options.Find().SetProjection(bson.D{{Key: "unom", Value: 1}}))

	if err != nil {
		return nil, err
	}
	return unoms, nil
}

// GetMunicipalDistricts возвращает список районов.
// Оставляет приставку муниципальный район <value> или городской округ <value>.
func (r *addressRegistryRepository) GetMunicipalDistricts(ctx context.Context) (result []string, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetMunicipalDistricts")
	defer span.End()

	res, err := r.mongo.Distinct(ctx, addressCollectionName, "municipal_district", nil)
	result = make([]string, len(res))
	for i, v := range res {
		result[i] = v.(string)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetByMunicipalDistrict возвращает список объектов недвижимости по району.
func (r *addressRegistryRepository) GetByMunicipalDistrict(ctx context.Context, municipalDistricts []string) (result []models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetByMunicipalDistrict", trace.WithAttributes(attribute.StringSlice("municipalDistrict", municipalDistricts)))
	defer span.End()

	// /Академический|Алексеевкий|Алтуфьевский|Арбат|Аэропорт/
	reStr := "//" + strings.Join(municipalDistricts, "|") + "/"

	err = r.mongo.FindMany(
		ctx,
		addressCollectionName,
		bson.D{{Key: "municipal_district", Value: bson.D{{Key: "$regex", Value: reStr}}}},
		&result,
		options.Find().SetProjection(addressProjection),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGeoDataInRadius возвращает геоданные в радиусе.
func (r *addressRegistryRepository) GetGeoDataInRadius(ctx context.Context, latitude, longitude float64, distance int) (result []models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetGeoDataInRadius", trace.WithAttributes(attribute.Float64("latitude", latitude), attribute.Float64("longitude", longitude)))
	defer span.End()

	filter := bson.M{
		"geo_center": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": bson.A{latitude, longitude},
				},
				"$maxDistance": distance,
			},
		},
	}
	log.Info().Interface("filter", filter).Msg("filter")

	err = r.mongo.FindMany(ctx, addressCollectionName, filter, &result)
	log.Info().Int("len", len(result)).Interface("err", err).Msg("events")
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *addressRegistryRepository) GetHeatSourceByConsumerUnom(ctx context.Context, unom int64) (result models.HeatingPoint, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetHeatSourceByConsumerUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	filter := bson.M{"consumer_full_address.unom": bson.M{"$eq": unom}}
	err = r.mongo.FindOne(ctx, "consumers", filter, &result)
	if err != nil {
		return models.HeatingPoint{}, err
	}
	return result, nil
}

func (r *addressRegistryRepository) GetHetSourceBySrcUnom(ctx context.Context, unom int64) (result models.HeatingPoint, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetHetSourceBySrcUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	filter := bson.M{"heating_point_full_address.unom": bson.M{"$eq": unom}}
	err = r.mongo.FindOne(ctx, "consumers", filter, &result)
	if err != nil {
		return models.HeatingPoint{}, err
	}
	return result, nil
}

// GetAllObjectsByUnom возвращает все объекты по уникальному номеру.
func (r *addressRegistryRepository) GetAllObjectsByUnom(ctx context.Context, unom int64) (models.UnomResult, error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetAllObjectsByUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	filter := bson.M{"unom": unom}
	var model models.UnomResult
	model.Unom = unom
	err := r.mongo.FindOne(ctx, "dispatch_services", filter, &model.DispatchService)
	if err != nil {
		log.Info().Err(err).Msg("error during getting related objects dispatch_services")
		model.DispatchService = nil
	}
	err = r.mongo.FindOne(ctx, "mkd", filter, &model.Consumer)
	if err != nil {
		log.Info().Err(err).Msg("error during getting related objects mkd")
		model.Consumer = nil
	}
	heatPoint, err := r.GetHeatSourceByConsumerUnom(ctx, unom)
	if err != nil {
		log.Info().Err(err).Msg("error during getting related objects consumers")
		model.HeatingPoint = nil
	}
	model.HeatingPoint = &heatPoint

	return model, nil
}

// GetConsumersUnomsByHeatingPoint возвращает список уникальных номеров потребителей по уникальному номеру теплового пункта.
func (r *addressRegistryRepository) GetConsumersUnomsByHeatingPoint(ctx context.Context, unom int64) ([]int64, error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetConsumersUnomsByHeatingPoint", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	var result []models.HeatingPoint
	filter := bson.M{"heating_point_full_address.unom": unom}
	err := r.mongo.FindMany(ctx, "consumers", filter, &result)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(result))
	for i, v := range result {
		ids[i] = v.ConsumerAddress.Unom
	}

	return ids, nil
}
