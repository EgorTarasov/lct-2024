package models

import (
	"time"
)

type PredictionMessage struct {
	Unoms        []int64   `json:"unoms"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Threshold    float32   `json:"threshold"`
	RegionName   string    `json:"regionName"`
	PredictionID int64     `json:"predictionID"`
}

type UploadMessage struct {
	UploadID int64  `json:"uploadID"`
	S3Key    string `json:"s3Key"`
}
