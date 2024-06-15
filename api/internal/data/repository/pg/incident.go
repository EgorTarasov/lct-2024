package pg

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/data/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"go.opentelemetry.io/otel/trace"
)

type incidentRepo struct {
	pg     *postgres.Database
	tracer trace.Tracer
}

func NewIncidentRepo(pg *postgres.Database, tracer trace.Tracer) *incidentRepo {
	return &incidentRepo{
		pg:     pg,
		tracer: tracer,
	}
}

// GetRecent возвращает список инцидентов, отсортированных по времени открытия в обратном порядке.
func (repo *incidentRepo) GetRecent(ctx context.Context, limit, offset int) ([]models.Incident, error) {
	ctx, span := repo.tracer.Start(ctx, "data.GetRecent")
	defer span.End()

	const q = `
select 
	id,
	opened_at,
	closed_at,
	title, 
	status,
	priority,
	unom,
	created_at,
	updated_at
from incidents
order by opened_at desc
limit $1 offset $2;
`
	var incidents []models.Incident
	if err := repo.pg.Select(ctx, &incidents, q, limit, offset); err != nil {
		return nil, err
	}
	return incidents, nil
}

// Create создает новый инцидент.
func (repo *incidentRepo) Create(ctx context.Context, title, status string, priority int, unom int64) (int64, error) {
	ctx, span := repo.tracer.Start(ctx, "data.Create")
	defer span.End()

	const q = `
insert into 
    incidents( 
              title,
              status,
              priority,
              unom 
		  )
values ($1, $2, $3, $4)
returning id;
`
	var id int64
	if err := repo.pg.Get(ctx, &id, q, title, status, priority, unom); err != nil {
		return 0, err
	}
	return id, nil
}

// GetByID возвращает инцидент по его идентификатору.
func (repo *incidentRepo) GetByID(ctx context.Context, id int64) (models.Incident, error) {
	ctx, span := repo.tracer.Start(ctx, "data.GetByID")
	defer span.End()

	const q = `
select 
	id,
	opened_at,
	closed_at,
	title, 
	status,
	priority,
	unom,
	created_at,
	updated_at
from incidents
where id = $1;
`
	var incident models.Incident
	if err := repo.pg.Get(ctx, &incident, q, id); err != nil {
		return models.Incident{}, err
	}
	return incident, nil
}
