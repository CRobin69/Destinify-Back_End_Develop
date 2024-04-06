package service

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/repository"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/helper"
)

type Service struct {
	UserService        IUserService
	TicketService      ITicketService
	CityService        ICityService
	PlaceService       IPlaceService
	CulinaryService    ICulinaryService
	GuideService       IGuideService
	OrderService       IOrderService // Add missing import for IOrderService
	TransactionService ITransactionService
	CommentService     ICommentService
}

func NewService(Repository *repository.Repository) *Service {
	userService := NewUserService(Repository.UserRepository)
	cityService := NewCityService(Repository.CityRepository)
	placeService := NewPlaceService(Repository.PlaceRepository)
	culinaryService := NewCulinaryService(Repository.CulinaryRepository)
	guideService := NewGuideService(Repository.GuideRepository)
	ticketService := NewTicketService(Repository.TicketRepository, Repository.OrderRepository, Repository.PlaceRepository, Repository.GuideRepository)
	OrderService := NewOrderService(Repository.OrderRepository)
	TransactionService := NewTransactionService(Repository.TransactionRepository, &helper.MdtClient{}, Repository.UserRepository, Repository.OrderRepository, Repository.TicketRepository, Repository.GuideRepository, Repository.PlaceRepository, Repository.CommentRepository)
	CommentService := NewCommentService(Repository.CommentRepository)
	
	return &Service{
		UserService:        userService,
		TicketService:      ticketService,
		CityService:        cityService,
		PlaceService:       placeService,
		CulinaryService:    culinaryService,
		GuideService:       guideService,
		OrderService:       OrderService,
		TransactionService: TransactionService,
		CommentService:     CommentService,
	}
}
