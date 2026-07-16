package service

import (
	"context"
	"fmt"
	"log"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/domain"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"

	"encoding/json"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTrip(ctx context.Context, fare domain.RideFareModel) (*domain.TripModel, error) {
	t := domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserId:   fare.UserId,
		RideFare: &fare,
		Status:   "PENDING",
	}

	createdTrip, err := s.repo.CreateTrip(ctx, t)
	if err != nil {
		return nil, err
	}
	return &createdTrip, nil
}

func (s *Service) GetRoute(ctx context.Context, pickup, destination *types.Coordinates) (*types.Route, error) {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full", pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get route from OSRM: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the response %v", err)
	}
	var osrmResp types.OSRMResponse
	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return nil, fmt.Errorf("Failed to parse the response %v", err)
	}

	if len(osrmResp.Routes) == 0 {
		return nil, fmt.Errorf("No route found")
	}

	route := &types.Route{
		Distance: osrmResp.Routes[0].Distance,
		Duration: osrmResp.Routes[0].Duration,
		Polyline: osrmResp.Routes[0].Polyline,
	}

	return route, nil
}

func (s *Service) GetFares(userId string, distance float64, duration float64) ([]*types.RideShare, error) {
	packages := []struct {
		slug       string
		rateperkm  float64
		ratepermin float64
		bookingfee float64
	}{
		{"sedan", 15.0, 2.0, 30.0},
		{"suv", 20.0, 3.0, 50.0},
		{"van", 25.0, 4.0, 70.0},
		{"luxury", 40.0, 6.0, 120.0},
	}
	fares := make([]*types.RideShare, len(packages))
	for i, p := range packages {
		fareId := primitive.NewObjectID().Hex()
		fares[i] = &types.RideShare{
			Id:          fareId,
			UserId:      userId,
			PackageSlug: p.slug,
			TotalPrice:  distance/1000*p.rateperkm + duration/60*p.ratepermin + p.bookingfee,
		}
	}
	return fares, nil
}

func (s *Service) PreviewTrip(ctx context.Context, userId string, pickup *types.Coordinates, destination *types.Coordinates) (*types.PreviewTripResponse, error) {
	tripId := primitive.NewObjectID().Hex()

	route, err := s.GetRoute(ctx, pickup, destination)
	if err != nil {
		log.Printf("PreviewTrip: failed to get route: %v", err)
		return nil, fmt.Errorf("failed to get route: %w", err)
	}

	ridefares, err := s.GetFares(userId, route.Distance, route.Duration)
	if err != nil {
		log.Printf("PreviewTrip: failed to get fares: %v", err)
		return nil, fmt.Errorf("failed to get fares: %w", err)
	}

	// Save fares to the in-memory repository
	domainFares := make([]*domain.RideFareModel, len(ridefares))
	for i, f := range ridefares {
		objID, _ := primitive.ObjectIDFromHex(f.Id)
		domainFares[i] = &domain.RideFareModel{
			ID:          objID,
			TripID:      tripId,
			UserId:      f.UserId,
			PackageSlug: f.PackageSlug,
			TotalPrice:  f.TotalPrice,
		}
	}

	if err := s.repo.SaveRideFares(ctx, domainFares); err != nil {
		log.Printf("PreviewTrip: failed to save fares: %v", err)
		return nil, fmt.Errorf("failed to save fares: %w", err)
	}

	return &types.PreviewTripResponse{
		TripId:    tripId,
		Route:     route,
		RideFares: ridefares,
	}, nil
}

func (s *Service) TripStart(ctx context.Context, rideDetails *types.RideShare) (*types.TripStartResponse, error) {
	return nil, nil
}
