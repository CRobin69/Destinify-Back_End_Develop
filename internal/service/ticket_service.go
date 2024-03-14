package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type ITicketService interface {
	CreateTicket(param model.TicketCreate) error
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

func (ts *TicketService) CreateTicket(param model.TicketCreate) error {
	totalPrice := param.TicketPrice * param.TicketQuantity
	ticket := entity.Ticket{
		ID:             param.ID,
		PlaceID:        param.PlaceID,
		TicketPrice:    param.TicketPrice,
		TicketDate:     param.TicketDate,
		TicketQuantity: param.TicketQuantity,
		TotalPrice:     totalPrice,
		UserID:         param.UserID,
	}

	_, err := ts.tr.CreateTicket(ticket)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TicketService) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	return ts.tr.GetTicketByID(param)
}
