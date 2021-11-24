package controllers

import (
	"fmt"
	"strings"
	// "log"
	// "github.com/fuadsuleyman/go-couriers/database"
	// "github.com/fuadsuleyman/go-couriers/helper"
	// "github.com/fuadsuleyman/go-couriers/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gin-gonic/gin"
)

func ReverseProxy(ctx *gin.Context) {
	fmt.Println("I am from gin")
}


// I don't need this api
func MyProxy(c *fiber.Ctx) error {

	fmt.Println("port: ", c.Port())
	fmt.Println("BaseURL: ", c.BaseURL())
	fmt.Println("Hostname: ", c.Hostname())
	fmt.Println("Path: ", c.Path())

	 
	baseUrl := c.BaseURL()
	myLen := len(baseUrl) - 4
	port := baseUrl[myLen:]
	path := c.Path()
	
	fmt.Println("custom port: ", port)

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	fmt.Println("parts", parts)

	targetHost := fmt.Sprintf("svc-%s", parts[1])
	targetNamespace := fmt.Sprintf("svc-%s", parts[2])

	targetAddr := fmt.Sprintf(
		"http://%s.%s:%d/api/%s",
		targetHost, targetNamespace, 10000, strings.Join(parts[3:], "/"),
	)

	fmt.Println("targetHost:", targetHost)
	fmt.Println("targetNamespace:", targetNamespace)
	fmt.Println("targetAddr:", targetAddr)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "look at terminal",
	})
}