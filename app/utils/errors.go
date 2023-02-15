package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ValidationError(body interface{}) error {
	res := Response{
		Status:  ERROR,
		Message: "Validation failed",
		Data:    body,
	}
	return echo.NewHTTPError(http.StatusUnprocessableEntity, res)
}

func BindingError(err error) error {
	var message = err.(*echo.HTTPError).Message
	res := Response{
		Status:  ERROR,
		Message: message,
	}
	return echo.NewHTTPError(http.StatusBadRequest, res)
}

func OperationError(message string) error {
	res := Response{
		Status:  ERROR,
		Message: message,
	}
	return echo.NewHTTPError(http.StatusUnprocessableEntity, res)
}
