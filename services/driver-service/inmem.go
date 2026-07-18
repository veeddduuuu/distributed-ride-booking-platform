package driverservice

import (
	"context"
	"fmt"
)

type inmemRepository struct {
	driver map[string]*Driver
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		driver: make(map[string]*Driver),
	}
}

func (r *inmemRepository) AddDriver(ctx context.Context, driver Driver) (*Driver, error) {
	r.driver[driver.UserId] = &driver
	return &driver, nil
}

func (r *inmemRepository) RemoveDriver(ctx context.Context, driverId string) (*Driver, error) {
	driver, exists := r.driver[driverId]
	if !exists {
		return nil, fmt.Errorf("driver %s not found", driverId)
	}
	delete(r.driver, driverId)
	return driver, nil
}

func (r *inmemRepository) GetDriver(ctx context.Context, driverId string) (*Driver, error) {
	driver, exists := r.driver[driverId]
	if !exists {
		return nil, fmt.Errorf("driver %s not found", driverId)
	}
	return driver, nil
}