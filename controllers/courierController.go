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
func GetCouriers(c *fiber.Ctx) error {

	var couriers []models.Courier

	database.DB.Find(&couriers)

	if len(couriers) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "Don't find any couriers!",
		})
	}

	couriersMessage := fmt.Sprintf("couriers(%v)", len(couriers))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		couriersMessage: couriers,
	})

}

func GetCourier(c *fiber.Ctx) error {
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
			"warning": "You have not permission to update courier!",
		})
	}

	// read id from param
	id := c.Params("id")

	var courier models.Courier

	database.DB.Find(&courier, "id = ?", id)

	notExistsMesssage := fmt.Sprintf("Courier with id %v is not exists!", id)

	if courier.Id == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsMesssage,
		})
	}

	// You can not get other couriers data!
	if courier.Username != resMap["Username"] {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You can not get other couriers data!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"courier": courier,
	})

	// fmt.Println("courierin id-si", courier.Id)

	// paramId := fmt.Sprintf("Get Single Courier with id: %v", id)

	// return c.SendString(paramId)
}

func CreateCourier(c *fiber.Ctx) error {
	fmt.Println("Hi, from Create Courier!")
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
			"warning": "You have not permission to create courier!",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var existsCourier models.Courier

	// check if entered username exists in db or not
	database.DB.Where("username = ?", resMap["Username"]).First(&existsCourier)

	if existsCourier.Id > 0 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"warning": "This cook with this username is alredy exits!",
		})
	}

	courier := models.Courier{
		Username:  resMap["Username"],
		Usertype:  resMap["Usertype"],
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
	}

	errors := helper.ValidateStruct(courier)

	if errors != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": errors,
		})
	}

	// create user
	createVal := database.DB.Create(&courier)
	fmt.Println("createVal:", createVal.Error)

	if createVal.Error != nil {
		return c.JSON(fiber.Map{
			"warning": createVal.Error.Error(),
		})	
	}

	responseMessage := fmt.Sprintf("Courier with id %v is successfully created!", courier.Id)
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": responseMessage,
	})

}

func UpdateCourier(c *fiber.Ctx) error {

	fmt.Println("Hi, from UpdateCourier!")
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
			"warning": "You have not permission to update courier!",
		})
	}

	var courier models.Courier

	// Read the param noteId
	id := c.Params("id")

	// Find the courier with the given Id
	database.DB.Find(&courier, "id = ?", id)

	notExistsMesssage := fmt.Sprintf("Courier with id %v is not exists!", id)

	// You can not get other couriers data!
	if courier.Username != resMap["Username"] {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": "You can not update other couriers data!",
		})
	}

	if courier.Id == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"warning": notExistsMesssage,
		})
	}

	type updateCourier struct {
		Firstname      string `json:"first_name"`
		Lastname       string `json:"last_name"`
		Patronymic     string `json:"patronymic"`
		Email          string `json:"email"`
		Phone          string `json:"phone"`
		Transport      string `json:"transport"`
		WorkExperience int    `json:"work_experience"`
		IsAvailable    bool   `json:"is_available"`
		Location       string `json:"location"`
	}

	// Store the body containing the updated data and return error if encountered
    var updateCourierData updateCourier
    err := c.BodyParser(&updateCourierData)

	if err != nil {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"warning": "Review your input", "data": err})
    }

	// Edit the courier
    courier.Firstname = updateCourierData.Firstname
    courier.Lastname = updateCourierData.Lastname
    courier.Patronymic = updateCourierData.Patronymic
	courier.Phone = updateCourierData.Phone
	courier.Email = updateCourierData.Email
	courier.Transport = updateCourierData.Transport
	courier.WorkExperience = updateCourierData.WorkExperience
	courier.IsAvailable = updateCourierData.IsAvailable
	courier.Location = updateCourierData.Location

    // Save the Changes
    saveVal := database.DB.Save(&courier)
	fmt.Println("SaveVal in update: ", saveVal)

	if saveVal.Error != nil {
		return c.JSON(fiber.Map{
			"warning": saveVal.Error.Error(),
		})	
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Seccessfully Updated": courier})

}

func DeleteCourier(c *fiber.Ctx) error {
	return c.SendString("Delete Courier")
}
