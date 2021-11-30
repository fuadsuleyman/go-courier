package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	fmt.Println("Ertered to not found")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"warning": "Page Not Found",
	})

}
