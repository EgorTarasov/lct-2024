package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/EgorTarasov/lct-2024/api/internal/users/models"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type userDataRepo struct {
	pg     *postgres.Database
	tracer trace.Tracer
}

// NewDataRepo конструктор репозитория для работы с данными пользовательских аккаунтов
// содержит информацию об статусе обработки загруженных пользователем файлов.
func NewDataRepo(pg *postgres.Database, tracer trace.Tracer) *userDataRepo {
	return &userDataRepo{
		pg:     pg,
		tracer: tracer,
	}
}

// CheckIdempotencyKey проверяет наличие ключа идемпотентности в базе данных.
func (repo *userDataRepo) CheckIdempotencyKey(ctx context.Context, idempotencyKey string) (bool, error) {
	ctx, span := repo.tracer.Start(ctx, "data.CheckIdempotencyKey", trace.WithAttributes(attribute.String("idempotencyKey", idempotencyKey)))
	defer span.End()
	const q = `select exists(select 1 from uploads where idempotency_key = $1);`
	var exists bool
	if err := repo.pg.Get(ctx, &exists, q, idempotencyKey); err != nil {
		return false, fmt.Errorf("can't check idempotency key: %w", err)
	}
	return exists, nil
}

// CreateUpload создает запись о новом файле в системе.
func (repo *userDataRepo) CreateUpload(ctx context.Context, filename, idempotencyKey, s3Key string, userID int64) (int64, error) {
	ctx, span := repo.tracer.Start(ctx, "data.CreateUploads", trace.WithAttributes(attribute.StringSlice("filename, s3Key", []string{filename, s3Key})))
	defer span.End()
	var id int64

	exists, err := repo.CheckIdempotencyKey(ctx, idempotencyKey)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("idempotency key already exists")
	}

	const q = `
insert into uploads(filename, status, idempotency_key, s3_key, user_id) values ($1, 'pending', $2, $3, $4) returning id;
`
	if err := repo.pg.Get(ctx, &id, q, filename, idempotencyKey, s3Key, userID); err != nil {
		return 0, fmt.Errorf("can't create upload: %w", err)
	}
	return id, nil
}

// UpdateStatus обновляет статус обработки файла.
func (repo *userDataRepo) UpdateStatus(ctx context.Context, id int64, status string) error {
	ctx, span := repo.tracer.Start(ctx, "data.UpdateStatus", trace.WithAttributes(attribute.Int64("id", id), attribute.String("status", status)))
	defer span.End()
	const q = `update uploads set status = $1 where id = $2;`
	if _, err := repo.pg.Exec(ctx, q, status, id); err != nil {
		return fmt.Errorf("can't update upload status: %w", err)
	}
	return nil
}

// GetUpload возвращает информацию о файле по его идентификатору.
func (repo *userDataRepo) GetUpload(ctx context.Context, id int64) (models.Upload, error) {
	ctx, span := repo.tracer.Start(ctx, "data.GetUpload", trace.WithAttributes(attribute.Int64("id", id)))
	defer span.End()
	const q = `select 
	id,
	idempotency_key,
	filename,
	s3_key,
	user_id,
	status,
	created_at
	from uploads where id = $1;
`
	var upload models.Upload
	if err := repo.pg.Get(ctx, &upload, q, id); err != nil {
		return upload, fmt.Errorf("can't get upload: %w", err)
	}
	return upload, nil
}

// GetAllUploads возвращает список всех файлов загруженных пользователями.
func (repo *userDataRepo) GetAllUploads(ctx context.Context) ([]models.Upload, error) {
	ctx, span := repo.tracer.Start(ctx, "data.GetAllUploads")
	defer span.End()
	const q = `select 
	id,
	idempotency_key,
	filename,
	s3_key,
	user_id,
	status,
	created_at
	from uploads;
`
	var uploads []models.Upload
	if err := repo.pg.Select(ctx, &uploads, q); err != nil {
		return uploads, fmt.Errorf("can't get uploads: %w", err)
	}
	return uploads, nil
}
