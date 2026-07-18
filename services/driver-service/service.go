package driverservice

import (
	"context"
	"fmt"

	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
)

type Service struct {
	repo DriverRepository
}

func NewService(repo DriverRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterDriver(ctx context.Context, driverId string, packageSlug string, location *types.Coordinates) (Driver, error) {
	existing, _ := s.repo.GetDriver(ctx, driverId)
	if existing != nil {
		return Driver{}, fmt.Errorf("driver %s is already registered", driverId)
	}

	driver := Driver{
		UserId:      driverId,
		PackageSlug: packageSlug,
		Location:    location,
		Name: "Max Verstappen",
		CarNumber: "TuTuDuDu",
		ProfilePicture: "https://www.printables.com/model/836977-max-verstappen-sid-meme-keyring/user-gcodes",
	}

	saved, err := s.repo.AddDriver(ctx, driver)
	if err != nil {
		return Driver{}, fmt.Errorf("failed to register driver: %w", err)
	}

	return *saved, nil
}

func (s *Service) UnregisterDriver(ctx context.Context, driverId string) (Driver, error) {
	removed, err := s.repo.RemoveDriver(ctx, driverId)
	if err != nil {
		return Driver{}, fmt.Errorf("failed to unregister driver: %w", err)
	}

	return *removed, nil
}
