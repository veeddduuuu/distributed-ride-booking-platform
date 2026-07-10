package service

import (
	"context"
	"fmt"

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
