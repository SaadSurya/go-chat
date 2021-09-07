package messages

import (
	"strconv"
	"time"

	socket "github.com/saadsurya/go-chat/sockets"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func SendMessage(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	message := new(Message)
	if err := c.BodyParser(message); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return nil
	}

	message.From = username

	CreateMessage(message)

	if s := socket.Sockets[message.To]; s != nil {
		s.WriteJSON(message)
	}

	c.JSON(message)
	return nil
}

func GetMessages(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	ofUser := c.Query("ofUser")

	const layout = "2006-01-02T15:04:05-0700"
	before := time.Now()
	if c.Query("before") != "" {
		beforeParam, err := time.Parse(layout, c.Query("before"))
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(err)
			return nil
		}
		before = beforeParam
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	messages := Retrieve(username, ofUser, before, limit)

	c.JSON(messages)
	return nil
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/message/send", SendMessage)
	app.Get("/message/retrieve", GetMessages)
}
