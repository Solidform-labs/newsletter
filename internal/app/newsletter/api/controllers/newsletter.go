package controllers

import (
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/Solidform-labs/newsletter/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

// AddSubscriber godoc
// @Summary Add a new subscriber
// @Description Add a new subscriber to the newsletter
// @Tags subscribers
// @Accept json
// @Produce json
// @Param subscriber body models.Subscriber true "Subscriber object"
// @Success 201
// @Failure 400 {object} models.BaseError "Bad Request Error message"
// @Failure 500 {object} models.BaseError "Internal Error message"
// @Router /newsletter/subscribers [post]
func AddSubscriber(c *fiber.Ctx) error {
	sub := new(models.Subscriber)

	if err := c.BodyParser(sub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BaseError{
			Message: "Invalid JSON format",
			Error:   err.Error(),
		})
	}

	if !validation.IsValidEmail(sub.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(models.BaseError{
			Message: "Invalid email",
			Error:   "Email is not valid",
		})
	}

	if err := db.CreateSubscriber(sub.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.BaseError{
			Message: "Could not subscribe",
			Error:   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DeleteSubscriber godoc
// @Summary Delete a subscriber
// @Description Delete a subscriber from the newsletter
// @Tags subscribers
// @Param id path string true "Subscriber ID or email"
// @Success 204
// @Failure 400 {object} models.BaseError "Bad Request Error message"
// @Failure 500 {object} models.BaseError " Internal Error message"
// @Router /newsletter/subscribers/{id} [delete]
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
