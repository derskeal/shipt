package dto

type ShipmentResponse struct {
	OriginLocationCode      string  `json:"originLocationCode"`
	DestinationLocationCode string  `json:"destinationLocationCode"`
	Weight                  float32 `json:"weight"`
	Price                   float32 `json:"price"`
}
