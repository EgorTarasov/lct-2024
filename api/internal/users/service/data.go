package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
	"github.com/EgorTarasov/lct-2024/api/internal/users/models"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type fileRepo interface {
	CreateUpload(ctx context.Context, filename, idempotencyKey, s3Key string, userID int64) (int64, error)
	UpdateStatus(ctx context.Context, id int64, status string) error
	GetUpload(ctx context.Context, id int64) (models.Upload, error)
	GetAllUploads(ctx context.Context) ([]models.Upload, error)
	CheckIdempotencyKey(ctx context.Context, idempotencyKey string) (bool, error)
}

func buildSaramaMessage(uploadID int64, s3Key string) *sarama.ProducerMessage {
	jsonMessage := shared.UploadMessage{
		UploadID: uploadID,
		S3Key:    s3Key,
	}
	rawBytes, _ := json.Marshal(jsonMessage)

	messageHeader := sarama.RecordHeader{
		Key:   []byte("type"),
		Value: []byte("upload"),
	}
	return &sarama.ProducerMessage{
		Topic:     "",
		Value:     sarama.ByteEncoder(rawBytes),
		Partition: -1,
		Headers: []sarama.RecordHeader{
			messageHeader,
		},
	}

}

// CreateUploads создает запись о новом файле в системе.
func (s *service) CreateUploads(ctx context.Context, file io.Reader, filename, idempotencyKey string, fileSize int64, userID int64) (int64, error) {
	s3Key := uuid.New().String()
	ctx, span := s.tracer.Start(ctx, "service.CreateUploads", trace.WithAttributes(attribute.StringSlice("idempotencyKey, filename, s3Key", []string{idempotencyKey, filename, s3Key})))
	defer span.End()

	exists, err := s.data.CheckIdempotencyKey(ctx, idempotencyKey)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("idempotency key already exists")
	}

	// upload file to s3

	_, err = s.s3.PutObject(ctx, "uploads", s3Key, file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return 0, err
	}

	id, err := s.data.CreateUpload(ctx, filename, idempotencyKey, s3Key, userID)
	if err != nil {
		return 0, err
	}
	msg := buildSaramaMessage(id, s3Key)
	msg.Topic = s.topic

	s.producer.SendAsyncMessage(msg)

	return id, err
}

// CheckFileProcessing проверяет статус обработки файла.
func (s *service) CheckFileProcessing(ctx context.Context, id int64) (models.Upload, error) {
	ctx, span := s.tracer.Start(ctx, "service.CheckFileProcessing", trace.WithAttributes(attribute.Int64("id", id)))
	defer span.End()

	upload, err := s.data.GetUpload(ctx, id)
	if err != nil {
		return models.Upload{}, err
	}

	return upload, nil
}

// ListUploads возвращает список всех загруженных файлов.
func (s *service) ListUploads(ctx context.Context) ([]models.Upload, error) {
	ctx, span := s.tracer.Start(ctx, "service.ListUploads")
	defer span.End()

	uploads, err := s.data.GetAllUploads(ctx)
	if err != nil {
		return nil, err
	}

	return uploads, nil
}
