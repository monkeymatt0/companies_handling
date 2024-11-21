package middlewares

import (
	"companies_handling/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtCheck(c *gin.Context) {
	tokenString, err := c.Cookie("Bearer")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization error"})
		c.Abort()
		return
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization token"})
		c.Abort()
		return
	}
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		},
	)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication error"})
		c.Abort()
		return
	}

	c.Next()
}
