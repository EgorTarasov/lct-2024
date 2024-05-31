package pg

import (
	"context"

	"github.com/EgorTarasov/lct-2024/api/internal/auth/models"
	"github.com/EgorTarasov/lct-2024/api/internal/auth/repository"
	"github.com/EgorTarasov/lct-2024/api/pkg/postgres"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type userAccountRepo struct {
	pg     *postgres.Database
	tracer trace.Tracer
}

func NewAccountRepo(pg *postgres.Database, tracer trace.Tracer) *userAccountRepo {
	return &userAccountRepo{
		pg:     pg,
		tracer: tracer,
	}
}

// Create сохраняет запись о пользователе в бд
func (ur *userAccountRepo) Create(ctx context.Context, user models.UserCreate) (int64, error) {
	ctx, span := ur.tracer.Start(ctx, "postgres.CreateUser")
	defer span.End()

	var newId int64
	query := `insert into "users"(first_name, last_name)  values($1, $2) returning id;`
	err := ur.pg.Get(ctx, &newId, query, user.FirstName, user.LastName)
	if err != nil {
		log.Error().Err(err).Msg("Create")
		return newId, repository.ErrUserNotCreated
	}
	return newId, nil
}

// GetById получение данных пользователя по id
func (ur *userAccountRepo) GetById(ctx context.Context, id int64) (models.UserDao, error) {
	ctx, span := ur.tracer.Start(ctx, "postgres.GetById")
	defer span.End()

	var user models.UserDao

	query := `
select id, first_name, last_name, role, created_at, updated_at
from "users"
where deleted_at is null and id = $1;`

	err := ur.pg.Get(ctx, &user, query, id)
	if err != nil {
		log.Error().Err(err).Msg("GetById")
		return user, repository.ErrUserNotFound
	}
	return user, nil
}

// CreateEmail создание записи для входа по email
// password - закодировано представление пароля
// ip - адрес пользователя в сети с которого был создан аккаунт
func (ur *userAccountRepo) CreateEmail(ctx context.Context, userId int64, email, password, ip string) error {
	ctx, span := ur.tracer.Start(ctx, "postgres.CreateEmail")
	defer span.End()

	query := `insert into email_auth(user_id, email, password, last_ip) values($1, $2, $3, $4)`

	if _, err := ur.pg.Exec(ctx, query, userId, email, password, ip); err != nil {
		log.Error().Err(err).Msg("CreateEmail")
		return repository.ErrEmailNotCreated
	}

	return nil
}

// GetPasswordHash для авторизации пользователя в системе
func (ur *userAccountRepo) GetPasswordHash(ctx context.Context, email string) (int64, string, error) {
	ctx, span := ur.tracer.Start(ctx, "postgres.GetPasswordHash")
	defer span.End()

	query := `
select user_id, password
from email_auth
where deleted_at is null and email = $1;
`
	var (
		userId         int64
		hashedPassword string
	)
	row := ur.pg.ExecQueryRow(ctx, query, email)
	if err := row.Scan(&userId, &hashedPassword); err != nil {
		log.Error().Err(err).Msg("GetPasswordHash")
		return 0, "", repository.ErrUserPasswordNotFound
	}

	return userId, hashedPassword, nil
}

// UpdateEmailUsage обновление о авторизации в системе
func (ur *userAccountRepo) UpdateEmailUsage(ctx context.Context, userId int64, ip string) error {
	ctx, span := ur.tracer.Start(ctx, "postgres.GetPasswordHash")
	defer span.End()

	query := `
update email_auth
set
    last_used = now(),
    last_ip = $1
where user_id = $2;`

	_, err := ur.pg.Exec(ctx, query, ip, userId)

	if err != nil {
		log.Error().Err(err).Msg("UpdateEmailUsage")
		return repository.ErrEmailUpdate
	}
	return nil
}

// GetVkUserData Получение данных аккаунта пользователя по vk id
func (ur *userAccountRepo) GetVkUserData(ctx context.Context, vkId int64) (models.UserDao, error) {
	ctx, span := ur.tracer.Start(ctx, "postgres.GetVkUserData")
	defer span.End()
	// check if record exists
	// if not create otherwise update
	var (
		dbID int64
		user models.UserDao
	)

	query := `
select vk.vk_id, u.id, u.first_name, u.last_name, u.role, u.created_at, u.updated_at
from vk_auth vk
join "users" u on vk.user_id = u.id
where
    vk.vk_id = $1 and
    vk.deleted_at is null and u.deleted_at is null;
`

	row := ur.pg.ExecQueryRow(ctx, query, vkId)
	err := row.Scan(&dbID,
		&user.Id, &user.FirstName, &user.LastName, &user.Role,
		&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("GetVkUserData")
		return user, repository.ErrVkUserNotFound
	}
	return user, nil
}

// SaveVkUserData сохранение данных от вк
func (ur *userAccountRepo) SaveVkUserData(ctx context.Context, userData models.VkUserData) error {
	ctx, span := ur.tracer.Start(ctx, "postgres.SaveVkUserData")
	defer span.End()

	query := `
insert into vk_auth(user_id, vk_id, first_name, last_name, birth_date, city, photo, sex)
values ($1, $2, $3, $4, $5, $6, $7, $8);
`
	if _, err := ur.pg.Exec(ctx, query, userData.UserId, userData.VkId, userData.FirstName, userData.LastName, userData.BirthDate, userData.City, userData.Photo, userData.Sex); err != nil {
		log.Error().Err(err).Msg("SaveVkUserData")
		return repository.ErrVkUserSave
	}

	return nil
}

// UpdateVkUserData обновление данных от вк
func (ur *userAccountRepo) UpdateVkUserData(ctx context.Context, userData models.VkUserData) error {
	ctx, span := ur.tracer.Start(ctx, "postgres.UpdateVkUserData")
	defer span.End()

	query := `
update vk_auth
set
    first_name = $2,
    last_name = $3,
    birth_date = $4,
    city = $5,
    photo = $6,
    sex = $7
where vk_id = $1 and deleted_at is null;
`
	_, err := ur.pg.Exec(ctx, query, userData.VkId, userData.FirstName, userData.LastName, userData.BirthDate, userData.City, userData.Photo, userData.Sex)
	if err != nil {
		log.Error().Err(err).Msg("UpdateVkUserData")
		return repository.ErrEmailUpdate
	}
	return nil
}

func (ur *userAccountRepo) GetUserId(ctx context.Context, email string) (int64, error) {
	ctx, span := ur.tracer.Start(ctx, "postgres.GetUserId")
	defer span.End()

	query := `
select
	user_id
from email_auth where email = $1 and deleted_at is null;
`
	var userId int64
	if err := ur.pg.Get(ctx, &userId, query, email); err != nil {
		log.Error().Err(err).Msg("GetUserId")
		return userId, repository.ErrUserNotFound
	}
	return userId, nil
}
