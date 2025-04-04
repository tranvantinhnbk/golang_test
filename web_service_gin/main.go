package main

import (
	_ "example/web-service-gin/docs"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"example/web-service-gin/router"
)

// @title Example Web Service API
// @version 1.0
// @description A simple API using Gin and Swagger
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	// Swagger route
	r.GET("/docs", func(c *gin.Context) {
		// Redirect to Swagger UI directly without 'index.html'
		c.Redirect(302, "/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	router.SetupRoutes(r)

	r.Run(":8080")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// Graceful shutdown with a timeout
	fmt.Println("\nShutting down the server...")

	// Add a timeout to allow pending requests to complete
	time.Sleep(2 * time.Second)

	// Now the server will be stopped gracefully
	fmt.Println("Server stopped.")
}
