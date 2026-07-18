package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/api-gateway/grpc_clients"
	pbdriver "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/driver"
	pbtrip "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/trip"
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
	resp, err := tripservice.Client.PreviewTrip(r.Context(), &pbtrip.PreviewTripRequest{
		UserId: reqBody.UserId,
		Pickup: &pbtrip.Coordinate{
			Latitude:  reqBody.Pickup.Latitude,
			Longitude: reqBody.Pickup.Longitude,
		},
		Destination: &pbtrip.Coordinate{
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

func handleCreateTrip(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Received create trip request: ", reqBody)

	tripservice, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Printf("failed to connect to trip service: %v", err)
		http.Error(w, "Trip service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer tripservice.Close()

	resp, err := tripservice.Client.CreateTrip(r.Context(), &pbtrip.CreateTripRequest{
		FareId: reqBody.FareId,
	})
	if err != nil {
		log.Printf("CreateTrip grpc call failed: %v", err)
		http.Error(w, "CreateTrip grpc call failed", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func handleRegisterDriver(w http.ResponseWriter, r *http.Request) {
	var reqBody RegisterDriverRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.DriverId == "" {
		http.Error(w, "DriverId is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Received register driver request:", reqBody)

	driverservice, err := grpc_clients.NewDriverServiceClient()
	if err != nil {
		log.Printf("failed to connect to driver service: %v", err)
		http.Error(w, "Driver service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer driverservice.Close()

	resp, err := driverservice.Client.RegisterDriver(r.Context(), &pbdriver.RegisterRequest{
		DriverId:    reqBody.DriverId,
		PackageSlug: reqBody.PackageSlug,
		Location: &pbdriver.Coordinate{
			Latitude:  reqBody.Location.Latitude,
			Longitude: reqBody.Location.Longitude,
		},
	})
	if err != nil {
		log.Printf("RegisterDriver grpc call failed: %v", err)
		http.Error(w, fmt.Sprintf("Failed to register driver: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func handleUnregisterDriver(w http.ResponseWriter, r *http.Request) {
	var reqBody UnregisterDriverRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.DriverId == "" {
		http.Error(w, "DriverId is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Received unregister driver request:", reqBody)

	driverservice, err := grpc_clients.NewDriverServiceClient()
	if err != nil {
		log.Printf("failed to connect to driver service: %v", err)
		http.Error(w, "Driver service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer driverservice.Close()

	resp, err := driverservice.Client.UnregisterDriver(r.Context(), &pbdriver.UnregisterRequest{
		DriverId: reqBody.DriverId,
	})
	if err != nil {
		log.Printf("UnregisterDriver grpc call failed: %v", err)
		http.Error(w, fmt.Sprintf("Failed to unregister driver: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}