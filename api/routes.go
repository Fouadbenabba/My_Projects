package api

import (
	"first/api/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/cities/:id", controllers.GetCity)
	e.GET("/cities/all", controllers.GetAllCities)
	e.POST("/cities", controllers.CreateCity)
	e.DELETE("/cities/:id", controllers.DeleteCity)
	e.PUT("/cities/:id", controllers.UpdateCity)
}
