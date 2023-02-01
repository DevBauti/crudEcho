package routes

import (
	"github.com/DevBauti/crudEcho/controllers"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetAUser)
	e.PUT("/user/:userId", controllers.EditAUser)
	e.DELETE("/user/:userId", controllers.DeleteAUser)
	e.GET("/users", controllers.GetAllUsers)
}
