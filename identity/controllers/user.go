package controllers

import (
	"net/http"

	// "surge/identity/m/models"

	"github.com/gin-gonic/gin"
)

type UserSignupInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ObtainJWT godoc
// @Accept   json
// @Produce  json
// @Router   /api/users/login [post]
func ObtainJWT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SignupUser godoc
// @Accept   json
// @Produce  json
// @Router   /api/users/signup [post]
func SignupUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
