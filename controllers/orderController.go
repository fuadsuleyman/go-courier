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
	fmt.Println("Ertered to get orders")

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


// I don't need this api
func GetOrder(c *fiber.Ctx) error {

	var order models.Order

	id := c.Params("id")

	database.DB.Find(&order, "id = ?", id)


	notExistsMesssage := fmt.Sprintf("Order with id %v is not exists!", id)

	if order.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsMesssage,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"order": order,
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
	// id := c.Params("id")

	// fmt.Println("Parametrdeki ID: ", id)

	var order models.Order

	database.DB.Where("courier.username = ? AND is_active = ?", resMap["Username"], true).Find(&order)
	// database.DB.Find(&order, "courier.username = ?", resMap["Username"])

	fmt.Println("Tapdigim orderin id-si:", order.ID)

	notExistsOrderMesssage := fmt.Sprintf("This Courier have not active order to pick up!")

	if order.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsOrderMesssage,
		})
	}

	// check courier
	username := resMap["Username"]

	var currentCourier models.Courier

	database.DB.Find(&currentCourier, "username = ?", username)

	notExistsCourierMesssage := fmt.Sprintf("Courier with username %v is not exists!", username)

	if currentCourier.Id == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsCourierMesssage,
		})
	}
	
	if int64(currentCourier.Id) != order.CourierId {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You have not permission pick up order with this token!",
		})	
	}
	
	if order.IsActive != true {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You can not pick up not active order!",
		})
	}

	order.Status = "courier on the way to client"

	// Save the Changes
    saveVal := database.DB.Save(&order)
	fmt.Println("SaveVal in update: ", saveVal)

	if saveVal.Error != nil {
		return c.JSON(fiber.Map{
			"warning": saveVal.Error.Error(),
		})	
	}

	seccessMes := fmt.Sprintf("Order with %v id is picked by %v", order.ID, username)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": seccessMes,
	})

}


func DeliverOrder(c *fiber.Ctx) error {
	fmt.Println("Hi, from DeliverOrder!")
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

	notExistsOrderMesssage := fmt.Sprintf("Order with id %v is not exists!", id)

	if order.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsOrderMesssage,
		})
	}

	// check courier
	username := resMap["Username"]

	var currentCourier models.Courier

	database.DB.Find(&currentCourier, "username = ?", username)

	notExistsCourierMesssage := fmt.Sprintf("Courier with username %v is not exists!", username)

	if currentCourier.Id == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsCourierMesssage,
		})
	}

	if int64(currentCourier.Id) != order.CourierId {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You have not permission deliver order with this token!",
		})	
	}
	
	if order.IsActive != true {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You can not deliver not active order!",
		})
	}

	if order.Status != "courier on the way to client"{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "First pick order from cook!",
		})
	}

	order.Status = "order delivered"

	// Save the Changes
    saveVal := database.DB.Save(&order)
	fmt.Println("SaveVal in update: ", saveVal)

	if saveVal.Error != nil {
		return c.JSON(fiber.Map{
			"warning": saveVal.Error.Error(),
		})	
	}

	seccessMes := fmt.Sprintf("Order with %v id is delived to client by %v", order.ID, username)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": seccessMes,
	})

}