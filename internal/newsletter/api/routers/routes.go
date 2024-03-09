package routers

import (
	"github.com/Solidform-labs/newsletter/internal/newsletter/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/v1")

	newsletter := v1.Group("/newsletter")

	subs := newsletter.Group("/subscribers")
	// subs.Get("/", controllers.ListSubscribers)
	subs.Post("/", controllers.AddSubscriber)
	// subs.Get("/:id", controllers.GetSubscriber)
	subs.Delete("/:id", controllers.DeleteSubscriber)
}
