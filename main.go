package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"shipt/app/config"
)

func main() {
	e := echo.New()

	e.Validator = &config.CustomValidator{Validator: validator.New()}
	ConfigureRoutes(e)
	config.InitDatabase()
	config.MigrateDatabase()

	url := fmt.Sprintf(":%s", config.Envs.Port)
	e.Logger.Fatal(e.Start(url))
}
