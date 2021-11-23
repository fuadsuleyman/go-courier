package controllers


import (
	"fmt"
	// "log"
	"github.com/fuadsuleyman/go-couriers/database"
	"github.com/fuadsuleyman/go-couriers/helper"
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

func PickOrder(c *fiber.Ctx) error {
	fmt.Println("Hi, from GetCourier!")
	header := c.Get("Authorization")
	resMap := helper.CheckToken(header)

	if _, ok := resMap["warning"]; ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": resMap["warning"],
		})
	}
	fmt.Println("resMap: ", resMap)

	if resMap["Usertype"] != "2" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "Only Courier have not permission to pick up order!",
		})
	}

	// check order with param
	id := c.Params("id")

	fmt.Println("Parametrdeki ID: ", id)

	var order models.Order

	database.DB.Find(&order, "id = ?", id)

	fmt.Println("Tapdigim orderin id-si:", order.ID)

	notExistsMesssage := fmt.Sprintf("Order with id %v is not exists!", id)

	if order.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsMesssage,
		})
	}

	// check courier
	username := resMap["Username"]

	var currentCourier models.Courier

	database.DB.Find(&currentCourier, "username = ?", username)
	// currentCourier deyilse error
	fmt.Println("currentCourier: ", currentCourier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"warning": "terminala bax",
	})

}