package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shipt/app/controllers"
)

func ConfigureRoutes(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.POST("/shipments", controllers.CreateShipment)

	app.GET("/shipments", controllers.GetShipments)
}
