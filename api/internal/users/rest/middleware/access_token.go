package middleware

import (
	"github.com/EgorTarasov/lct-2024/api/internal/users/token"
	jwtware "github.com/gofiber/contrib/jwt"
)

// UserClaimsMiddleware - middleware для получение данных из jwt token.
var UserClaimsMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(token.Key)},
	ContextKey: "userClaims",
	Claims:     &token.UserClaims{},
	KeyFunc:    nil,
	JWKSetURLs: nil,
})
