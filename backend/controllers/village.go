package controllers

import (
	"chukotka/models"
	"chukotka/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVillages - возвращает список всех деревень
func GetVillages(c *gin.Context) {
	villages, err := repositories.GetAllVillages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch villages"})
		return
	}
	c.JSON(http.StatusOK, villages)
}

// GetVillageByID - возвращает деревню по ID
func GetVillageByID(c *gin.Context) {
	id := c.Param("id")
	village, err := repositories.GetVillageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch village"})
		return
	}
	if village == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Village not found"})
		return
	}
	c.JSON(http.StatusOK, village)
}

// CreateVillage - создает новую деревню
func CreateVillage(c *gin.Context) {
	var village models.Village
	if err := c.ShouldBindJSON(&village); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateVillage(&village); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create village"})
		return
	}
	c.JSON(http.StatusCreated, village)
}

// UpdateVillage - обновляет деревню по ID
func UpdateVillage(c *gin.Context) {
	id := c.Param("id")
	village, err := repositories.GetVillageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch village"})
		return
	}
	if village == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Village not found"})
		return
	}
	if err := c.ShouldBindJSON(village); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.UpdateVillage(village); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update village"})
		return
	}
	c.JSON(http.StatusOK, village)
}

// DeleteVillage - удаляет деревню по ID
func DeleteVillage(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteVillage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete village"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Village deleted successfully"})
}
