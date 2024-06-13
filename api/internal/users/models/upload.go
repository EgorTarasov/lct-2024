package models

import (
	"time"
)

// Upload represents a file uploaded by a user.
type Upload struct {
	ID             int64     `db:"id" json:"id"`
	IdempotencyKey string    `db:"idempotency_key" json:"-"`
	Filename       string    `db:"filename" json:"filename"`
	S3Key          string    `db:"s3_key" json:"-"`
	UserID         int64     `db:"user_id" json:"userID"`
	Status         string    `db:"status" json:"status"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	Url            string    `json:"url"`
}
