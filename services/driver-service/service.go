package driverservice

type Service struct{
	repo DriverRepository
}

func NewService(repo DriverRepository) (*Service){
	return &Service{
		repo : repo,
	}
}

