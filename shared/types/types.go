package types

// Coordinates represents a geographical location with latitude and longitude.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Route represents a calculated path between two points, typically including distance, duration, and the path polyline.
type Route struct {
	Distance float64 `json:"distance"` // Distance in meters
	Duration float64 `json:"duration"` // Duration in seconds
	Polyline string  `json:"polyline"` // Encoded polyline string
}

type OSRMResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Polyline string  `json:"geometry"`
	} `json:"routes"`
}

type WSMessage struct {
	Type string `json:"type"`
	Data any `json:"payload"`
}