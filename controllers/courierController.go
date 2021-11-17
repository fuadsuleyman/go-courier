package controllers

import "github.com/gofiber/fiber/v2"

func GetCouriers(c *fiber.Ctx) error {
	return c.SendString("Get All Couriers")
}

func GetCourier(c *fiber.Ctx) error {
	return c.SendString("Get Single Courier")
}

func CreateCourier(c *fiber.Ctx) error {
	return c.SendString("Create Courier")
}

func UpdateCourier(c *fiber.Ctx) error {
	return c.SendString("Update Courier")
}

func DeleteCourier(c *fiber.Ctx) error {
	return c.SendString("Delete Courier")
}

