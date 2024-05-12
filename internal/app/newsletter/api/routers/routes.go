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

	//unprotected routes
	subs.Delete("/:id", controllers.DeleteSubscriber)

	newsletter.Get("/auth", controllers.AuthenticateAndSendToken)

	app.Get("/swagger/*", swagger.HandlerDefault)

	//protected routes
	app.Use(middleware.CheckAuth)
	subs.Post("/", controllers.AddSubscriber)
	// subs.Get("/", controllers.ListSubscribers)
	// subs.Get("/:id", controllers.GetSubscriber)
	subs.Post("/send", controllers.SendEmailToSubscribers)
}
