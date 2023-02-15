package repositories

import (
	"fmt"
	"shipt/app/config"
	"shipt/app/dto"
	"shipt/app/models"
)

func CreateShipment(shipment dto.ShipmentRequest, originCountry models.Country, destinationCountry models.Country, price float32) (models.Shipment, error) {
	newS := models.Shipment{
		Weight: shipment.Weight,
		Price:  price,
	}
	err := config.DB.Model(&newS).Association("OriginLocation").Append(&originCountry)
	if err != nil {
		return models.Shipment{}, err
	}
	err = config.DB.Model(&newS).Association("DestinationLocation").Append(&destinationCountry)
	if err != nil {
		return models.Shipment{}, err
	}

	result := config.DB.Create(&newS)

	fmt.Printf("ID: %v. Rows: %d", newS.ID, result.RowsAffected)
	return newS, result.Error
}

func GetShipments() ([]models.Shipment, error) {
	var shipments []models.Shipment

	result := config.DB.Preload("OriginLocation").Preload("DestinationLocation").Find(&shipments)

	return shipments, result.Error
}
