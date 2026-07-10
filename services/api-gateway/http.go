package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody TripPreviewRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err!=nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//validation
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
	
	resp, err := http.Post("http://localhost:8083/preview", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Body)
}
