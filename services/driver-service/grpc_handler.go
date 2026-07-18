package driverservice

import (
	"context"

	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/driver"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DriverServiceHandler struct{
	pb.UnimplementedDriverServiceServer
	svc DriverService
}

func NewDriverServiceHandler(svc DriverService) *DriverServiceHandler {
	return &DriverServiceHandler{svc: svc}
}

func (h *DriverServiceHandler) RegisterDriver(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	driverId := req.DriverId
	packageSlug := req.PackageSlug

	var location *types.Coordinates
	if req.Location != nil {
		location = &types.Coordinates{
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		}
	}

	driver, err := h.svc.RegisterDriver(ctx, driverId, packageSlug, location)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not register driver: %v", err)
	}

	resp := &pb.RegisterResponse{
		Driver: driverToProto(&driver),
	}
	return resp, nil
}

func (h *DriverServiceHandler) UnregisterDriver(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	driverId := req.DriverId
	driver, err := h.svc.UnregisterDriver(ctx, driverId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not unregister driver: %v", err)
	}

	resp := &pb.UnregisterResponse{
		Driver: driverToProto(&driver),
	}
	return resp, nil
}

// driverToProto converts a domain Driver to a protobuf Driver.
// Keeps the Pack logic in one place — DRY.
func driverToProto(d *Driver) *pb.Driver {
	pbDriver := &pb.Driver{
		UserId:         d.UserId,
		Name:           d.Name,
		CarNumber:      d.CarNumber,
		ProfilePicture: d.ProfilePicture,
		PackageSlug:    d.PackageSlug,
	}
	if d.Location != nil {
		pbDriver.Location = &pb.Coordinate{
			Latitude:  d.Location.Latitude,
			Longitude: d.Location.Longitude,
		}
	}
	return pbDriver
}
