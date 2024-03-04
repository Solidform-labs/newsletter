package controllers

import (
	"github.com/Solidform-labs/newsletter/internal/newsletter/api/models"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Subscribe(c *fiber.Ctx) error {
	sub := new(models.Subscriber)

	if err := c.BodyParser(sub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(sub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email",
			"error":   err.Error(),
		})
	}

	if err := db.CreateSubscriber(sub.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not subscribe",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func Unsubscribe(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email is required",
		})
	}

	sub := models.Subscriber{
		Email: email,
	}

	validate := validator.New()
	if err := validate.Struct(sub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email",
			"error":   err.Error(),
		})
	}

	if err := db.DeleteSubscriber(email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not unsubscribe",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
