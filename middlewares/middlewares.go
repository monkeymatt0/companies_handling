package middlewares

import (
	"companies_handling/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtCheck(c *gin.Context) {
	tokenString := c.GetHeader("Cookie")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization token"})
		c.Abort()
		return
	}
	tokenString = tokenString[7:] // Removing Bearer=
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.Claims{
			Email: "matteo@test.com",
		},
		func(t *jwt.Token) (interface{}, error) {
			return os.Getenv("SECRET"), nil
		},
	)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication error"})
		c.Abort()
		return
	}

	c.Next()
}
