package controllers

import (
	"chukotka/repositories"
	"chukotka/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAboutPage(c *gin.Context) {
	about, err := repositories.GetAboutPage()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "About page not found"})
		return
	}
	c.JSON(http.StatusOK, about)
}

func UpdateAboutPage(c *gin.Context) {
	var request models.AboutPage

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateAboutPage(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update about page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "About page updated successfully"})
}