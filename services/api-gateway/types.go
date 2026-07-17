package main

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TripPreviewRequest struct {
	UserId      string      `json:"userId"`
	Pickup      Coordinates `json:"pickup"`
	Destination Coordinates `json:"destination"`
}

type CreateTripRequest struct {
	FareId string `json:"fareId"`
}

