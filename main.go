package main

import (
	"bwastartup-be/handler"
	"bwastartup-be/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// hardcoded way
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test simpan dari service"
	// userInput.Email = "test@service.com"
	// userInput.Occupation = "Testing"
	// userInput.Password = "secret123"

	// userService.RegisterUser(userInput)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}
