package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Model User
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// func init() {
// 	errLoad := godotenv.Load()
// 	if errLoad != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")

// 	// Gunakan nilai default jika DB_PORT kosong
// 	if dbPort == "" {
// 		dbPort = "5432"
// 	}

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		dbHost, dbUser, dbPassword, dbName, dbPort)

// 	fmt.Println("DSN:", dsn) // Debugging string koneksi

// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Gagal koneksi ke database: %v", err)
// 	}

// 	if err := db.AutoMigrate(&User{}); err != nil {
// 		log.Fatalf("Gagal melakukan migrasi: %v", err)
// 	}
// }

// Create User
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Get All Users
func getUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func main() {
	// Setup Gin router
	router := gin.Default()

	// Define CRUD Endpoints
	router.POST("/users", createUser)
	router.GET("/users", getUsers)

	// Run the server
	port := ":8080"
	fmt.Println("Server berjalan di port", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Server gagal berjalan:", err)
	}
}
