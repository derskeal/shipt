package dto

type ShipmentRequest struct {
	OriginLocationCode      string  `json:"originLocationCode" validate:"required,iso3166_1_alpha2"`
	DestinationLocationCode string  `json:"destinationLocationCode" validate:"required,iso3166_1_alpha2"`
	Weight                  float32 `json:"weight" validate:"required,numeric"`
}
