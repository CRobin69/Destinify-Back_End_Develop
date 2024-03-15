package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	CityRepository ICityRepository
	PlaceRepository IPlaceRepository
	CulinaryRepository ICulinaryRepository
	TicketRepository ITicketRepository
	GuideRepository IGuideRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	cityRepository := NewCityRepository(db)
	placeRepository := NewPlaceRepository(db)
	culinaryRepository := NewCulinaryRepository(db)
	ticketRepository := NewTicketRepository(db)
	guideRepository := NewGuideRepository(db)

	return &Repository{
		UserRepository: userRepository,
		CityRepository: cityRepository,
		PlaceRepository: placeRepository,
		CulinaryRepository: culinaryRepository,
		TicketRepository: ticketRepository,
		GuideRepository: guideRepository,
		
	}
}