package authentication

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saadsurya/go-chat/common"
	"github.com/saadsurya/go-chat/users"
)

func Register(c *fiber.Ctx) error {
	user := new(users.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return nil
	}

	users.CreateUser(user)
	c.JSON(user)
	return nil
}

func Login(c *fiber.Ctx) error {
	user := new(users.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
		return nil
	}

	if user := Authenticate(user.Username, user.Password); user == nil {
		c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{Message: "Invalid username or password"})
		return nil
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["displayName"] = user.DisplayName()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	tokenResponse := Token{
		AuthToken: t,
	}
	c.JSON(tokenResponse)
	return nil
}

func RegisterRoutes(app *fiber.App) {
	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)
}
