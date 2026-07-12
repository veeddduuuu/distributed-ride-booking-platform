package main

import (
	"log"
	"net/http"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRiderWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err:= upgrader.Upgrade(w, r, nil)
	if err!=nil {
		log.Printf("Websocket Upgrade failed : %v", err)
		return
	}
	defer conn.Close()
	userId:= r.URL.Query().Get("userId")
	log.Printf("Rider connected with userId: %s", userId)
	if(userId==""){
		log.Printf("userId is empty")
		return
	}
	for{
		_, message, err := conn.ReadMessage()
		if err!=nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message from rider: %s", message)
	}
}

func handleDriverWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err:= upgrader.Upgrade(w, r, nil)
	if err!=nil {
		log.Printf("Websocket Upgrade failed : %v", err)
		return
	}
	defer conn.Close()
	
	userId:= r.URL.Query().Get("userId")
	log.Printf("Driver connected with userId: %s", userId)
	if(userId==""){
		log.Printf("userId is empty")
		return
	}

	packageSlug:= r.URL.Query().Get("packageSlug")
	log.Printf("Driver connected with packageSlug: %s", packageSlug)
	if(packageSlug==""){
		log.Printf("packageSlug is empty")
		return
	}

	type Driver struct{
		UserId string `json:"userId"`
		Name string `json:"name"`
		CarNumber string `json:"carNumber"`
		ProfilePicture string `json:"profilePicture"`
		PackageSlug string `json:"packageSlug"`
	}

	msg := types.WSMessage{
		Type: "driver_registered",
		Data: Driver{
			UserId: userId,
			Name: "Max Verstappen",
			CarNumber: "ABC123",
			ProfilePicture: "https://example.com/profile.jpg",
			PackageSlug: packageSlug,
		},
	}

	if err:= conn.WriteJSON(msg); err!=nil {
		log.Printf("Error sending message: %v", err)
		return
	}


	for{
		_, message, err := conn.ReadMessage()
		if err!=nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message from driver: %s", message)
	}
}
