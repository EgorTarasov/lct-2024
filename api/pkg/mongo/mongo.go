// Package mongo provides a MongoDB client.
package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo provides a MongoDB client and tracing for operations.
type Mongo struct {
	mongo *mongo.Client
	db    string
}

// MustNew creates a new Mongo instance or panics if failed.
func MustNew(cfg *Config) (Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.RetryTimeout)*time.Second)
	defer cancel()
	log.Info().Str("url", cfg.URL()).Msg("connecting to mongo")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://Dino:Dino-misos2024@192.168.1.70:54002/")) //FIXME: cfg
	if err != nil {
		return Mongo{
			mongo: nil,
			db:    "",
		}, fmt.Errorf("failed to create mongo client")
	}
	return Mongo{
		db:    cfg.DB,
		mongo: client,
	}, nil
}

// Ping checks the connection to the MongoDB.
func Ping(ctx context.Context, m Mongo) error {
	var result bson.M
	if err := m.mongo.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return err
	}
	return nil
}

// Close closes the MongoDB client.
func (m Mongo) Close() error {
	if err := m.mongo.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("failed to close mongo conn: %v", err)
	}
	return nil
}

// InsertOne inserts a single document into the given collection.
func (m Mongo) InsertOne(ctx context.Context, coll string, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	data, err := bson.Marshal(doc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal document")
	}

	res, err := m.mongo.Database(m.db).Collection(coll).InsertOne(ctx, data, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert one document")
	}
	return res, nil
}

// FindOne finds a single document in the given collection.
func (m Mongo) FindOne(ctx context.Context, coll string, filter, dest interface{}, opts ...*options.FindOneOptions) error {
	return errors.Wrap(m.mongo.Database(m.db).Collection(coll).FindOne(ctx, filter, opts...).Decode(dest), "failed to find one document")
}

// FindMany finds multiple documents in the given collection.
func (m Mongo) FindMany(ctx context.Context, coll string, filter, dest interface{}, opts ...*options.FindOptions) error {
	cursor, err := m.mongo.Database(m.db).Collection(coll).Find(ctx, filter, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to find many documents")
	}
	return errors.Wrap(cursor.All(ctx, dest), "failed to decode many documents")
}

// UpdateOne updates a single document in the given collection.
func (m Mongo) UpdateOne(ctx context.Context, coll string, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := m.mongo.Database(m.db).Collection(coll).UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update one document")
	}
	return res, nil
}

// UpdateMany updates multiple documents in the given collection.
func (m Mongo) UpdateMany(ctx context.Context, coll string, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := m.mongo.Database(m.db).Collection(coll).UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update many documents")
	}
	return res, nil
}

// DeleteOne deletes a single document from the given collection.
func (m Mongo) DeleteOne(ctx context.Context, coll string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	res, err := m.mongo.Database(m.db).Collection(coll).DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete one document")
	}
	return res, nil
}
