package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saadsurya/go-chat/authentication"
	"github.com/saadsurya/go-chat/chats"
	"github.com/saadsurya/go-chat/database"
	"github.com/saadsurya/go-chat/messages"
	"github.com/saadsurya/go-chat/users"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=go_chat sslmode=disable password=10pearls1+")
	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to database")
	}
	log.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&users.User{})
	database.DBConn.AutoMigrate(&messages.Message{})
	log.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	authentication.RegisterRoutes(app)

	app.Use(authentication.GetJWTMiddleware())

	users.RegisterRoutes(app)
	messages.RegisterRoutes(app)
	chats.RegisterRoutes(app)

	app.Listen(":3000")
}
