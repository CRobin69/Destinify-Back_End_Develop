package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"log"

	"INTERN_BCC/model"

	"github.com/google/uuid"
)

type ITicketService interface {
	BuyTickets(param model.TicketBuy) (entity.Order, []entity.Ticket, error)
	GetTicketByID(param model.TicketParam) (entity.Ticket, error)
	GetTicketByUserID(param model.TicketParam) ([]entity.Ticket, error)
}

type TicketService struct {
	tr repository.ITicketRepository
	or repository.IOrderRepository
	pr repository.IPlaceRepository
	gr repository.IGuideRepository
}

func NewTicketService(ticketRepository repository.ITicketRepository, orderRepository repository.IOrderRepository, placeRepository repository.IPlaceRepository, guideRepository repository.IGuideRepository) ITicketService {
	return &TicketService{
		tr: ticketRepository,
		or: orderRepository,
		pr: placeRepository,
		gr: guideRepository,
	}
}

func (ts *TicketService) BuyTickets(param model.TicketBuy) (entity.Order, []entity.Ticket, error) {
	placeParam := model.PlaceParam{
		ID: param.PlaceID,
	}

	place, err := ts.pr.GetPlaceByID(placeParam.ID)
	if err != nil {
		return entity.Order{}, nil, err
	}
	var guide entity.Guide
	if param.GuideID != 0 {
		g, err := ts.gr.BookGuideByID(param.GuideID)
		if err != nil {
			return entity.Order{}, nil, err
		}
		guide = g
	}
	param.TicketPrice = place.Price
	newOrder := entity.Order{
		ID:         uuid.New(),
		UserID:     param.UserID,
		GuideID:    param.GuideID,
		TotalPrice: param.TicketPrice * param.TicketQuantity,
	}

	if param.GuideID != 0 {
		newOrder.GuideID = param.GuideID
		newOrder.TotalPrice += guide.GuidePrice
	}
	
	createdOrder, err := ts.or.CreateOrder(newOrder)
	if err != nil {
		return newOrder, nil, err
	}
	
	var ticketsID []uuid.UUID

	var tickets []entity.Ticket
	for i := 0; i < param.TicketQuantity; i++ {
		ticket := entity.Ticket{
			ID:          uuid.New(),
			UserID:      param.UserID,
			PlaceID:     param.PlaceID,
			TicketPrice: param.TicketPrice,
			OrderID:     newOrder.ID,
		}
		createdTicket, err := ts.tr.BuyTicket(ticket)
		if err != nil {
			return createdOrder, tickets, err
		}
		tickets = append(tickets, createdTicket)
		ticketsID = append(ticketsID, createdTicket.ID)
	}
	newOrder.Tickets = ticketsID
	err = ts.or.PassTicketToOrder(newOrder.ID, ticketsID)
	if err != nil {
		log.Println("Failed to associate tickets with order", err)
	}
	return createdOrder, tickets, nil
}

func (ts *TicketService) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	return ts.tr.GetTicketByID(param)
}

func (ts *TicketService) GetTicketByUserID(param model.TicketParam) ([]entity.Ticket, error) {
	return ts.tr.GetTicketByUserID(param)
}
