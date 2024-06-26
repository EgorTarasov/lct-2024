package reports

import (
	"context"
	"time"

	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"go.opentelemetry.io/otel/trace"
)

type reportRepo struct {
	pg     *postgres.Database
	tracer trace.Tracer
}

// NewReportRepo конструктор репозитория для работы с отчетами.
func NewReportRepo(pg *postgres.Database, tracer trace.Tracer) *reportRepo {
	return &reportRepo{
		pg:     pg,
		tracer: tracer,
	}
}

// CreateReport creates new report.
func (repo *reportRepo) CreateReport(ctx context.Context, reportCreate CreateReport) (int64, error) {
	ctx, span := repo.tracer.Start(ctx, "reports.CreateReport")
	defer span.End()
	const q = `
insert into reports(
                    	title,
                    	user_id,
                    	start_date,
                    	end_date
) values ($1, $2, $3, $4) returning id;
`
	var id int64
	if err := repo.pg.Get(ctx, &id, q, reportCreate.Title, reportCreate.UserID, reportCreate.StartDate, reportCreate.EndDate); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *reportRepo) AddS3Key(ctx context.Context, s3Key string, id int64) error {
	ctx, span := repo.tracer.Start(ctx, "reports.UpdateReport")
	defer span.End()
	const q = `
update reports
set s3_key = $1,
    updated_at = now(),
    status = $2
where id = $3;
`
	if _, err := repo.pg.Exec(ctx, q, s3Key, ReportStateReady, id); err != nil {
		return err
	}
	return nil
}

func (repo *reportRepo) GetReportByID(ctx context.Context, id int64) (Report, error) {
	ctx, span := repo.tracer.Start(ctx, "reports.GetReportByID")
	defer span.End()
	const q = `select * from reports where id = $1;`
	var report Report
	if err := repo.pg.Get(ctx, &report, q, id); err != nil {
		return Report{}, err
	}
	return report, nil
}

// GetUsersReports returns all reports for given user.
func (repo *reportRepo) GetUsersReports(ctx context.Context, userID int64) ([]Report, error) {
	ctx, span := repo.tracer.Start(ctx, "reports.GetUsersReports")
	defer span.End()
	const q = `select * from reports where user_id = $1;`
	var reports []Report
	if err := repo.pg.Select(ctx, &reports, q, userID); err != nil {
		return nil, err
	}
	return reports, nil
}

func (repo *reportRepo) GetPredictionRecords(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]PredictionRecord, error) {
	ctx, span := repo.tracer.Start(ctx, "reports.GetPredictionRecords")
	defer span.End()
	const q = `select 
id,
p1_less_than_or_equal_to_0,
p2_less_than_or_equal_to_0,
t1_less_than_min,
t1_greater_than_max,
no,
lack_of_heating_in_the_house,
pipe_leak_in_the_entrance,
strong_leak_in_the_heating_system,
temperature_in_the_apartment_below_the_standard,
temperature_in_public_areas_below_the_standard,
leak_in_the_heating_system
from prediction_records 
where prediction_date between $1 and $2
limit $3 
offset $4;`
	var records []PredictionRecord
	if err := repo.pg.Select(ctx, &records, q, startDate, endDate, limit, offset); err != nil {
		return nil, err
	}
	return records, nil
}

func (repo *reportRepo) GetPredictionCount(ctx context.Context, startDate, endDate time.Time) (int64, error) {
	ctx, span := repo.tracer.Start(ctx, "reports.GetPredictionCount")
	defer span.End()
	const q = `
select 
    count(*)
from prediction_records
where prediction_date between $1 and $2;
;`
	var count int64
	if err := repo.pg.Get(ctx, &count, q, startDate, endDate); err != nil {
		return 0, err
	}
	return count, nil
}
