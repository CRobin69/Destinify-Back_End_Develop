package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"
	"gorm.io/gorm"
)

type ITicketRepository interface {
	CreateTicket(ticket entity.Ticket) (entity.Ticket, error)
	GetTicketByID(param model.TicketParam) (entity.Ticket, error)
}

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) ITicketRepository {
	return &TicketRepository{db: db}
}

func (t *TicketRepository) CreateTicket(ticket entity.Ticket) (entity.Ticket, error) {
	err := t.db.Debug().Create(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (t *TicketRepository) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	ticket := entity.Ticket{}
	err := t.db.Debug().Where(&param).First(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}
