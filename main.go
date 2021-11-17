package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/fuadsuleyman/go-couriers/controllers"
) 

func Setup(app *fiber.App) {
	app.Get("/api/v1.0/couriers", controllers.GetCouriers)
}

func main() {
    app := fiber.New()

    Setup(app)

    app.Listen(":3000")
}