package service

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/repository"

	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	GetOrderByID(orderID uuid.UUID) (entity.Order, error)
	GetOrderByUserID(userID uuid.UUID) (entity.Order, error)
}

type OrderService struct {
	orderRepository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (os *OrderService) CreateOrder(order entity.Order) (entity.Order, error) {
	createdOrder, err := os.orderRepository.CreateOrder(order)
	if err != nil {
		return entity.Order{}, err
	}
	return createdOrder, nil
}

func (os *OrderService) GetOrderByUserID(userID uuid.UUID) (entity.Order, error) {
	orders, err := os.orderRepository.GetOrdersByUserID(userID)
	if err != nil {
		return entity.Order{}, err
	}
	return orders, nil
}

func (os *OrderService) GetOrderByID(orderID uuid.UUID) (entity.Order, error) {
	order, err := os.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}