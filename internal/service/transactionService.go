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
	Update(orderID string) (entity.Transaction, error)
	GetSuccessByUserID(ID uuid.UUID) ([]entity.Transaction, error)
	CreateComment(param model.CommentCreate) error
}

type TransactionService struct {
	transactionRepository repository.ITransactionRepository
	midtransClient        *helper.MdtClient
	userRepository        repository.IUserRepository
	orderRepository       repository.IOrderRepository
	guideRepository       repository.IGuideRepository
	ticketRepository      repository.ITicketRepository
	placeRepository       repository.IPlaceRepository
	commentRepository     repository.ICommentRepository
}

func NewTransactionService(transactionRepository repository.ITransactionRepository, midtransClient *helper.MdtClient, userRepository repository.IUserRepository, orderRepository repository.IOrderRepository, ticketRepository repository.ITicketRepository, guideRepository repository.IGuideRepository, placeRepository repository.IPlaceRepository, commentRepository repository.ICommentRepository) ITransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		midtransClient:        midtransClient,
		userRepository:        userRepository,
		orderRepository:       orderRepository,
		guideRepository:       guideRepository,
		ticketRepository:      ticketRepository,
		placeRepository:       placeRepository,
		commentRepository:     commentRepository,
	}
}
func (ts *TransactionService) CreateTransaction(userID uuid.UUID, transaction model.TransactionPost) (entity.Transaction, error) {
	user, err := ts.userRepository.FindByID(userID)
	if err != nil {
		return entity.Transaction{}, err
	}

	order, err := ts.orderRepository.GetOrdersByUserID(user.ID)
	if err != nil {
		return entity.Transaction{}, err
	}

	paramTicket, err := ts.ticketRepository.GetTicketByOrderID(order.ID)
	if err != nil {
		return entity.Transaction{}, err
	}

	places, err := ts.placeRepository.GetPlaceByID(paramTicket.PlaceID)
	if err != nil {
		return entity.Transaction{}, err
	}

	var guide entity.Guide
	if order.GuideID != 0 {
		guide, err = ts.guideRepository.GetGuideByID(order.GuideID)
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
	formattedIDs := ts.orderRepository.ParseTicketIDs(orderTickets)

	transaction.Amount = (ticketPrice * len(formattedIDs)) + guidePrice
	taxTicket := float64(ticketPrice) * 0.05
	taxGuide := float64(guidePrice) * 0.15
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
		Place:         places,
		PlaceID:       places.ID,
	}

	create, err := ts.transactionRepository.CreateTransaction(data)
	if err != nil {
		return entity.Transaction{}, err
	}
	return create, nil
}

func (ts *TransactionService) Update(orderID string) (entity.Transaction, error) {
	data, err := ts.transactionRepository.FindByOrderID(orderID)
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
				comments := entity.Comment{
					UserID : data.UserID,
					TransactionID: data.ID,
					PlaceID : data.PlaceID,
				}
				_, err := ts.commentRepository.CreateComment(comments)
				if err != nil {
					return entity.Transaction{}, err
				}
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
	result, err = ts.transactionRepository.Update(data)
	if err != nil {
		log.Println("=========================\n", err, "\n3=========================")
		return entity.Transaction{}, err
	}
	return result, nil
}

func (ts *TransactionService) GetSuccessByUserID(ID uuid.UUID) ([]entity.Transaction, error) {
	transactions, err := ts.transactionRepository.FindSuccessByUserID(ID)
	if err != nil {
		return nil, err
	}
	for i, transaction := range transactions {
		place, err := ts.placeRepository.GetPlaceByID(transaction.PlaceID)
		if err != nil {
			return nil, err
		}
		transactions[i].Place = place
	}
	return transactions, nil
}

func (ts *TransactionService) CreateComment(param model.CommentCreate) error {
	successTransaction, err := ts.transactionRepository.FindSuccessByUserID(param.UserID)
	if err != nil {
		return err
	}

	for _, transaction := range successTransaction {
		comment := entity.Comment{
			UserID:     param.UserID,
			StarReview: param.StarReview,
			View:       param.View,
			Feedback:   param.Feedback,
			Opinion:    param.Opinion,
			PlaceID:    transaction.PlaceID,
		}
		_, err := ts.commentRepository.CreateComment(comment)
		if err != nil {
			return err
		}
	}
	return nil
}
