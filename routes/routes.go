package routes

import (
	"companies_handling/handlers"
	"companies_handling/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", userHandler.LoginUser)
	router.GET("/companies/:id", nil)

	userGroup := router.Group("/users")
	// @todo : Here you will check the user and restrict the access via middlware
	userGroup.Use(middlewares.JwtCheck)
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
		userGroup.DELETE("/:id/hard", userHandler.DeleteUserHard)

		userGroup.POST("/companies", nil)
		userGroup.PATCH("/companies/:id", nil)
		userGroup.DELETE("/companies/:id", nil)
	}
}
