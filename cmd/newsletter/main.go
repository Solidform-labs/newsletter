package main

import (
	"os"

	"github.com/Solidform-labs/newsletter/configs"
	"github.com/Solidform-labs/newsletter/internal/newsletter/api/routers"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	// Only load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	config := configs.GetConfig()

	defer db.GetDB().Close()

	app := fiber.New()
	defer app.Shutdown()

	routers.Setup(app)

	log.Infof("Server is running on port %s", config.ApiPort)
	app.Listen(":" + config.ApiPort)
}
