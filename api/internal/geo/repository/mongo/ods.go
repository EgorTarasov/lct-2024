package mongo

import (
	"context"

	mongoDB "github.com/EgorTarasov/lct-2024/api/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/trace"
)

const odsCollectionName = "ods"

// odsCollectionName is a name of the collection in the mongoDB.
type odsRepository struct {
	db *mongoDB.Mongo
	tr trace.Tracer
}

// NewOdsRepository creates a new property repository.
func NewOdsRepository(db *mongoDB.Mongo, tracer trace.Tracer) *odsRepository {
	return &odsRepository{
		db: db,
		tr: tracer,
	}
}

// GetAllCPH returns all CHP from the database.
func (pr *odsRepository) GetAllCPH(ctx context.Context) ([]string, error) {
	result, err := pr.db.Distinct(ctx, odsCollectionName, "ctp", bson.D{
		{
			Key: "ctp",
			Value: bson.D{
				{
					Key:   "$ne",
					Value: "NaN",
				},
			},
		},
	})
	names := make([]string, len(result))
	for i, name := range result {
		names[i] = name.(string)
	}

	if err != nil {
		return nil, err
	}
	return names, nil
}
