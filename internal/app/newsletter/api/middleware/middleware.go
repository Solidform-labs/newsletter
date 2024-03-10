package middleware

import (
	"github.com/Solidform-labs/newsletter/configs"
	"github.com/Solidform-labs/newsletter/internal/pkg/fiber_storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	storage := fiber_storage.Create(config)
	app.Use(limiter.New(limiter.Config{
		Storage:    storage,
		Max:        config.ApiMaxRequests,
		Expiration: config.ApiRequestsExpiration,
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
	}))
}
