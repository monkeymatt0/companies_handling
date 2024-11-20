package routes

import (
	"companies_handling/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	router.POST("/user", userHandler.CreateUser)
	router.GET("/user/:id", userHandler.GetUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)
	router.DELETE("/user/:id/hard", userHandler.DeleteUserHard)
	router.POST("/login", userHandler.LoginUser)

	router.GET("/companies/:id", nil)
	userGroup := router.Group("/user")
	// @todo : Here you will check the user and restrict the access via middlware
	userGroup.Use(nil)
	{
		router.POST("/companies", nil)
		router.PATCH("/companies/:id", nil)
		router.DELETE("/companies/:id", nil)
	}
}
