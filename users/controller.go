package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetAll(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	c.JSON(GetAllUsers(username))
	return nil
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/user/all", GetAll)
}
