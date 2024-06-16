package service

import (
	"context"
	"sync"

	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// service для поиска по данным системы.
type service struct {
	ar addressRegistry
	cr statePropertyRepo
	fr consumerRepo
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
	GetHeatSourceByConsumerUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error)
	GetHetSourceBySrcUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error)
}

type consumerRepo interface {
	GetSearchFilter(ctx context.Context) (filters []shared.Filter, err error)
	SearchWithFilters(ctx context.Context, filters []shared.Filter) ([]shared.HeatingPoint, error)
	GetDispatcherInfoByUnoms(ctx context.Context, unoms []int64) ([]shared.DispatchServices, error)
	GetMKDConsumersByUnoms(ctx context.Context, unoms []int64) ([]shared.MKDConsumer, error)
	GetStateConsumersByUnoms(ctx context.Context, unoms []int64) (values []shared.StateConsumer, err error)
}

// statePropertyRepo доступ к данным о потребителях.
type statePropertyRepo interface {
	SearchSimilarObjects(ctx context.Context, query string) (result []shared.StateProperty, err error)
}

// NewService конструктор сервиса для поиска по данным.
func NewService(ar addressRegistry, cr statePropertyRepo, fr consumerRepo, tracer trace.Tracer) *service {
	return &service{
		ar: ar,
		fr: fr,
		cr: cr,
		tr: tracer,
	}
}

// SearchStateProperties поиск по тексту социальных / промышленных объектов.
func (s *service) SearchStateProperties(ctx context.Context, query string) ([]shared.StatePropertySearchResult, error) {
	ctx, span := s.tr.Start(ctx, "service.SearchObjects", trace.WithAttributes(attribute.String("query", query)))
	defer span.End()
	stateProperties, err := s.cr.SearchSimilarObjects(ctx, query)
	result := make([]shared.StatePropertySearchResult, len(stateProperties))
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
func (s *service) ListFilters(ctx context.Context) ([]shared.Filter, error) {
	ctx, span := s.tr.Start(ctx, "service.ListFilters")
	defer span.End()

	filters, err := s.fr.GetSearchFilter(ctx)
	if err != nil {
		return nil, err
	}
	return filters, nil
}

// SearchWithFilters поиск по коллекции consumers с учетом фильтров.
func (s *service) SearchWithFilters(ctx context.Context, filters []shared.Filter) (response []shared.HeatingPoint, err error) {
	ctx, span := s.tr.Start(ctx, "service.SearchWithFilters")
	defer span.End()

	result, err := s.fr.SearchWithFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GeoDataByUnom получение гео данных по уникальному номеру объекта.
func (s *service) GeoDataByUnom(ctx context.Context, unom int64) (shared.Address, error) {
	ctx, span := s.tr.Start(ctx, "service.GeoDataByUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	return s.ar.GetGeoDataByUnom(ctx, unom)
}

// GeoDataByUnoms получение гео данных по уникальным номерам объектов.
func (s *service) GeoDataByUnoms(ctx context.Context, unoms []int64) ([]shared.Address, error) {
	ctx, span := s.tr.Start(ctx, "service.GeoDataByUnoms", trace.WithAttributes(attribute.Int64Slice("unoms", unoms)))
	defer span.End()

	return s.ar.GetGeoDataByUnoms(ctx, unoms)
}

// GetHeatingPointByConsumerUnom получение данных о теплоснабжающем объекте по уникальному номеру.
func (s *service) GetHeatingPointByConsumerUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error) {
	ctx, span := s.tr.Start(ctx, "service.GetHeatingPointByConsumerUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	return s.ar.GetHeatSourceByConsumerUnom(ctx, unom)
}

// GetHeatingPointBySrcUnom получение данных о теплоснабжающем объекте по уникальному номеру.
func (s *service) GetHeatingPointBySrcUnom(ctx context.Context, unom int64) (shared.HeatingPoint, error) {
	ctx, span := s.tr.Start(ctx, "service.GetHeatingPointBySrcUnom", trace.WithAttributes(attribute.Int64("unom", unom)))
	defer span.End()

	return s.ar.GetHetSourceBySrcUnom(ctx, unom)
}

// GetConsumersInfo получение информации о потребителях.
func (s *service) GetConsumersInfo(ctx context.Context, unoms []int64) (interface{}, error) {
	ctx, span := s.tr.Start(ctx, "service.GetConsuerInfo", trace.WithAttributes(attribute.Int64Slice("unoms", unoms)))
	defer span.End()
	// бежим в 3 места и собираем ифнормацию

	var (
		stateConsumers []shared.StateConsumer
		dispatchers    []shared.DispatchServices
		mkdConsumers   []shared.MKDConsumer
	)

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(routineCtx context.Context, unoms []int64) {
		defer wg.Done()
		var err error
		stateConsumers, err = s.fr.GetStateConsumersByUnoms(routineCtx, unoms)
		if err != nil {
			log.Err(err).Msg("failed to get state consumers")
			span.RecordError(err)
		}
	}(ctx, unoms)
	go func(routineCtx context.Context, unoms []int64) {
		defer wg.Done()
		var err error
		dispatchers, err = s.fr.GetDispatcherInfoByUnoms(routineCtx, unoms)
		if err != nil {
			log.Err(err).Msg("failed to get dispatchers")
			span.RecordError(err)
		}
	}(ctx, unoms)
	go func(routineCtx context.Context, unoms []int64) {
		defer wg.Done()
		var err error
		mkdConsumers, err = s.fr.GetMKDConsumersByUnoms(routineCtx, unoms)
		if err != nil {
			log.Err(err).Msg("failed to get mkd consumers")
			span.RecordError(err)
		}
	}(ctx, unoms)

	wg.Wait()

	return struct {
		StateConsumers []shared.StateConsumer    `json:"stateHeatConsumers"`
		Dispatchers    []shared.DispatchServices `json:"dispatchers"`
		MKDConsumers   []shared.MKDConsumer      `json:"mkdConsumers"`
	}{
		StateConsumers: stateConsumers,
		Dispatchers:    dispatchers,
		MKDConsumers:   mkdConsumers,
	}, nil
}
