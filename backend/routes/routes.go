package routes

import (
	"chukotka/controllers"
	"chukotka/middlewares"

	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	registerPublicRoutes(api)
	registerAdminRoutes(api)
}

func registerPublicRoutes(api *gin.RouterGroup) {
	api.GET("/districts", controllers.GetDistricts)
	api.GET("/districts/:id", controllers.GetDistrictByID)

	api.GET("/villages", controllers.GetVillages)
	api.GET("/villages/:id", controllers.GetVillageByID)

	api.GET("/terms", controllers.GetTerms)
	api.GET("/terms/:id", controllers.GetTermByID)

	api.POST("/register", controllers.RegisterAdmin)
	api.POST("/login", controllers.LoginAdmin)

	api.GET("/is-admin", controllers.ValidateTokenBool)
}

func registerAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middlewares.AdminAuthMiddleware())

	admin.POST("/districts", controllers.CreateDistrict)
	admin.PUT("/districts/:id", controllers.UpdateDistrict)
	admin.DELETE("/districts/:id", controllers.DeleteDistrict)

	admin.POST("/villages", controllers.CreateVillage)
	admin.PUT("/villages/:id", controllers.UpdateVillage)
	admin.DELETE("/villages/:id", controllers.DeleteVillage)
	
	admin.POST("/terms", controllers.CreateTerm)
	admin.PUT("/terms/:id", controllers.UpdateTerm)
	admin.DELETE("/terms/:id", controllers.DeleteTerm)

	admin.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
	})
}

