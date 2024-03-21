package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"

	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrderByID(orderID uuid.UUID) (entity.Order, error)
	GetOrderByUserID(userID uuid.UUID) (entity.Order, error)
}

type OrderService struct {
	or repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{
		or: orderRepository,
	}
}

func (os *OrderService) CreateOrder(order entity.Order) (entity.Order, error) {
	createdOrder, err := os.or.CreateOrder(order)
	if err != nil {
		return entity.Order{}, err
	}
	return createdOrder, nil
}

func (os *OrderService) GetOrderByUserID(userID uuid.UUID) (entity.Order, error) {
	orders, err := os.or.GetOrdersByUserID(userID)
	if err != nil {
		return entity.Order{}, err
	}
	return orders, nil
}

func (os *OrderService) GetOrderByID(orderID uuid.UUID) (entity.Order, error) {
	order, err := os.or.GetOrderByID(orderID)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}