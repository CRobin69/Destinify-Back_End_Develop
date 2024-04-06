package repository

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITicketRepository interface {
	BuyTicket(tickets entity.Ticket) (entity.Ticket, error)
	GetTicketByID(param model.TicketParam) (entity.Ticket, error)
	GetTicketByUserID(param model.TicketParam) ([]entity.Ticket, error)
	GetTicketByOrderID(orderID uuid.UUID) (entity.Ticket, error)
}

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) ITicketRepository {
	return &TicketRepository{db: db}
}

func (tr *TicketRepository) BuyTicket(ticket entity.Ticket) (entity.Ticket, error) {
	tx := tr.db.Begin()
	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return entity.Ticket{}, err
	}
	tx.Commit()
	return ticket, nil
}

func (tr *TicketRepository) GetTicketByID(param model.TicketParam) (entity.Ticket, error) {
	ticket := entity.Ticket{}
	err := tr.db.Debug().Where(&param).First(&ticket).Error
	if err != nil {
		return entity.Ticket{}, err
	}
	return ticket, nil
}

func (tr *TicketRepository) GetTicketByUserID(param model.TicketParam) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := tr.db.Debug().Where(&param).Find(&tickets).Error
	if err != nil {
		return []entity.Ticket{}, err
	}
	return tickets, nil
}

func (tr *TicketRepository) GetTicketByOrderID(orderID uuid.UUID) (entity.Ticket, error) {
	ticket := entity.Ticket{}
	err := tr.db.Debug().Where("order_id = ?", orderID).First(&ticket).Error
	if err != nil {
		return entity.Ticket{}, err
	}
	return ticket, nil
}
