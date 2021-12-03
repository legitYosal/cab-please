package controllers

import (
	"net/http"

	"surge/identity/m/models"

	"github.com/gin-gonic/gin"
)

type UserSignupInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ObtainJWT godoc
// @Summary   Obtain JWT
// @Accept   json
// @Param user body UserLogin true "User Data"
// @Success 200 {object} object
// @Produce  json
// @tags 	  User
// @Router   /api/users/login [post]
func ObtainJWT(c *gin.Context) {
	var input UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// SignupUser godoc
// @Summary   Signup user
// @tags 	  User
// @Param user body UserSignupInput true "User Data"
// @Success 200 {object} object
// @Accept    json
// @Produce   json
// @Router    /api/users/signup [post]
func SignupUser(c *gin.Context) {
	var input UserSignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: input.Username,
		Password: models.HashPassword(input.Password),
	}
	result := models.DB.Create(&user)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}
}

// GetProfile godoc
// @Summary   get user profile
// @Accept   json
// @Produce  json
// @tags 	  User
// @Security ApiKeyAuth
// @Router   /api/users/profile [get]
func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
