package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	serviceHost *string
	servicePort *string
)

func init() {
	// config, err := initializers.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Error with loading enironments!", err)
	// }
	// initializers.ConnectDB(&config)
	server = gin.Default()
}

func main() {
	// config, err := initializers.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Error with loading environments!", err)
		
	// }

	servicePort := os.Getenv("PORT")
	serviceHost := os.Getenv("HOST")

	// Router Setup
	router := server.Group("/api")
	router.GET("/health-check", func(ctx *gin.Context) {
		msg := "Service started!"
		ctx.JSON(http.StatusOK, gin.H{"status":"succes", "content":&msg})
	})
	server.Run(serviceHost + ":" + servicePort)
}