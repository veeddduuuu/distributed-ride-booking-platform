package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
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

	server:= &http.Server{
		Addr:   ":8083",
		Handler: mux,
	}

	serverErrors := make(chan error, 1)

	go func(){
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err:= <-serverErrors:
		log.Printf("Error starting server: %v\n", err)
	case err:= <-shutdown:
		log.Printf("Received shutdown signal: %v\n", err)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down server gracefully: %v\n", err)
			server.Close()
		}
	}
}
