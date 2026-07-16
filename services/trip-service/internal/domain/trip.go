package domain

import (
	"context"

	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto"
)

type RideFareModel struct {
	ID          primitive.ObjectID
	TripID      string
	UserId      string
	PackageSlug string
	TotalPrice  float64
}

type TripModel struct {
	ID       primitive.ObjectID
	UserId   string
	Status   string
	RideFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip TripModel) (TripModel, error)
	SaveRideFares(ctx context.Context, fares []*RideFareModel) error
	GetRideFare(ctx context.Context, fareId string) (*RideFareModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination *types.Coordinates) (*types.Route, error)
	GetFares(userId string, distance float64, duration float64) ([]*types.RideShare, error)
	PreviewTrip(ctx context.Context, userId string, pickup *types.Coordinates, destination *types.Coordinates) (*types.PreviewTripResponse, error)
	TripStart(ctx context.Context, ridedetails *types.RideShare) (*types.TripStartResponse, error)
}
