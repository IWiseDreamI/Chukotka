package controllers

import (
	"chukotka/models"
	"chukotka/repositories"
	"chukotka/utils"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

func RegisterAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	admin.Password = hashedPassword

	if err := repositories.CreateAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered successfully"})
}

func LoginAdmin(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Поиск администратора по имени пользователя
	admin, err := repositories.FindAdminByUsername(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find admin"})
		return
	}
	if admin == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Проверка пароля
	if !utils.CheckPasswordHash(credentials.Password, admin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(admin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ValidateTokenBool проверяет валидность JWT-токена и возвращает результат в виде bool.
func ValidateTokenBool(c *gin.Context) {
	// Получаем токен из заголовка Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusOK, gin.H{"valid": false})
		return
	}

	// Валидируем токен
	token, err := utils.ValidateJWT(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusOK, gin.H{"valid": false})
		return
	}

	// Проверка структуры claims (необязательно, если нужно только валидность)
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"valid": false})
		return
	}

	// Если все проверки пройдены, токен действителен
	c.JSON(http.StatusOK, gin.H{"valid": true})
}