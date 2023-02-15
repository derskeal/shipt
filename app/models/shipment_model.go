package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model

	OriginLocationID      int
	DestinationLocationID int
	OriginLocation        Country
	DestinationLocation   Country

	Weight float32
	Price  float32
}
