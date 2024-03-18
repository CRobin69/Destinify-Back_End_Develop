package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type ITicketRepository interface {
	BuyTicket(ticket entity.Ticket) (entity.Ticket, error)
	GetTicketByID(param model.TicketParam) (entity.Ticket, error)
}

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) ITicketRepository {
	return &TicketRepository{db: db}
}

func (tr *TicketRepository) BuyTicket(ticket entity.Ticket) (entity.Ticket, error) {
	if err := tr.db.Create(&ticket).Error; err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (tr *TicketRepository) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	ticket := entity.Ticket{}
	err := tr.db.Debug().Where(&param).First(&ticket).Error
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}
