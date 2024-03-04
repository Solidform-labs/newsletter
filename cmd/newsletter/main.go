package main

import (
	"github.com/Solidform-labs/newsletter/configs"
	"github.com/Solidform-labs/newsletter/internal/newsletter/api/routers"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := configs.GetConfig()

	defer db.GetDB().Close()

	app := fiber.New()
	defer app.Shutdown()

	routers.Setup(app)

	log.Infof("Server is running on port %s", config.ApiPort)
	app.Listen(":" + config.ApiPort)
}
