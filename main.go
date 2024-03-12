package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"tugas2-hacktiv8/database"
	"tugas2-hacktiv8/handlers"
	"tugas2-hacktiv8/models"
)

var db *gorm.DB

func main() {
	var err error
	db, err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	if err := pingDB(); err != nil {
		log.Fatal(err)
	}

	// Automigrate models
	migrate()

	// Initialize Gin router
	router := gin.Default()

	// Routes for orders
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders", handlers.GetOrders)
	router.GET("/orders/:id", handlers.GetOrder)
	router.PUT("/orders/:id", handlers.UpdateOrder)
	router.DELETE("/orders/:id", handlers.DeleteOrder)

	// Run the server
	log.Println("Server started on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func pingDB() error {
	// Query database to ensure connection
	var version string
	if err := db.Raw("SELECT version()").Scan(&version).Error; err != nil {
		return err
	}

	log.Printf("Connected to database. Database version: %s\n", version)
	return nil
}

func migrate() {
	err := db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		log.Fatal(err)
	}
}
