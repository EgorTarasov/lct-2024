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

// NewAddressRegistryRepository справочная
func NewAddressRegistryRepository(mongo mongo.Mongo, tracer trace.Tracer) *addressRegistryRepository {
	return &addressRegistryRepository{
		mongo:  mongo,
		tracer: tracer,
	}
}

// GetGeoDataByUnom возвращает геоданные по уникальному номеру объекта недвижимости.
func (r *addressRegistryRepository) GetGeoDataByUnom(ctx context.Context, unom int64) (result models.Address, err error) {
	ctx, span := r.tracer.Start(ctx, "addressRegistry.GetGeoDataByUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	err = r.mongo.FindOne(ctx, addressCollectionName,
		bson.D{{Key: "unom", Value: unom}},
		result,
		options.
			FindOne().
			SetProjection(addressProjection),
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

// GetGeoDataInRadius
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
