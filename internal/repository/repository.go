package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository        IUserRepository
	CityRepository        ICityRepository
	PlaceRepository       IPlaceRepository
	CulinaryRepository    ICulinaryRepository
	TicketRepository      ITicketRepository
	GuideRepository       IGuideRepository
	OrderRepository       IOrderRepository
	TransactionRepository ITransactionRepository
	CommentRepository     ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	cityRepository := NewCityRepository(db)
	placeRepository := NewPlaceRepository(db)
	culinaryRepository := NewCulinaryRepository(db)
	ticketRepository := NewTicketRepository(db)
	guideRepository := NewGuideRepository(db)
	orderRepository := NewOrderRepository(db)
	transactionRepository := NewTransactionRepository(db)
	commentRepository := NewCommentRepository(db)

	return &Repository{
		UserRepository:        userRepository,
		CityRepository:        cityRepository,
		PlaceRepository:       placeRepository,
		CulinaryRepository:    culinaryRepository,
		TicketRepository:      ticketRepository,
		GuideRepository:       guideRepository,
		OrderRepository:       orderRepository,
		TransactionRepository: transactionRepository,
		CommentRepository:     commentRepository,
	}
}
