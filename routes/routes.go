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

	userGroup := router.Group("/users")
	userGroup.Use(middlewares.JwtCheck)
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
		userGroup.DELETE("/:id/hard", userHandler.DeleteUserHard)

		userGroup.GET("/:id/companies/:uuid", companyHandler.GetCompany)
		userGroup.POST("/:id/companies", companyHandler.CreateCompany)
		userGroup.PATCH("/:id/companies/:uuid", companyHandler.EditCompany)
		userGroup.DELETE("/:id/companies/:uuid", companyHandler.DeleteCompany)
		userGroup.DELETE("/:id/companies/:uuid/hard", companyHandler.DeleteCompanyHard)
	}
}
