package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/api-gateway/grpc_clients"
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
	
	bodyBytes, _ := json.Marshal(reqBody)
	importBytes := bytes.NewBuffer(bodyBytes)
	
	tripservice, err := grpc_clients.NewTripServiceClient()
	if err!=nil{
		log.Fatal(err)
	}

	defer tripservice.Close()


	resp, err := http.Post("http://trip-service:8083/preview", "application/json", importBytes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error forwarding request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	
	io.Copy(w, resp.Body)
}
