package main

import (
	"context"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/domain"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/infrastructure/repository"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/service"
	"log"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	fare := &domain.RideFareModel{
		UserId:      "01",
		PackageSlug: "VIP",
		TotalPrice:  100,
	}
	t, err := svc.CreateTrip(ctx, *fare)
	if err != nil {
		log.Println(err)
	}
	log.Println(t)
}
