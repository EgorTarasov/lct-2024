package pg

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/geo/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"go.opentelemetry.io/otel/trace"
)

type objectRepo struct {
	pg     *postgres.Database
	tracer trace.Tracer
}

// NewObjectRepo конструктор репозитория для работы с данными об объектах недвижимости.
func NewObjectRepo(pg *postgres.Database, tracer trace.Tracer) *objectRepo {
	return &objectRepo{
		pg:     pg,
		tracer: tracer,
	}
}

// SelectByID возвращает объект недвижимости по его глобальному идентификатору.
const SelectByID = `
select 
	object.global_id,
	object.address,
	object.in_moscow,
	object.type,
	object.unom,
	object.kladr,
	object.fias,
	object.kad_n,
	object.kad_zu,
	object.geo_data,
	object.geodata_center
from object where global_id = $1;
`

// SelectByID возвращает объект недвижимости по его глобальному идентификатору.
func (or *objectRepo) SelectByID(ctx context.Context, globalID int64) (*models.PropertyDAO, error) {
	var objectDao models.PropertyDAO
	row := or.pg.ExecQueryRow(ctx, SelectByID, globalID)
	err := row.Scan(&objectDao.GlobalID,
		&objectDao.Address,
		&objectDao.InMoscow,
		&objectDao.Type,
		&objectDao.Unom,
		&objectDao.Kladr,
		&objectDao.GeoData,
		&objectDao.GeoDataCenter,
	)
	if err != nil {
		return nil, err
	}
	return &objectDao, nil
}
