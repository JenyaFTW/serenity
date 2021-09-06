package main

import (
	"backend/handlers"
	"backend/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, make sure you cloned .env.example")
	}

	models.ConnectDatabase()

	r := gin.Default()

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/auth/login", handlers.AuthLogin)
		apiRouter.POST("/auth/signup", handlers.AuthSignup)
	}

	r.Run()
}
