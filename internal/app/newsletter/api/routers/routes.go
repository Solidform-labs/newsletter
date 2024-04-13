package routers

import (
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/controllers"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/v1")

	newsletter := v1.Group("/newsletter")

	subs := newsletter.Group("/subscribers")

	app.Use(middleware.CheckAuth)

	//unprotected routes
	subs.Post("/", controllers.AddSubscriber)
	subs.Delete("/:id", controllers.DeleteSubscriber)
	app.Get("/swagger/*", swagger.HandlerDefault)

	//protected routes
	// subs.Get("/", controllers.ListSubscribers)
	// subs.Get("/:id", controllers.GetSubscriber)
	subs.Post("/send", controllers.SendEmailToSubscribers)
	subs.Post("/send/:id", controllers.SendEmailToSubscriber)
}
