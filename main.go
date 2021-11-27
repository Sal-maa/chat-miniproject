package main

import (
	"chat-miniproject/config"
	"chat-miniproject/contract"
	"chat-miniproject/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DatabaseMigration() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	db.AutoMigrate(contract.Users{}, contract.Chat{})
}

func main() {
	DatabaseMigration()

	app := fiber.New()

	app.Post("/users", handler.CreateNewUser)
	app.Get("/users", handler.GetAllUser)
	app.Post("/users/:id", handler.SendNewText)
	app.Get("/users/:id", handler.ReadNewText)
	app.Get("/users/:id?name", handler.GetChatRoom)

	app.Listen(":8080")
}
