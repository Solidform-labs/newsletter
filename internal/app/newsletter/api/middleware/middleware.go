package middleware

import (
	"github.com/Solidform-labs/newsletter/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/storage/postgres/v2"
)

func Setup(app *fiber.App) {
	config := configs.GetConfig()

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.Environment == "development"
		},
	}))

	app.Use(helmet.New())

	app.Use(logger.New())

	app.Use(healthcheck.New())

	storage := postgres.New(postgres.Config{
		ConnectionURI: config.DbConnectionString,
	})
	app.Use(limiter.New(limiter.Config{
		Storage: storage,
	}))
}
