package routes

import (
	"be_test_linkque/docs/swagger/v1/extl/swaggo"
	"be_test_linkque/infrastructure/v1/mysql/repository"
	"be_test_linkque/utils/config"
	"be_test_linkque/utils/helper"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	userService "be_test_linkque/core/service/user"

	userCtrl "be_test_linkque/interface/extl/v1/user"
	// echoSwagger "github.com/swaggo/echo-swagger"
)

// API godoc
// @title  be_test_linkque API
// @version 1.0
// @description This is a  be_test_linkque API Documentation.
// @termsOfService
// @contact.name
// @contact.url
// @contact.message
// @securityDefinitions.basic  jwt

func API(e *echo.Echo) {
	// var m mid.AuthJwt
	// swagger
	config.LoadEnvVars()
	swaggo.SwaggerInfo.Host = os.Getenv("APP_BASE_URL")
	// Instance DB
	db := config.MySQL
	helper := helper.GetInstance()
	userRepo := repository.NewUserRepository(db)
	clientServ := userService.New(userRepo, helper)

	e.Use(middleware.RemoveTrailingSlash())

	baseUrl := "/api/v1"
	userHandler := userCtrl.New(clientServ)
	userRoute := e.Group(fmt.Sprintf("%s/users", baseUrl))
	userRoute.POST("", userHandler.Create)

	// swagger
	swaggerRoute := e.Group(fmt.Sprintf("%s/swagger/*", baseUrl))
	swaggerRoute.GET("", echoSwagger.WrapHandler)
}
