package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Subscribe(c *fiber.Ctx) error {
	return c.SendString("Subscribe")
}

func Unsubscribe(c *fiber.Ctx) error {
	return c.SendString("Unsubscribe")
}
