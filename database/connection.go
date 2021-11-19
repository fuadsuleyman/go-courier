package database

import (
	"fmt"
	"log"

	"github.com/fuadsuleyman/go-couriers/models"

	// "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
// auth_db_1
func Connect() {

	// db, err := gorm.Open("postgres", "host=192.168.31.74  user=lezzetly password=lezzetly123 dbname=db_name port=5432 sslmode=disable Timezone=Asia/Baku")

	// if err != nil {
	// 	fmt.Println(err, "Error is  here")
	// 	log.Println("Connection Failed to Open")
	// } else {
	// 	log.Println("Connection Established")
	// }
	// DB = db

	x := gorm.Open
	connection, err := x(postgres.Open("host=192.168.31.74  user=lezzetly password=lezzetly123 dbname=db_name port=5432 sslmode=disable Timezone=Asia/Baku"))

	
	// connection, err := x(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
	// viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), viper.GetString("db.dbname"), viper.GetString("db.password"), viper.GetString("db.sslmode"), viper.GetString("db.TimeZone"))), &gorm.Config{})

	if err != nil {
		fmt.Println(err, "Error is  here")
		panic("could not connect to the database!")
	}else {
		log.Println("Connection Established")
	}

	DB = connection

	DB.AutoMigrate(&models.Courier{})

}

