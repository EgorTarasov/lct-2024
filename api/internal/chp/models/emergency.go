package models

import (
	shared "github.com/EgorTarasov/lct-2024/api/internal/shared/models"
)

type EmergencyInfoDTO struct {
	Purpose string `json:"purpose"`
}

// Event информация о событие
type Event struct {
	// Округ
	Country string `json:"country" bson:"country"`
	// Описание
	Title string `json:"title" bson:"title"`
	// Источник
	Source string `json:"source" bson:"src"`
	// Дата начала
	StartDate string `json:"startDate" bson:"external_created_at"`
	// Дата окончания
	EndedAt string `json:"EndedAt" bson:"external_closed"`
	// Дата закрытия в системе
	ExternalEndedAt string `json:"externalEndedAt" bson:"external_ended_at"`

	Unom int64 `json:"unom" bson:"unom"`

	Address shared.Address `json:"geo" bson:"geo"`
}
