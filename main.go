package main

import (
	"fmt"
	"net/http"
	"os"
	"service/src/api"
	"service/src/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server *gin.Engine
)

func init() {
	// Loading Environments
	err := godotenv.Load("app.env")
	if err != nil {
		panic("Error loading environments")
	}
	server = gin.Default()
}

func main() {
	servicePort := os.Getenv("SERVICE_PORT")
	serviceHost := os.Getenv("SERVICE_HOST")

	// Create DSN For DB Connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
	)

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected to DB")

	// AutoMigrate Models
	// sc := database.ShipmentController{}
	// sc.Migrate()
	db.AutoMigrate(&model.ShipmentEntity{})

	// Router Setup
	router := server.Group("/api")
	router.GET("/health-check", func(ctx *gin.Context) {
		msg := "Service started!"
		ctx.JSON(http.StatusOK, gin.H{"status":"success", "content":&msg})
	})
	sr := api.NewShipmentRouter(db)
	sr.ShipmentRoute(router)
	sr.CreateShipment(router)
	server.Run(serviceHost + ":" + servicePort)
}