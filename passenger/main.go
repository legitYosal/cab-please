package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/usefss/cab-please/identity/middleware"
	utils "github.com/usefss/cab-please/identity/utils"
	"github.com/usefss/cab-please/passenger/controllers"

	_ "github.com/usefss/cab-please/passenger/docs"
)

// @title           Swagger Passenger API
// @description     Passenger swagger docs.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	port := utils.GetEnv("PORT")
	host := utils.GetEnv("HOST")

	r := gin.Default()
	r.Use(middleware.AuthorizationMiddleware)

	r.POST("/api/journey/", controllers.RequestJourney)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(host + ":" + port)
}
