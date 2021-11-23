package controllers


import (
	"fmt"
	// "log"
	"github.com/fuadsuleyman/go-couriers/database"
	// "github.com/fuadsuleyman/go-couriers/helper"
	"github.com/fuadsuleyman/go-couriers/models"
	"github.com/gofiber/fiber/v2"
)

// I don't need this api
func GetOrders(c *fiber.Ctx) error {

	var orders []models.Order

	database.DB.Find(&orders)

	if len(orders) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "Don't find any orders!",
		})
	}

	ordersMessage := fmt.Sprintf("orders(%v)", len(orders))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		ordersMessage: orders,
	})

}