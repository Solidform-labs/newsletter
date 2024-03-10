package main

import (
	"os"

	"github.com/Solidform-labs/newsletter/configs"
	_ "github.com/Solidform-labs/newsletter/docs"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/middleware"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/routers"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

// @title Newsletter API
// @version 1.0
// @description This is Solidform's Newsletter API to handle subscriptions and sending emails
// @host https://newsletter-test-4aaa4eezza-ew.a.run.app
// @basepath /v1
// @schemes https
// @produce json
// @accept json

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
	defer func() {
		log.Info("Server is shutting down")
	}()

	middleware.Setup(app)

	routers.Setup(app)

	log.Infof("Server is running on port %s", config.ApiPort)
	app.Listen(":" + config.ApiPort)
}
