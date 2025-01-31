package controllers

import (
	"chukotka/models"
	"chukotka/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTerms - возвращает список всех терминов
func GetTerms(c *gin.Context) {
	terms, err := repositories.GetAllTerms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch terms"})
		return
	}
	c.JSON(http.StatusOK, terms)
}

// GetTermByID - возвращает термин по ID
func GetTermByID(c *gin.Context) {
	id := c.Param("id")
	term, err := repositories.GetTermByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch term"})
		return
	}
	if term == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Term not found"})
		return
	}
	c.JSON(http.StatusOK, term)
}

// CreateTerm - создает новый термин
func CreateTerm(c *gin.Context) {
	var term models.Term
	if err := c.ShouldBindJSON(&term); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateTerm(&term); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create term"})
		return
	}
	c.JSON(http.StatusCreated, term)
}

// UpdateTerm - обновляет термин по ID
func UpdateTerm(c *gin.Context) {
	id := c.Param("id")
	term, err := repositories.GetTermByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch term"})
		return
	}
	if term == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Term not found"})
		return
	}
	if err := c.ShouldBindJSON(term); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.UpdateTerm(term); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update term"})
		return
	}
	c.JSON(http.StatusOK, term)
}

// DeleteTerm - удаляет термин по ID
func DeleteTerm(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteTerm(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete term"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Term deleted successfully"})
}