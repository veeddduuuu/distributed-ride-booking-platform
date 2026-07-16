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

type RideShare struct {
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PackageSlug   string                 `protobuf:"bytes,3,opt,name=packageSlug,proto3" json:"packageSlug,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
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

type PreviewTripResponse struct {
	TripId        string                 `protobuf:"bytes,1,opt,name=tripId,proto3" json:"tripId,omitempty"`
	Route         *Route                 `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	RideFares     []*RideShare           `protobuf:"bytes,3,rep,name=rideFares,proto3" json:"rideFares,omitempty"`
}

