package controllers


import (
	"fmt"
	// "log"
	// "github.com/fuadsuleyman/go-couriers/database"
	// "github.com/fuadsuleyman/go-couriers/helper"
	// "github.com/fuadsuleyman/go-couriers/models"
	"github.com/gofiber/fiber/v2"
)

// I don't need this api
func MyProxy(c *fiber.Ctx) error {

	fmt.Println("port: ", c.Port())
	fmt.Println("BaseURL: ", c.BaseURL())
	fmt.Println("Hostname: ", c.Hostname())
	fmt.Println("Path: ", c.Path())

	 
	baseUrl := c.BaseURL()
	myLen := len(baseUrl) - 4
	port := baseUrl[myLen:]
	
	fmt.Println("custom port: ", port)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "look at terminal",
	})
}