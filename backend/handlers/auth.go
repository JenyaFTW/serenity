package handlers

import (
	"backend/auth"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

var jwtAuthentication = auth.TokenManager{}

func AuthLogin(c *gin.Context) {
	var postData loginJSON

	if err := c.BindJSON(&postData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	var findUser models.User
	if err := models.DB.First(&findUser, "email = ?", postData.Email); err == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid email/password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(postData.Password)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid email/password"})
		return
	}

	ts, err := jwtAuthentication.CreateToken(findUser.ID.String(), findUser.Username)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	})
}

func AuthSignup(c *gin.Context) {
	var postData signUpJSON

	if err := c.BindJSON(&postData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	var findUser models.User
	if err := models.DB.First(&findUser, "email = ? OR username = ?", postData.Email, postData.Username).Error; err == nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "User with such email/username already exists"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(postData.Password), bcrypt.DefaultCost)
	user := models.User{Username: postData.Username, Email: postData.Email, Password: string(hashedPassword)}
	models.DB.Create(&user)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully registered account"})
}
