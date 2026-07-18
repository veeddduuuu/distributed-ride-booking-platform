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
	return nil, status.Error(codes.Unimplemented, "method RegisterDriver not implemented")

}

func (h *DriverServiceHandler) UnregisterDriver(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UnregisterDriver not implemented")
}
