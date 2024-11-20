package routes

import "github.com/gin-gonic/gin"

func SetUpRoutes(router *gin.Engine) {
	router.POST("/user", nil)
	router.GET("/user/:id", nil)
	router.DELETE("/user/:id", nil)
	userGroup := router.Group("/user")
	// @todo : Here you will check the user and restrict the access via middlware
	userGroup.Use(nil)
	{
		router.POST("/companies", nil)
		router.PATCH("/companies/:id", nil)
		router.GET("/companies/:id", nil)
		router.DELETE("/companies/:id", nil)
	}
}
