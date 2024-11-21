package routes

import (
	"companies_handling/handlers"
	"companies_handling/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	companyHandler *handlers.CompanyHandler,
) {
	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", userHandler.LoginUser)
	router.GET("/companies/:uuid", companyHandler.GetCompany)

	userGroup := router.Group("/users")
	userGroup.Use(middlewares.JwtCheck)
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
		userGroup.DELETE("/:id/hard", userHandler.DeleteUserHard)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.JwtCheck)
	{
		authGroup.POST("/companies", companyHandler.CreateCompany)
		authGroup.PATCH("/companies/:uuid", companyHandler.EditCompany)
		authGroup.DELETE("/companies/:uuid", companyHandler.DeleteCompany)
		authGroup.DELETE("/companies/:uuid/hard", companyHandler.DeleteCompanyHard)
	}
}
