package main

import (
	"PR-Tracker/api/controller"
	"PR-Tracker/api/database"
	"PR-Tracker/api/middleware"
	"PR-Tracker/api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "fmt"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Workout{})
	database.Database.AutoMigrate(&model.Record{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/workout", controller.CreateWorkout)
	protectedRoutes.GET("/workout", controller.GetAllWorkouts)
	protectedRoutes.DELETE("/workout/:id", controller.DeleteWorkout)

	protectedRoutes.POST("/workout/:id/records", controller.AddRecord)
	protectedRoutes.GET("/workout/:id/records", controller.GetAllRecords)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}
