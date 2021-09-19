package handlers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetSubdomains(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	projectId := c.Locals("id")

	var project models.Project
	models.DB.Find(&project, "id = ? AND user_id = ?", projectId, userId)
	// Add error handling

	var subdomains []models.Subdomain
	models.DB.Find(&subdomains, "project_id = ?", project.ID)

	return c.JSON(fiber.Map{
		"subdomains": subdomains,
	})
}
