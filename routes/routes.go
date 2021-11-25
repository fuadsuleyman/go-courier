package routes

import (
	"github.com/fuadsuleyman/go-couriers/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	
	app.Get("/api/v1.0/couriers", controllers.GetCouriers)

	app.Get("/api/v1.0/couriers/:id", controllers.GetCourier)

	app.Post("/api/v1.0/couriers", controllers.CreateCourier)

	app.Put("/api/v1.0/couriers/:id", controllers.UpdateCourier)

	app.Delete("/api/v1.0/couriers/:id", controllers.DeleteCourier)

	// order apis

	app.Get("/api/v1.0/orders", controllers.GetOrders)
	app.Get("/api/v1.0/orders/:id/pick-up", controllers.PickOrder)
	app.Get("/api/v1.0/orders/:id/deliver", controllers.DeliverOrder)

	// proxy

	app.Get("/api/v1.0/proxy", controllers.MyProxy)

	

}

