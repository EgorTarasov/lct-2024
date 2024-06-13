package constants

// Role пользователя в системе.
type Role int

// User - пользователь
// Admin - администратор платформы
// Moderator - модератор платформы.
const (
	User Role = iota
	Admin
	Moderator
)

// Sex Обозначение пола по аналогии с вк
// https://dev.vk.com/ru/reference/objects/user#sex
type Sex int

// Undefined - не указан
// Female - женский
// Male - мужской.
const (
	Undefined Sex = iota
	Female
	Male
)
