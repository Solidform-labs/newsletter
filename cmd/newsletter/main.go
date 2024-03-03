package main

import (
	"fmt"

	"github.com/Solidform-labs/newsletter/configs"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := configs.GetConfig()

	db.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Server is running on port", config.ApiPort)
	app.Listen(":" + config.ApiPort)
}
