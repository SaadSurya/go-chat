package chats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AllChat(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	chats := GetChats(username)
	c.JSON(chats)
	return nil
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/chat/all", AllChat)
}
