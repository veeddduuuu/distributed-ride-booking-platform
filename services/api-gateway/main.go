package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting API Gateway on port 8080...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the API Gateway!"))
	})

	mux.HandleFunc("POST /trip/preview", handleTripPreview)
	mux.HandleFunc("POST /ws/rider", handleRiderWebSocket)
	mux.HandleFunc("POST /ws/driver", handleDriverWebSocket)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		fmt.Printf("Error starting server: %v\n", err)
	case sig := <-shutdown:
		fmt.Printf("Received shutdown signal: %v\n", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Error shutting down server gracefully: %v\n", err)
			server.Close()
		}
	}
}
