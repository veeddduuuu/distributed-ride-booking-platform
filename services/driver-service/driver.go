package driverservice

import "context"

type Driver struct{
	UserId string
	Name string
	CarNumber string
	ProfilePicture string
	PackageSlug string
}

type DriverRepository interface {
	AddDriver(ctx context.Context, driver Driver) (*Driver, error)
	RemoveDriver(ctx context.Context, driverId string) (*Driver, error)
	GetDriver(ctx context.Context, driverId string) (*Driver, error)
}

type DriverService interface{
	RegisterDriver(ctx context.Context, driverId string, packageSlug string) (Driver, error)
	UnregisterDriver(ctx context.Context, driverId string) (Driver, error)
}