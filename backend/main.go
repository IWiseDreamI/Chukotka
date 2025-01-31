package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"chukotka/routes"
	"chukotka/core/db"
	"chukotka/core/config"
	"chukotka/middlewares"
)

func main() {
	config.Load()
	db.InitDatabase()

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	routes.RegisterRoutes(r)

	port := config.GetEnv("PORT", "8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
