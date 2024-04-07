package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Solidform-labs/newsletter/configs"
	_ "github.com/Solidform-labs/newsletter/docs"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/middleware"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/routers"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/Solidform-labs/newsletter/internal/pkg/fiber_storage"
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
	envFile := ".env"
	if configs.Environment() != "production" {
		envFile = ".env.dev"
	}
	log.Infof("Loading %s file", envFile)
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Load(envFile); err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	config := configs.GetConfig()

	db.Connect()
	db.Init()

	app := fiber.New()
	middleware.Setup(app)
	routers.Setup(app)

	go func() {
		if err := app.Listen(":" + config.ApiPort); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()
	log.Infof("Server is running on port %s", config.ApiPort)

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)                                             // Create a channel to listen for OS signals
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM) // Notify the channel when an interrupt or termination signal is received
	<-sigs                                                                      // Block the main thread until an interrupt or termination signal is received
	log.Info("Shutting down server...")
	app.Shutdown()
	db.GetDB().Close()
	fiber_storage.GetStorage().Close()
	log.Info("Server has been shut down")
}
