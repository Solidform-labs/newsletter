package controllers

import (
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/Solidform-labs/newsletter/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

func AddSubscriber(c *fiber.Ctx) error {
	sub := new(models.Subscriber)

	if err := c.BodyParser(sub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}

	if !validation.IsValidEmail(sub.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email",
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

func DeleteSubscriber(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	if validation.IsValidEmail(id) {
		if err := db.DeleteSubscriberByEmail(id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not unsubscribe",
				"error":   err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
	if isNumericID, intID := validation.ParseNumericID(id); isNumericID {
		if err := db.DeleteSubscriberByID(intID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not unsubscribe",
				"error":   err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "id is invalid",
	})

}
