package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"

	"github.com/google/uuid"
)

type ITicketService interface {
	BuyTicket(param model.TicketBuy) ([]entity.Ticket, error)
	GetTicketByID(param model.TicketParam) (entity.Ticket, error)
}

type TicketService struct {
	tr repository.ITicketRepository
}

func NewTicketService(ticketRepository repository.ITicketRepository) ITicketService {
	return &TicketService{
		tr: ticketRepository,
	}
}

func (ts *TicketService) BuyTicket(param model.TicketBuy) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	for i := 0; i < param.TicketQuantity; i++ {
		ticket := entity.Ticket{
			ID:          uuid.New(),
			PlaceID:     param.PlaceID,
			UserID:      param.UserID,
			TicketPrice: param.TicketPrice,
		}
		createdTicket, err := ts.tr.BuyTicket(ticket)
		if err != nil {
			return tickets, err
		}
		tickets = append(tickets, createdTicket)
	}
	return tickets, nil
}

func (ts *TicketService) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	return ts.tr.GetTicketByID(param)
}
