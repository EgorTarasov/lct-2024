package service

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/search/models"
	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// service для поиска по данным системы.
type service struct {
	ar addressRegistry
	cr statePropertyRepo
	fr filterRepo
	tr trace.Tracer
}

// addressRegistry справочник с гео информацией.
type addressRegistry interface {
	GetGeoDataInRadius(ctx context.Context, latitude, longitude float64, distance int) (result []shared.Address, err error)
	GetGeoDataByUnom(ctx context.Context, unom int64) (shared.Address, error)
	GetGeoDataByUnoms(ctx context.Context, unoms []int64) ([]shared.Address, error)
	GetUnomsInRadius(ctx context.Context, latitude, longitude float64, distance int) ([]int64, error)
	GetMunicipalDistricts(ctx context.Context) ([]string, error)
	GetByMunicipalDistrict(ctx context.Context, municipalDistricts []string) ([]shared.Address, error)
}

type filterRepo interface {
	GetSearchFilter(ctx context.Context) (filters []models.Filter, err error)
}

// statePropertyRepo доступ к данным о потребителях.
type statePropertyRepo interface {
	SearchSimilarObjects(ctx context.Context, query string) (result []models.StateProperty, err error)
}

// NewService конструктор сервиса для поиска по данным.
func NewService(ar addressRegistry, cr statePropertyRepo, fr filterRepo, tracer trace.Tracer) *service {
	return &service{
		ar: ar,
		fr: fr,
		cr: cr,
		tr: tracer,
	}
}

// SearchStateProperties поиск по тексту социальных / промышленных объектов.
func (s *service) SearchStateProperties(ctx context.Context, query string) ([]models.StatePropertySearchResult, error) {
	ctx, span := s.tr.Start(ctx, "service.SearchObjects", trace.WithAttributes(attribute.String("query", query)))
	defer span.End()
	stateProperties, err := s.cr.SearchSimilarObjects(ctx, query)
	result := make([]models.StatePropertySearchResult, len(stateProperties))
	if err != nil {
		return nil, err
	}

	// find geoData by unom
	unomIdx := make(map[int64]int, len(stateProperties))
	unoms := make([]int64, 0, len(stateProperties))

	for idx, property := range stateProperties {
		unomIdx[property.Unom] = idx
		unoms = append(unoms, property.Unom)
	}

	geoData, err := s.ar.GetGeoDataByUnoms(ctx, unoms)
	if err != nil {
		return nil, err
	}

	for _, value := range geoData {
		result[unomIdx[value.Unom]].Address = value
		result[unomIdx[value.Unom]].StateProperty = stateProperties[unomIdx[value.Unom]]
	}

	return result, err
}

// ListFilters получение списка всех фильтров для поиска по коллекции consumers.
func (s *service) ListFilters(ctx context.Context) ([]models.Filter, error) {
	ctx, span := s.tr.Start(ctx, "service.ListFilters")
	defer span.End()

	filters, err := s.fr.GetSearchFilter(ctx)
	if err != nil {
		return nil, err
	}
	return filters, nil
}
