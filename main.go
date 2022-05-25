package main

import (
	"github.com/labstack/echo/v4"
	"jwt-auth.com/config"
	"jwt-auth.com/controller"
)

func main() {
	e := echo.New()

	config := config.Initconfig()

	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.GET("/details", controller.Details)

	e.Start(":" + config.Port)

}
