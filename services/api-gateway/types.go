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

type RegisterDriverRequest struct {
	DriverId    string      `json:"driverId"`
	PackageSlug string      `json:"packageSlug"`
	Location    Coordinates `json:"location"`
}

type UnregisterDriverRequest struct {
	DriverId string `json:"driverId"`
}