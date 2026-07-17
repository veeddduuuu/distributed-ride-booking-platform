package main

import (
	"context"
)

type inmemRepository struct {
	trips     map[string]*TripModel
	rideFares map[string]*RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trips:     make(map[string]*TripModel),
		rideFares: make(map[string]*RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip TripModel) (TripModel, error) {
	r.trips[trip.ID] = &trip
	return trip, nil
}

func (r *inmemRepository) GetTrip(ctx context.Context, tripId string) (*TripModel, error){
	trip, exists := r.trips[tripId]
	if !exists{
		return nil, context.DeadlineExceeded
	}
	return trip, nil
}

func (r *inmemRepository) SaveRideFares(ctx context.Context, fares []*RideFareModel) error {
	for _, fare := range fares {
		r.rideFares[fare.ID] = fare
	}
	return nil
}

func (r *inmemRepository) GetRideFare(ctx context.Context, fareId string) (*RideFareModel, error) {
	fare, exists := r.rideFares[fareId]
	if !exists {
		return nil, context.DeadlineExceeded // Using a dummy error for now, should use a custom not found error
	}
	return fare, nil
}
