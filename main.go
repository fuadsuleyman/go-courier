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
	"github.com/gin-gonic/gin"
) 


func main() {


	// start gin part
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	// end gin part
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	database.Connect()
	
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(logger.New())

    routes.SetupRoutes(app)

    app.Listen(fmt.Sprintf(":%s", viper.GetString("port")))
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}