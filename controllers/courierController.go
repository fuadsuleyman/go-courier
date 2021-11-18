package controllers

import (
	"fmt"
	// "log"
	"github.com/gofiber/fiber/v2"
	"github.com/fuadsuleyman/go-couriers/helper"
)

func GetCouriers(c *fiber.Ctx) error {

	return c.SendString("Get All Couriers")
}

func GetCourier(c *fiber.Ctx) error {
	fmt.Println("Hi, from GetCourier!")
	header := c.Get("Authorization")
	resMap := helper.CheckToken(header)
	if _, ok := resMap["warning"]; ok{
		return c.JSON(fiber.Map{
			"warning": resMap["warning"],
		})
	}
	fmt.Println("resMap: ", resMap)

	return c.SendString("Get Single Courier")
}

func CreateCourier(c *fiber.Ctx) error {
	fmt.Println("Hi, from Create Courier!")
	header := c.Get("Authorization")
	resMap := helper.CheckToken(header)
	if _, ok := resMap["warning"]; ok{
		return c.JSON(fiber.Map{
			"warning": resMap["warning"],
		})
	}
	fmt.Println("resMap: ", resMap)
	return c.SendString("Create Courier")
}

func UpdateCourier(c *fiber.Ctx) error {
	return c.SendString("Update Courier")
}

func DeleteCourier(c *fiber.Ctx) error {
	return c.SendString("Delete Courier")
}

