package main

import (
	"fmt"

	"github.com/fuadsuleyman/go-couriers/database"
	"github.com/fuadsuleyman/go-couriers/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
) 


func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	database.Connect()
	
    app := fiber.New()

    routes.SetupRoutes(app)

    app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}