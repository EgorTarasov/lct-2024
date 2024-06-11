package service

import (
	"context"
)

type odsRepository interface {
	GetAllCPH(ctx context.Context) ([]string, error)
}

// GetAllCPH returns all CHP from the database.
func (s *service) GetAllCPH(ctx context.Context) ([]string, error) {
	ctx, span := s.tracer.Start(ctx, "service.GetAllCPH")
	defer span.End()

	return s.ods.GetAllCPH(ctx)
}
