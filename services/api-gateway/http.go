package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/api-gateway/grpc_clients"
	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/trip"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody TripPreviewRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validation
	if reqBody.UserId == "" {
		http.Error(w, "UserId is required", http.StatusBadRequest)
		return
	}
	if reqBody.Pickup.Latitude == 0 || reqBody.Pickup.Longitude == 0 {
		http.Error(w, "Pickup coordinates are required", http.StatusBadRequest)
		return
	}
	if reqBody.Destination.Latitude == 0 || reqBody.Destination.Longitude == 0 {
		http.Error(w, "Destination coordinates are required", http.StatusBadRequest)
		return
	}

	fmt.Println("Received trip preview request:", reqBody)

	// create the gRPC client
	tripservice, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Printf("failed to connect to trip service: %v", err)
		http.Error(w, "Trip service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer tripservice.Close()

	// call PreviewTrip over gRPC (not HTTP)
	resp, err := tripservice.Client.PreviewTrip(r.Context(), &pb.PreviewTripRequest{
		UserId: reqBody.UserId,
		Pickup: &pb.Coordinate{
			Latitude:  reqBody.Pickup.Latitude,
			Longitude: reqBody.Pickup.Longitude,
		},
		Destination: &pb.Coordinate{
			Latitude:  reqBody.Destination.Latitude,
			Longitude: reqBody.Destination.Longitude,
		},
	})
	if err != nil {
		log.Printf("PreviewTrip gRPC call failed: %v", err)
		http.Error(w, fmt.Sprintf("Failed to preview trip: %v", err), http.StatusInternalServerError)
		return
	}

	// write the proto response back as JSON to the HTTP client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
