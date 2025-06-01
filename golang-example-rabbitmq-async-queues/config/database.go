package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){

	database, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=orders_with_rabbit_mq port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}