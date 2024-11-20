package main

import (
	"companies_handling/internal"
	"companies_handling/routes"
	"fmt"
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
	var config internal.Config
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
	err5 := db.AutoMigrate(&internal.User{}, &internal.Company{})
	if err5 != nil {
		panic(err5)
	}

	r := gin.Default()
	routes.SetUpRoutes(r)
	fmt.Println(r.Routes())
}
