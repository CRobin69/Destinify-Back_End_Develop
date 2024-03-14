package service

import "INTERN_BCC/internal/repository"

type Service struct {
	UserService IUserService
	TicketService ITicketService
	CityService ICityService
	PlaceService IPlaceService
	CulinaryService ICulinaryService
}

func NewService(Repository *repository.Repository) *Service {
	userService := NewUserService(Repository.UserRepository)
	ticketService := NewTicketService(Repository.TicketRepository)
	cityService := NewCityService(Repository.CityRepository)
	placeService := NewPlaceService(Repository.PlaceRepository)
	culinaryService := NewCulinaryService(Repository.CulinaryRepository)

	return &Service{
		UserService: userService,
		TicketService: ticketService,
		CityService: cityService,
		PlaceService: placeService,
		CulinaryService: culinaryService,
	}
}