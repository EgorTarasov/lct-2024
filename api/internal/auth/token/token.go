package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EgorTarasov/lct-2024/api/internal/shared/constants"
	"github.com/golang-jwt/jwt/v5"
)

// UserPayload пользовательские данные в jwt token.
type UserPayload struct {
	UserID   int64          `json:"user_id"`
	AuthType string         `json:"auth_type"`
	Role     constants.Role `json:"role"`
}

// UserClaims jwt token с пользовательскими данными.
type UserClaims struct {
	UserPayload
	jwt.RegisteredClaims
}

// Key - для подписывания jwt.
const Key = "supersecretkey"

// Encode создание jwt токена.
func Encode(_ context.Context, data UserPayload) (string, error) {
	payload := UserClaims{
		UserPayload: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "larek.tech",
			Subject:  "",
			Audience: nil,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour),
			},
			NotBefore: nil,
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			ID: "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// FIXME move Key into cfg
	t, err := token.SignedString([]byte(Key))
	if err != nil {
		return "", fmt.Errorf("jwt err: %v", err.Error())
	}
	return t, nil
}

// Decode получение данных из jwt токена.
func Decode(_ context.Context, tokenString string) (UserPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	if err != nil {
		return UserPayload{}, nil
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return claims.UserPayload, nil
	}
	return UserPayload{}, errors.New("invalid jwt token")
}
