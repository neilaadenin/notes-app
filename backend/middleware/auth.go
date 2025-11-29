package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/models"
	"backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var user models.User
		database.DB.Where("email = ?", email).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", user.ID)
		c.Next()
	}
}

