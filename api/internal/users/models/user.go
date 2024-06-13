package models

import (
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/shared/constants"
)

// UserCreate данные необходимые для записи о пользователе.
type UserCreate struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// avatar    string         `` // TODO: add profile pictures via s3.
	Role constants.Role `json:"role"`
}

// UserDao представление пользователя на Data слое.
type UserDao struct {
	ID        int64          `db:"id" json:"id"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  string         `db:"last_name" json:"last_name"`
	Role      constants.Role `db:"role" json:"role"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt time.Time      ``
}

// Dto преобразование UserDao -> UserDto.
func (ud *UserDao) Dto() UserDto {
	return UserDto{
		ID:        ud.ID,
		FirstName: ud.FirstName,
		LastName:  ud.LastName,
	}
}

// UserDto модель пользователя на уровне сервиса.
type UserDto struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
