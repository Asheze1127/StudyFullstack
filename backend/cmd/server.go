package main

import (
	"log"
	"os"
	
	// _ "backend/docs"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Fullstack Study API
// @version         1.0
// @description     認証機能を含むAPI
// @host      localhost:8088
// @BasePath  /v1
func main() {
	r, err := Bootstrap()
	if err != nil {
		log.Fatal("Failed to bootstrap:", err)
	}
	
	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}