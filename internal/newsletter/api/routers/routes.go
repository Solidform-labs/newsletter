package routers

import (
	"github.com/Solidform-labs/newsletter/internal/newsletter/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/v1")
	newsletter := v1.Group("/newsletter")
	newsletter.Post("/subscribe", controllers.Subscribe)
	newsletter.Delete("/unsubscribe", controllers.Unsubscribe)
}
