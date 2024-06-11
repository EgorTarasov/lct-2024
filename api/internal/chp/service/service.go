package service

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/chp/models"
	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/rs/zerolog/log"
)

// addressRegistry справочник с гео информацией
type addressRegistry interface {
	GetGeoDataInRadius(ctx context.Context, latitude, longitude float64, distance int) (result []shared.Address, err error)
	GetGeoDataByUnom(ctx context.Context, unom int64) (shared.Address, error)
	GetGeoDataByUnoms(ctx context.Context, unoms []int64) ([]shared.Address, error)
	GetUnomsInRadius(ctx context.Context, latitude, longitude float64, distance int) ([]int64, error)
	GetMunicipalDistricts(ctx context.Context) ([]string, error)
	GetByMunicipalDistrict(ctx context.Context, municipalDistricts []string) ([]shared.Address, error)
}

type eventRepo interface {
	GetEmergencyEvents(ctx context.Context, unoms []int64) (events []models.Event, err error)
}

type service struct {
	ar addressRegistry
	ev eventRepo
}

// NewService конструктор сервиса для работы с картой
func NewService(ar addressRegistry, ev eventRepo) *service {
	return &service{
		ar: ar,
		ev: ev,
	}
}

// GetEmergencyEventsInRadius получение аварийных ситуаций и информации об объекте.
func (s *service) GetEmergencyEventsInRadius(ctx context.Context, latitude, longitude float64, distance int) (events []models.Event, err error) {
	addresses, err := s.ar.GetGeoDataInRadius(ctx, latitude, longitude, distance)
	if err != nil {
		return nil, err
	}
	log.Info().Int("addresses", len(addresses)).Msg("addresses in service")

	unomsIdx := make(map[int64]int, len(addresses))
	unoms := make([]int64, 0, len(addresses))
	for i, address := range addresses {
		unomsIdx[address.Unom] = i
		unoms = append(unoms, address.Unom)
	}

	events, err = s.ev.GetEmergencyEvents(ctx, unoms)
	if err != nil {
		return nil, err
	}
	for k, v := range events {
		events[k].Address = addresses[unomsIdx[v.Unom]]
	}

	return events, nil
}

//eventUnom := make(map[int64]*models.Event)
//unoms := make([]int64, 0, len(events))
//for _, event := range events {
//	//eventUnom[event.Unom] = event
//	unoms = append(unoms, event.Unom)
//}
//
//addressInfos, err := s.ar.GetGeoDataByUnoms(ctx, unoms)
//if err != nil {
//	return nil, err
//}
//for _, v := range addressInfos {
//	eventUnom[v.Unom].Address = v
//}
