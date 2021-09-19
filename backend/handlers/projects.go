package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type postProjectsJSON struct {
	Name       string `json:"name"`
	MainDomain string `json:"main_domain"`
}

func PostProjects(c *fiber.Ctx) error {
	c.Accepts("json", "text")

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	postData := new(postProjectsJSON)
	if err := c.BodyParser(postData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}

	project := models.Project{
		Name:       postData.Name,
		MainDomain: postData.MainDomain,
		UserId:     uuid.MustParse(id),
	}

	models.DB.Create(&project)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Added new project"})
}

func GetProjects(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var projects []models.Project

	models.DB.Find(&projects, "user_id = ?", id)

	return c.JSON(fiber.Map{
		"projects": projects,
	})
}

func GetProjectById(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var project models.Project

	models.DB.Find(&project, "id = ? AND user_id = ?", c.Params("id"), id)

	return c.JSON(project)
}

func DeleteProjectById(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var project models.Project
	models.DB.Delete(&project, "id = ? AND user_id = ?", c.Params("id"), id)

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Successfully deleted project"})
}
