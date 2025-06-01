package main

import (
	"go-async-orders/config"
	models "go-async-orders/model"
	"log"
)

func main() {

	config.ConnectDatabase()

	// AutoMigrate
	err := config.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated.")

}