package driverservice

import "context"

type Driver struct{
	UserId string
	Name string
	CarNumber string
	ProfilePicture string
	PackageSlug string
}

type DriverService interface{
	RegisterDriver(ctx context.Context, driverId string, pacakgeSlug string) (Driver,error)
	UnregisterDriver(ctx context.Context, driverId string) (Driver, error)
}