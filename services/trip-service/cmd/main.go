package main

import (
	"context"
	"go-ride-platform/services/trip-service/internal/domain"
	"go-ride-platform/services/trip-service/internal/infrastructure/repository"
	"go-ride-platform/services/trip-service/internal/service"
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
