package repository

import (
	"context"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/domain"
)

type inmemRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip domain.TripModel) (domain.TripModel, error) {
	r.trips[trip.ID] = &trip
	return trip, nil
}

func (r *inmemRepository) SaveRideFares(ctx context.Context, fares []*domain.RideFareModel) error {
	for _, fare := range fares {
		r.rideFares[fare.ID] = fare
	}
	return nil
}

func (r *inmemRepository) GetRideFare(ctx context.Context, fareId string) (*domain.RideFareModel, error) {
	fare, exists := r.rideFares[fareId]
	if !exists {
		return nil, context.DeadlineExceeded // Using a dummy error for now, should use a custom not found error
	}
	return fare, nil
}
