package main

import (
	"surge/identity/m/controllers"
	"surge/identity/m/models"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "surge/identity/m/docs"
)

// @title           Swagger Identity API
// @description     Identity swagger docs.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("api/users/signup", controllers.SignupUser)
	r.POST("api/users/login", controllers.ObtainJWT)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
