package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite" // Use GORM's SQLite driver
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Subscription represents the subscription model
type Subscription struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	Plan   string `json:"plan"`
	Active bool   `json:"active"`
}

// Database instance
var db *gorm.DB

// Initialize the database connection
func initDB() {
	var err error

	// Open a new SQLite database connection
	db, err = gorm.Open(sqlite.Open("file::memory:?mode=memory&cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Use singular table names
		},
	})

	if err != nil {
		panic("failed to connect to the database" + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&Subscription{})

	// Insert sample data
	db.Create(&Subscription{UserID: 1, Plan: "Basic", Active: true})
	db.Create(&Subscription{UserID: 2, Plan: "Premium", Active: false})
}

// Get all subscriptions
func getSubscriptions(c *gin.Context) {
	var subscriptions []Subscription
	if err := db.Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

func main() {
	fmt.Println("Initializing DB")
	// Initialize the database
	initDB()

	// Set up the Gin router
	router := gin.Default()

	// Define the endpoint to get all subscriptions
	router.GET("/subscriptions", getSubscriptions)

	fmt.Println("Starting Server")
	fmt.Println("===========================")
	// Start the server
	router.Run(":8080")
}
