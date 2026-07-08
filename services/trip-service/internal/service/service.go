package service

import (
	"context"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/domain"

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
