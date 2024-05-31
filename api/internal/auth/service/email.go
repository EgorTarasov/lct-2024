package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/EgorTarasov/lct-2024/api/internal/auth/models"
	"github.com/EgorTarasov/lct-2024/api/internal/auth/token"

	"golang.org/x/crypto/bcrypt"
)

type EmailUserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int64, error)
	GetById(ctx context.Context, id int64) (models.UserDao, error)
	CreateEmail(ctx context.Context, userId int64, email, password, ip string) error
	GetPasswordHash(ctx context.Context, email string) (int64, string, error)
	UpdateEmailUsage(ctx context.Context, userId int64, ip string) error
	GetUserId(ctx context.Context, email string) (int64, error)
}

// CreateUserEmail создание аккаунта пользователя с использованием email + паролю
func (s *service) CreateUserEmail(ctx context.Context, data models.UserCreate, email, password, ip string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "service.CreateUserEmail")
	defer span.End()

	id, err := s.ur.GetUserId(ctx, email)
	if err != nil || id != 0 {
		return "", fmt.Errorf("user with given email already exists")
	}

	id, err = s.ur.Create(ctx, data)
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	err = s.ur.CreateEmail(ctx, id, email, string(hash), ip)
	if err != nil {
		return "", err
	}

	//// создание токена по id пользователя
	user, err := s.ur.GetById(ctx, id)
	if err != nil {
		return "", err
	}

	accessToken, err := token.Encode(ctx, token.UserPayload{
		UserId:   user.Id,
		AuthType: "email",
		Role:     user.Role,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// AuthorizeEmail авторизация в приложении с использованием email + пароль
func (s *service) AuthorizeEmail(ctx context.Context, email, password, ip string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "service.AuthorizeEmail")
	defer span.End()
	id, passwordHash, err := s.ur.GetPasswordHash(ctx, email)
	if err != nil {
		return passwordHash, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return passwordHash, errors.New("invalid password")
	}

	err = s.ur.UpdateEmailUsage(ctx, id, ip)
	if err != nil {
		return "", err
	}

	// создание jwt payload
	user, err := s.ur.GetById(ctx, id)
	if err != nil {
		return "", err
	}

	accessToken, err := token.Encode(ctx, token.UserPayload{
		UserId:   user.Id,
		AuthType: "email",
		Role:     user.Role,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}