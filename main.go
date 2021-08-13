package main

import (
	// "bwastartup/user"
	// "fmt"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	// "net/http"

	// "os/user"

	// "github.com/gin-gonic/gin"
	// "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:password147@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8081")


	// handler
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Hendra"
	// userInput.Email = "HendraCahyadi@outlook.com"
	// userInput.Occupation = "General Manager"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	// Input from user
	// handler, mapping from user -> struct input
	// service: doing mapping from input struct to User struct
	// repository
	// db 
}