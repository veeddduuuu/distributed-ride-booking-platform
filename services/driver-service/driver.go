package main

import (
	"context"

	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
)

type Driver struct {
	UserId         string
	Name           string
	CarNumber      string
	ProfilePicture string
	PackageSlug    string
	Location       *types.Coordinates
}

type DriverRepository interface {
	AddDriver(ctx context.Context, driver Driver) (*Driver, error)
	RemoveDriver(ctx context.Context, driverId string) (*Driver, error)
	GetDriver(ctx context.Context, driverId string) (*Driver, error)
}

type DriverService interface{
	RegisterDriver(ctx context.Context, driverId string, packageSlug string, location *types.Coordinates) (Driver, error)
	UnregisterDriver(ctx context.Context, driverId string) (Driver, error)
}