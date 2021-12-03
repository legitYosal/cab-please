package main

import (
	"github.com/usefss/cab-please/identity/middleware"
	utils "github.com/usefss/cab-please/identity/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(host + ":" + port)
}
