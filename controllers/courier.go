package controllers

import "github.com/gofiber/fiber/v2"

func GetCouriers(c *fiber.Ctx) error {
	return c.SendString("Get All Couriers")
}