package models

import (
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/shared/constants"
)

// VkUserData данные для авторизации по вк
// используется для передачи данных от сервиса к бд.
type VkUserData struct {
	UserID    int64
	VkID      int64
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate time.Time
	City      string
	Photo     string
	Sex       constants.Sex
}

// VkUserDataDao запись о пользователе в бд.
type VkUserDataDao struct {
	UserID    int64         `db:"user_id"`
	VkID      int64         `db:"vk_id"`
	FirstName string        `db:"first_name"`
	LastName  string        `db:"last_name"`
	BirthDate time.Time     `db:"birth_date"`
	City      string        `db:"city"`
	Photo     string        `db:"photo"`
	Sex       constants.Sex `db:"sex"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
	DeletedAt time.Time     ``
}
