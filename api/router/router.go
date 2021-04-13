package router

import (
	"github.com/GotoRen/go-restful-api/api/controller"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
}
