package main

import (
	"fmt"

	"github.com/fuadsuleyman/go-couriers/database"
	"github.com/fuadsuleyman/go-couriers/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/redirect/v2"
) 


func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	database.Connect()
	
    app := fiber.New()


	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
		  "/proxy":   "http://192.168.31.74/api/v1.0/meals/",
		  "/old/*": "/new/$1",
		},
		StatusCode: 301,
	  }))

    routes.SetupRoutes(app)

    app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}