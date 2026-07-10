package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/infrastructure/repository"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/service"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
)

type PreviewRequest struct {
	Pickup      types.Coordinates `json:"pickup"`
	Destination types.Coordinates `json:"destination"`
}

func main() {
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Trip Service!"))
	})

	mux.HandleFunc("POST /preview", func(w http.ResponseWriter, r *http.Request) {
		var req PreviewRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		tripResponse, err := svc.GetRoute(r.Context(), &req.Pickup, &req.Destination)
		if err != nil {
			http.Error(w, "Error creating trip", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tripResponse)
	})
	fmt.Println("Starting Trip Service at port 8083...")

	if err := http.ListenAndServe(":8083", mux); err != nil {
		log.Println("Error starting Trip Service:", err)
	}
}
