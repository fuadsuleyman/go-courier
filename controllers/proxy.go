package controllers

import (
	"fmt"
	// "strings"

	// "log"
	// "github.com/fuadsuleyman/go-couriers/database"
	// "github.com/fuadsuleyman/go-couriers/helper"
	// "github.com/fuadsuleyman/go-couriers/models"
	// "net/http"
	// "net/http/httputil"
	"net/url"

	"github.com/gofiber/fiber/v2"
	// "github.com/sirupsen/logrus"
	// "bytes"
	// "io/ioutil"
	// "github.com/google/uuid"
)

// I don't need this api
func MyProxy(c *fiber.Ctx) error {

	fmt.Println("port: ", c.Port())
	fmt.Println("BaseURL: ", c.BaseURL())
	fmt.Println("Hostname: ", c.Hostname())
	fmt.Println("Path: ", c.Path())

	origin, _ := url.Parse("http://192.168.31.74:8005/api/v1.0/proxy")
	fmt.Println("origin:", origin)
	baseUrl := c.BaseURL()
	myLen := len(baseUrl) - 4
	port := baseUrl[myLen:]
	path := c.Path()
	
	fmt.Println("custom port: ", port)
	fmt.Println("path: ", path)

	

	

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "look at terminal",
	})
}

// func readBody(response *http.Response) (uuid.UUID, string) {
// 	defer response.Body.Close()
// 	all, _ := ioutil.ReadAll(response.Body)
// 	u := uuid.New()
// 	var s string
// 	if len(all) > 0 {
// 		s = string(all)
// 	}
// 	return u, s
// }