package main

import (
	"surge/identity/m/controllers"
	"surge/identity/m/middleware"
	"surge/identity/m/models"

	"github.com/gin-gonic/gin"

	utils "surge/identity/m/utils"

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
	port := utils.GetEnv("PORT")
	host := utils.GetEnv("HOST")

	r := gin.Default()
	r.Use(middleware.AuthorizationMiddleware)

	models.ConnectDatabase()

	r.POST("api/users/signup", controllers.SignupUser)
	r.POST("api/users/login", controllers.ObtainJWT)
	r.GET("api/users/profile", controllers.GetProfile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(host + ":" + port)
}
