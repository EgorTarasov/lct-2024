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

const moekCollectionName = "moek"

// propertyRepository is a repository for properties.
type moekRepository struct {
	db *mongoDB.Mongo
	tr trace.Tracer
}

// NewMoekRepository creates a new property repository.
func NewMoekRepository(db *mongoDB.Mongo, tracer trace.Tracer) *moekRepository {
	return &moekRepository{
		db: db,
		tr: tracer,
	}
}

// SelectInRadius selects moek in a radius from the given point.
func (pr *moekRepository) SelectInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.MoekDAO, error) { //nolint
	ctx, span := pr.tr.Start(ctx, "mongo.moekRepository.SelectInRadius", trace.WithAttributes(attribute.Float64("longitude", longitude), attribute.Float64("latitude", latitude), attribute.Float64("radius", radius)))
	defer span.End()

	var moek []*models.MoekDAO
	err := pr.db.FindMany(ctx, moekCollectionName, bson.D{
		{Key: "geoData_center",
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
	}, &moek)
	if err != nil {
		return nil, err
	}

	return moek, nil
}
