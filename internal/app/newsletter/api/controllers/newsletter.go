package controllers

import (
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/Solidform-labs/newsletter/internal/pkg/db"
	"github.com/Solidform-labs/newsletter/internal/pkg/email"
	"github.com/Solidform-labs/newsletter/pkg/encryptdecrypt"
	"github.com/Solidform-labs/newsletter/pkg/tokens"
	"github.com/Solidform-labs/newsletter/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Add missing import statement

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
	deleteKey := c.Query("delete_key")
	if deleteKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "delete_key is required",
		})
	}

	if validation.IsValidEmail(id) {
		if err := db.DeleteSubscriberByEmail(id, deleteKey); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not unsubscribe",
				"error":   err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
	if isNumericID, intID := validation.ParseNumericID(id); isNumericID {
		if err := db.DeleteSubscriberByID(intID, deleteKey); err != nil {
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

func SendEmailToSubscribers(c *fiber.Ctx) error {
	ids := []string{}
	if err := c.BodyParser(ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BaseError{
			Message: "The body does not contain an array of users",
			Error:   err.Error(),
		})
	}

	if len(ids) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "the list of users passed to the function is empty",
		})
	}

	subscribers := []models.Subscriber{}

	for i, id := range ids {
		var subscriberPosition = &subscribers[i]

		if validation.IsValidEmail(id) {
			if err := db.GetSubscriberByEmail(id, subscriberPosition); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "No subscriber found with given Email",
					"error":   err.Error(),
				})
			}
		}

		if isNumericID, intID := validation.ParseNumericID(id); isNumericID {
			if err := db.GetSubscriberByid(intID, subscriberPosition); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "No subscriber found with given id",
					"error":   err.Error(),
				})
			}
		}
	}

	if err := email.SendNewsletter(subscribers, "Newsletter", "This is a test newsletter"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error sending emal to subscriber",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON("Email sent to subscriber")
}

func AuthenticateAndSendToken(c *fiber.Ctx) error {
	reqBody := new(models.User)

	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BaseError{
			Message: "Make sure to pass your email and password",
			Error:   err.Error(),
		})
	}

	var foundUser models.User

	if err := db.GetUserByEmail(reqBody.Email, &foundUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.BaseError{
			Message: "Could not find user",
			Error:   err.Error(),
		})
	}

	if !encryptdecrypt.CheckPassword(foundUser.Password, reqBody.Password) {
		log.Warn("Password did not match")
		return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
	}

	token, err := tokens.CreateToken(reqBody.Email)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
