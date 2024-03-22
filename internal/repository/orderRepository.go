package repository

import (
	"INTERN_BCC/entity"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrderByID(orderID uuid.UUID) (entity.Order, error)
	GetOrdersByUserID(userID uuid.UUID) (entity.Order, error)
	PassTicketToOrder(orderID uuid.UUID, ticketIDs []uuid.UUID) error
	ParseTicketIDs(ticketIDs []uuid.UUID) []string
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	tx := or.db.Begin()
	if order.GuideID != 0 {

		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			return entity.Order{}, err
		}
	} else {

		if err := tx.Omit("guide_id").Create(&order).Error; err != nil {
			tx.Rollback()
			return entity.Order{}, err
		}
	}

	tx.Commit()

	return order, nil
}

func (or *OrderRepository) GetOrderByID(orderID uuid.UUID) (entity.Order, error) {
	order := entity.Order{}
	err := or.db.Debug().Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (or *OrderRepository) PassTicketToOrder(orderID uuid.UUID, ticketIDs []uuid.UUID) error {

	array := pq.Array(ticketIDs)

	existingOrder := entity.Order{}
	if err := or.db.First(&existingOrder, "id = ?", orderID).Error; err != nil {
		return err
	}

	if err := or.db.Model(&existingOrder).Update("tickets", array).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) GetOrdersByUserID(userID uuid.UUID) (entity.Order, error) {
	var orders entity.Order
	var ticketIDs string
	var newestOrderID string

	if err := or.db.Model(&entity.Order{}).Where("user_id = ?", userID).Order("created_at ASC").Pluck("id", &newestOrderID).Error; err != nil {
		return entity.Order{}, err
	}

	if err := or.db.Model(&entity.Order{}).Where("id = ?", newestOrderID).Pluck("tickets", &ticketIDs).Error; err != nil {
		return entity.Order{}, err
	}

	if err := or.db.Select("user_id, guide_id, total_price, created_at, updated_at").Where("id = ?", newestOrderID).Find(&orders).Error; err != nil {
		return entity.Order{}, err
	}

	orderUUID, err := uuid.Parse(newestOrderID)
	if err != nil {
		return entity.Order{}, err
	}

	tickets, err := or.GetTicketIDsFromString(ticketIDs)
	if err != nil {
		return entity.Order{}, err
	}

	order := entity.Order{
		ID:         orderUUID,
		UserID:     userID,
		GuideID:    orders.GuideID,
		TotalPrice: orders.TotalPrice,
		Tickets:    tickets,
	}

	return order, nil
}

func (or *OrderRepository) GetTicketIDsFromString(formattedString string) ([]uuid.UUID, error) {
	formattedString = strings.Trim(formattedString, "{}")
	idStrings := strings.Split(formattedString, ",")

	var ticketIDs []uuid.UUID
	for _, idStr := range idStrings {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		ticketIDs = append(ticketIDs, id)
	}
	return ticketIDs, nil
}

func (or *OrderRepository) ParseTicketIDs(ticketIDs []uuid.UUID) []string {
	var formattedIDs []string
	for _, id := range ticketIDs {
		formattedIDs = append(formattedIDs, id.String())
	}
	return formattedIDs
}
