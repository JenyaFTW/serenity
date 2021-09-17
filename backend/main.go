package main

import (
	"backend/handlers"
	"backend/middlewares"
	"backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, make sure you cloned .env.example")
	}

	models.ConnectDatabase()

	app := fiber.New()

	app.Static("/", "../frontend/dist")

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", handlers.AuthLogin)
	auth.Post("/signup", handlers.AuthSignup)
	auth.Get("/me", middlewares.AuthRequired(), handlers.AuthMe)

	projects := api.Group("/projects", middlewares.AuthRequired())
	projects.Get("/", handlers.GetProjects)
	projects.Post("/", handlers.PostProjects)
	projects.Get("/:id", handlers.GetProjectById)
	projects.Delete("/:id", handlers.DeleteProjectById)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("../frontend/dist/index.html")
	})

	app.Listen(":8080")
}
