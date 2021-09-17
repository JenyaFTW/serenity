package handlers

import (
	"backend/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type signUpJSON struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginJSON struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthLogin(c *fiber.Ctx) error {
	c.Accepts("json", "text")

	postData := new(loginJSON)

	if err := c.BodyParser(postData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}

	var findUser models.User
	if err := models.DB.First(&findUser, "email = ?", postData.Email); err == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Invalid email/password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(postData.Password)); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Invalid email/password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // week expiry
	claims["id"] = findUser.ID

	signed, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": signed,
	})
}

func AuthSignup(c *fiber.Ctx) error {
	c.Accepts("json", "text")

	postData := new(signUpJSON)

	if err := c.BodyParser(postData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}

	var findUser models.User
	if err := models.DB.First(&findUser, "email = ? OR username = ?", postData.Email, postData.Username).Error; err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"message": "User with such email/username already exists"})
	}

	md5Hasher := md5.New()
	md5Hasher.Write([]byte(postData.Email))
	hashedEmail := hex.EncodeToString(md5Hasher.Sum(nil))

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(postData.Password), bcrypt.DefaultCost)
	user := models.User{Username: postData.Username, Email: postData.Email, Password: string(hashedPassword), Avatar: fmt.Sprint("https://gravatar.com/avatar/", hashedEmail)}
	models.DB.Create(&user)

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Successfully registered account"})
}

func AuthMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var findUser models.User
	if err := models.DB.First(&findUser, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	return c.Status(200).JSON(fiber.Map{
		"id":       findUser.ID,
		"email":    findUser.Email,
		"username": findUser.Username,
		"avatar":   findUser.Avatar,
		"role":     findUser.Role,
	})
}
