package mongo

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/geo/models"
	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// propertyCollectionNane is a name of the collection in the mongoDB.
const propertyCollectionNane = "properties"

// propertyRepository is a repository for properties.
type propertyRepository struct {
	db *mongoDB.Mongo
	tr trace.Tracer
}

// NewPropertyRepository creates a new property repository.
func NewPropertyRepository(db *mongoDB.Mongo, tracer trace.Tracer) *propertyRepository {
	return &propertyRepository{
		db: db,
		tr: tracer,
	}
}

// SelectByID selects a property by its global ID.
func (pr *propertyRepository) SelectByID(ctx context.Context, globalID int64) (*models.PropertyDAO, error) {
	ctx, span := pr.tr.Start(ctx, "mongo.propertyRepository.SelectByID", trace.WithAttributes(attribute.Int64("id", globalID)))
	defer span.End()

	var property models.PropertyDAO
	err := pr.db.FindOne(ctx, propertyCollectionNane, bson.D{{Key: "global_id", Value: globalID}}, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

// SelectInRadius selects properties in a radius from the given point.
func (pr *propertyRepository) SelectInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.PropertyDAO, error) { //nolint
	ctx, span := pr.tr.Start(ctx, "mongo.propertyRepository.SelectInRadius", trace.WithAttributes(attribute.Float64("longitude", longitude), attribute.Float64("latitude", latitude), attribute.Float64("radius", radius)))
	defer span.End()

	var properties []*models.PropertyDAO
	err := pr.db.FindMany(ctx, propertyCollectionNane, bson.D{
		{Key: "geodata_center",
			Value: bson.D{
				{
					Key: "$near",
					Value: bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": bson.A{latitude, longitude},
						},
						"$maxDistance": radius,
					},
				},
			},
		},
	}, &properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
