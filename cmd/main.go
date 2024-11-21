package main

import (
	"companies_handling/handlers"
	"companies_handling/models"
	"companies_handling/repositories"
	"companies_handling/routes"
	"companies_handling/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Reading .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cf, err2 := os.ReadFile("config.yaml")
	if err2 != nil {
		panic(err2)
	}

	cfs := os.Expand(string(cf), os.Getenv)
	// Decode yaml file
	var config models.Config
	if err3 := yaml.Unmarshal([]byte(cfs), &config); err3 != nil {
		panic(err3)
	}
	// Creating the DSN
	dsn := config.GetDSN()
	db, err4 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err4 != nil {
		panic(err4)
	}
	// Performin Migrations
	err5 := db.AutoMigrate(&models.User{}, &models.Company{})
	if err5 != nil {
		panic(err5)
	}

	// Creating the repo, service and handler user will use
	userRep := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRep)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	routes.SetUpRoutes(r, userHandler)
	fmt.Println(r.Routes())
	if err6 := r.Run(":8080"); err6 != nil {
		log.Fatalf("Failed to start the server: %v\n", err6)
	}
}
