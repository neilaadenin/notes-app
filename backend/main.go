// @title Notes App API
// @version 1.0
// @description API untuk aplikasi Notes
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"backend/database"
	"backend/routes"
	"time"

	docs "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Configure CORS 
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	config.ExposeHeaders = []string{"Content-Length", "Content-Type"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	database.Connect()

	docs.SwaggerInfo.BasePath = "/api"

	// Swagger 
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Handle OPTIONS preflight requests for all API routes
	r.OPTIONS("/api/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204)
	})

	routes.AuthRoutes(r)
	routes.NotesRoutes(r)

	r.Run(":8080")
}

