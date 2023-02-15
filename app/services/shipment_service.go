package services

import (
	"shipt/app/constants"
	"shipt/app/dto"
	"shipt/app/repositories"
	"shipt/app/utils"
)

func CreateShipment(shipment dto.ShipmentRequest) (dto.ShipmentResponse, error) {
	var (
		price       float32
		err         error
		shipmentRes dto.ShipmentResponse
	)

	// calculate weight-based price
	if shipment.Weight > 0 && shipment.Weight < 10 {
		price = constants.SMALL_WEIGHT_CLASS_PRICE
	} else if shipment.Weight >= 10 && shipment.Weight < 25 {
		price = constants.MEDIUM_WEIGHT_CLASS_PRICE
	} else if shipment.Weight >= 25 && shipment.Weight < 50 {
		price = constants.LARGE_WEIGHT_CLASS_PRICE
	} else if shipment.Weight >= 50 && shipment.Weight <= 1000 {
		price = constants.HUGE_WEIGHT_CLASS_PRICE
	} else {
		println("Unsupported weight")
		err = utils.OperationError(constants.UnsupportedWeight)
	}

	if err != nil {
		return shipmentRes, err
	}

	// calculate location-based price
	originCountry, err := repositories.FindCountryByCode(shipment.OriginLocationCode)
	if err != nil {
		return shipmentRes, err
	}
	destinationCountry, err := repositories.FindCountryByCode(shipment.DestinationLocationCode)
	if err != nil {
		return shipmentRes, err
	}

	if originCountry.ID == destinationCountry.ID {
		price = price * constants.SameZonePriceMultiplier
	} else if originCountry.IsEUMember && destinationCountry.IsEUMember {
		price = price * constants.EuZonePriceMultiplier
	} else {
		price = price * constants.WorldZonePriceMultiplier
	}

	_, err = repositories.CreateShipment(shipment, originCountry, destinationCountry, price)
	if err != nil {
		return shipmentRes, err
	}

	shipmentRes = dto.ShipmentResponse{
		OriginLocationCode:      shipment.OriginLocationCode,
		DestinationLocationCode: shipment.DestinationLocationCode,
		Weight:                  shipment.Weight,
		Price:                   price,
	}

	return shipmentRes, nil
}

func GetShipments() ([]dto.ShipmentResponse, error) {
	var shipmentDtos []dto.ShipmentResponse
	shipments, err := repositories.GetShipments()
	if err != nil {
		return nil, err
	}

	for _, shipment := range shipments {
		shipmentDtos = append(shipmentDtos, dto.ShipmentResponse{
			OriginLocationCode:      shipment.OriginLocation.Code,
			DestinationLocationCode: shipment.DestinationLocation.Code,
			Weight:                  shipment.Weight,
			Price:                   shipment.Price,
		})
	}

	return shipmentDtos, err
}
