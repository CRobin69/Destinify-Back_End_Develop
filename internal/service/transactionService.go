package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
	"INTERN_BCC/pkg/helper"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type ITransactionService interface {
	CreateTransaction(userID uuid.UUID, transaction model.TransactionPost) (entity.Transaction, error)
	// FindAll() ([]entity.Transaction, error)
	// FindByID(id uint) (entity.Transaction, error)
	// FindByUserID(userID string) ([]entity.Transaction, error)
	// FindByOrderID(orderID string) (entity.Transaction, error)
	Update(orderID string) (entity.Transaction, error)
}

type TransactionService struct {
	tr             repository.ITransactionRepository
	midtransClient *helper.MdtClient
	ur             repository.IUserRepository
	or             repository.IOrderRepository
	gr             repository.IGuideRepository
	t              repository.ITicketRepository
	pr             repository.IPlaceRepository
}

func NewTransactionService(transactionRepository repository.ITransactionRepository, midtransClient *helper.MdtClient, userRepository repository.IUserRepository, orderRepository repository.IOrderRepository, ticketRepository repository.ITicketRepository, guideRepository repository.IGuideRepository, placeRepository repository.IPlaceRepository) ITransactionService {
	return &TransactionService{
		tr:             transactionRepository,
		midtransClient: midtransClient,
		ur:             userRepository,
		or:             orderRepository,
		gr:             guideRepository,
		t:              ticketRepository,
		pr:             placeRepository,
	}
}
func (ts *TransactionService) CreateTransaction(userID uuid.UUID, transaction model.TransactionPost) (entity.Transaction, error) {
	user, err := ts.ur.FindByID(userID)
	if err != nil {
		return entity.Transaction{}, err
	}

	order, err := ts.or.GetOrdersByUserID(user.ID)
	if err != nil {
		return entity.Transaction{}, err
	}

	paramTicket, err := ts.t.GetTicketByOrderID(order.ID)
	if err != nil {
		return entity.Transaction{}, err
	}

	places, err := ts.pr.GetPlaceByID(paramTicket.PlaceID)
	if err != nil {
		return entity.Transaction{}, err
	}
	log.Println("lawak", order)
	var guide entity.Guide
	if order.GuideID != 0 {
		guide, err = ts.gr.GetGuideByID(order.GuideID)
		if err != nil {
			return entity.Transaction{}, err
		}
	}

	guideName := ""
	guideIDString := ""
	guidePrice := 0
	if guide.ID != 0 {
		guideName = guide.Name
		guideIDString = strconv.FormatUint(uint64(guide.ID), 10)
		guidePrice = guide.GuidePrice
	}

	placeName := places.Name
	ticketPrice := places.Price
	orderID := order.ID.String()
	orderTickets := order.Tickets
	log.Println(orderID, orderTickets)
	formattedIDs := ts.or.ParseTicketIDs(orderTickets)
	transaction.Amount = (ticketPrice * len(formattedIDs)) + guidePrice
	taxTicket := float64(ticketPrice) * 0.05
	taxGuide := float64(guidePrice) * 0.15
	log.Println(taxGuide)
	taxTotal := taxTicket + taxGuide
	midtrans, err := ts.midtransClient.CreateTransaction(orderID, guideIDString, ticketPrice, formattedIDs, guidePrice, taxGuide, taxTicket, taxTotal, guideName, placeName, user.Email, user.Name, user.HP, transaction.Method)
	if err != nil {
		return entity.Transaction{}, err
	}

	data := entity.Transaction{
		OrderID:       order.ID,
		UserID:        userID,
		Amount:        transaction.Amount,
		TransactionID: midtrans.TransactionID,
		VANumber:      midtrans.VaNumbers[0].VANumber,
		Method:        transaction.Method,
		Status:        "pending",
		PlaceID:       places.ID,
	}

	create, err := ts.tr.CreateTransaction(data)
	if err != nil {
		return entity.Transaction{}, err
	}
	return create, nil
}

func (ts *TransactionService) Update(orderID string) (entity.Transaction, error) {
	data, err := ts.tr.FindByOrderID(orderID)
	if err != nil {
		return entity.Transaction{}, err
	}

	transStatus, err := ts.midtransClient.NotifHandler(orderID)
	if err != nil {
		return entity.Transaction{}, err
	} else {
		if transStatus != nil {
			if transStatus.TransactionStatus == "capture" {
				if transStatus.FraudStatus == "challenge" {
					data.Status = "challenge"
				} else if transStatus.FraudStatus == "accept" {
					data.Status = "success"
				}
			} else if transStatus.TransactionStatus == "settlement" {
				data.Status = "success"
			} else if transStatus.TransactionStatus == "deny" {
				data.Status = "deny"
			} else if transStatus.TransactionStatus == "cancel" || transStatus.TransactionStatus == "expire" {
				data.Status = "failure"
			} else if transStatus.TransactionStatus == "pending" {
				data.Status = "pending"
			}
		}
	}

	var result entity.Transaction
	result, err = ts.tr.Update(data)
	if err != nil {
		log.Println("=========================\n", err, "\n3=========================")
		return entity.Transaction{}, err
	}
	return result, nil
}
