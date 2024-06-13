package service

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/geo/models"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type objectRepo interface {
	SelectByID(ctx context.Context, globalID int64) (*models.PropertyDAO, error)
	SelectInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.PropertyDAO, error)
}

type moekRepo interface {
	SelectInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.MoekDAO, error)
}

func (s *service) GetObjectByID(ctx context.Context, globalID int64) (*models.PropertyDTO, error) {
	ctx, span := s.tracer.Start(ctx, "service.GetObjectByID", trace.WithAttributes(attribute.Int64("id", globalID)))
	defer span.End()

	propertyDao, err := s.or.SelectByID(ctx, globalID)
	if err != nil {
		return nil, err
	}

	return propertyDao.ToDTO(), nil
}

// GetObjectsInRadius возвращает список объектов недвижимости в радиусе от заданной точки.
func (s *service) GetObjectsInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.PropertyDTO, error) { //nolint
	ctx, span := s.tracer.Start(ctx, "service.GetObjectsInRaduis", trace.WithAttributes(attribute.Float64("longitude", longitude), attribute.Float64("latitude", latitude), attribute.Float64("radius", radius)))
	defer span.End()

	propertiesDao, err := s.or.SelectInRadius(ctx, longitude, latitude, radius)
	if err != nil {
		return nil, err
	}
	result := make([]*models.PropertyDTO, len(propertiesDao))

	for idx, property := range propertiesDao {
		result[idx] = property.ToDTO()
	}

	return result, nil
}

// GetMoeksInRadius возвращает список моек в радиусе от заданной точки.
func (s *service) GetMoeksInRadius(ctx context.Context, longitude, latitude, radius float64) ([]*models.MoekDTO, error) { //nolint
	ctx, span := s.tracer.Start(ctx, "service.GetMoeksInRadius", trace.WithAttributes(attribute.Float64("longitude", longitude), attribute.Float64("latitude", latitude), attribute.Float64("radius", radius)))
	defer span.End()

	moeksDao, err := s.moek.SelectInRadius(ctx, longitude, latitude, radius)
	if err != nil {
		return nil, err
	}
	result := make([]*models.MoekDTO, len(moeksDao))

	for idx, moek := range moeksDao {
		result[idx] = moek.ToDTO()
	}

	return result, nil
}
