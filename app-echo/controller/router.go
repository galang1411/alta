package controller

import (
	"app-echo/service"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, us service.DBUserService) {
	userController := NewUserController(us)
	e.GET("/users", userController.GetAll)
	e.POST("/users", userController.Create)
}
