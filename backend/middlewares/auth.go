package middlewares

import (
	"chukotka/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware - проверяет токен администратора
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение токена из заголовка
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			c.Abort()
			return
		}

		// Проверка токена
		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
