package controllers

import (
	"chukotka/models"
	"chukotka/repositories"

	"net/http"
	"github.com/gin-gonic/gin"
)

// GetDistricts возвращает список всех районов
func GetDistricts(c *gin.Context) {
	districts, err := repositories.GetAllDistricts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch districts"})
		return
	}
	c.JSON(http.StatusOK, districts)
}

// GetDistrictByID возвращает район по ID
func GetDistrictByID(c *gin.Context) {
	id := c.Param("id")

	// Вызываем функцию репозитория
	district, err := repositories.GetDistrictByID(id)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve district"})
		return
	}

	// Проверяем, существует ли район
	if district == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	// Возвращаем успешный результат
	c.JSON(http.StatusOK, district)
}

// CreateDistrict создает новый район
func CreateDistrict(c *gin.Context) {
	var district models.District
	if err := c.ShouldBindJSON(&district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateDistrict(&district); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create district"})
		return
	}
	c.JSON(http.StatusCreated, district)
}

// UpdateDistrict обновляет данные района
func UpdateDistrict(c *gin.Context) {
	id := c.Param("id")

	district, err := repositories.GetDistrictByID(id)
	
	// Найти существующий район
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "District not found"})
		return
	}

	// Обновить данные
	if err := c.ShouldBindJSON(district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.UpdateDistrict(id, district); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update district"})
		return
	}
	c.JSON(http.StatusOK, district)
}

// DeleteDistrict удаляет район
func DeleteDistrict(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteDistrict(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete district"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "District deleted successfully"})
}