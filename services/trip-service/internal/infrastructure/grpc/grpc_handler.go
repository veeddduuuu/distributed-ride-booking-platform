package grpc

import (
	"context"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/domain"
	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/trip"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TripServiceHandler struct {
	pb.UnimplementedTripServiceServer
	svc domain.TripService
}

func NewTripServiceHandler(svc domain.TripService) *TripServiceHandler {
	return &TripServiceHandler{svc: svc}
}

func (h *TripServiceHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	userId := req.UserId

	pick := &types.Coordinates{
		Latitude:  req.Pickup.Latitude,
		Longitude: req.Pickup.Longitude,
	}
	dest := &types.Coordinates{
		Latitude:  req.Destination.Latitude,
		Longitude: req.Destination.Longitude,
	}

	preview, err := h.svc.PreviewTrip(ctx, userId, pick, dest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "preview trip failed: %v", err)
	}

	pbFares := make([]*pb.RideShare, len(preview.RideFares))
	for i, f := range preview.RideFares {
		pbFares[i] = &pb.RideShare{
			UserId:      f.UserId,
			PackageSlug: f.PackageSlug,
			TotalPrice:  f.TotalPrice,
		}
	}

	return &pb.PreviewTripResponse{
		TripId: preview.TripId,
		Route: &pb.Route{
			Distance: preview.Route.Distance,
			Duration: preview.Route.Duration,
			Polyline: preview.Route.Polyline,
		},
		RideFares: pbFares,
	}, nil
}
