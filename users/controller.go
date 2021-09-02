package users

import "github.com/gofiber/fiber/v2"

func GetAll(c *fiber.Ctx) error {
	c.JSON(GetAllUsers())
	return nil
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/user/all", GetAll)
}
