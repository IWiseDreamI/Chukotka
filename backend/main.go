package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"chukotka/routes"
	"chukotka/core/db"
	"chukotka/core/config"
)

func main() {
	config.Load()
	db.InitDatabase()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
	    c.Writer.Header().Set("Content-Encoding", "identity")
	    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	    c.Writer.Header().Set("Access-Control-Max-Age", "86400") // Cache preflight response
	    if c.Request.Method == "OPTIONS" {
	        c.AbortWithStatus(204) // Respond with empty 204 for preflight requests
	        return
	   }
    	   c.Next()
	})

	routes.RegisterRoutes(r)

	port := config.GetEnv("PORT", "8080")
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
