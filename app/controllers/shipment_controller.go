package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shipt/app/dto"
	"shipt/app/services"
	"shipt/app/utils"
)

func CreateShipment(context echo.Context) error {
	shipmentRequest := new(dto.ShipmentRequest)
	if err := context.Bind(shipmentRequest); err != nil {
		fmt.Println(err)
		return utils.BindingError(err)
	}
	if err := context.Validate(shipmentRequest); err != nil {
		return err
	}

	res, err := services.CreateShipment(*shipmentRequest)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, res)
}

func GetShipments(context echo.Context) error {
	res, err := services.GetShipments()
	if err != nil {
		println(err)
		return context.JSON(http.StatusInternalServerError, "Could not fetch shipments")
	}
	return context.JSON(http.StatusOK, utils.SuccessResponse(res))
}
