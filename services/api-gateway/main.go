package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting API Gateway on port 8080...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the API Gateway!"))
	})

	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
