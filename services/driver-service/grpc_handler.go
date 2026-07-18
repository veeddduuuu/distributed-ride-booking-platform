package driverservice

import (
	"context"

	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/driver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DriverServiceHandler struct{
	pb.UnimplementedDriverServiceServer
	svc DriverService
}

func NewDriverService(svc DriverService)(*DriverServiceHandler){
	return &DriverServiceHandler{svc : svc}
}

func (h *DriverServiceHandler) RegisterDriver(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	driverId:=req.DriverId
	packageSlug:=req.PackageSlug
	driver, err:= h.svc.RegisterDriver(ctx, driverId, packageSlug)
	if err!=nil {
		return nil, status.Errorf(codes.Internal, "Could not register driver: %v", err)
	}
	return &pb.RegisterResponse{
		Driver: &pb.Driver{
			UserId: driver.UserId,
			Name: driver.Name,
			CarNumber: driver.CarNumber,
			ProfilePicture: driver.ProfilePicture,
			PackageSlug: driver.PackageSlug,
		},
	}, nil
	// return nil, status.Error(codes.Unimplemented, "method RegisterDriver not implemented")

}

func (h *DriverServiceHandler) UnregisterDriver(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	driverId:=req.DriverId
	driver, err:= h.svc.UnregisterDriver(ctx, driverId)
	if err!=nil {
		return nil, status.Errorf(codes.Internal, "Could not unregister driver: %v", err)
	}
	return &pb.UnregisterResponse{
		Driver: &pb.Driver{
			UserId: driver.UserId,
			Name: driver.Name,
			CarNumber: driver.CarNumber,
			ProfilePicture: driver.ProfilePicture,
			PackageSlug: driver.PackageSlug,
		},
	}, nil
	// return nil, status.Error(codes.Unimplemented, "method UnregisterDriver not implemented")
}
