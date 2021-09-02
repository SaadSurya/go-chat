package authentication

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	_ "github.com/golang-jwt/jwt/v4"
)

func GetJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})
}
