package repository

import (
	"INTERN_BCC/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	FindByUserID(userID uuid.UUID) ([]entity.Transaction, error)
	FindByOrderID(orderID string) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	FindSuccessByUserID(userID uuid.UUID) ([]entity.Transaction, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (tr *TransactionRepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
    if err := tr.db.Create(&transaction).Error; err != nil {
        return entity.Transaction{}, err
    }
    return transaction, nil
}

func (tr *TransactionRepository) FindByUserID(userID uuid.UUID) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	if err := tr.db.Where("user_id = ?", userID).Where("status = ?", "pending").Find(&transaction).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return transaction, nil
}

func (tr *TransactionRepository) FindByOrderID(orderID string) (entity.Transaction, error) {
	var transaction entity.Transaction
	if err := tr.db.Where("order_id = ?", orderID).First(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (tr *TransactionRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	if err := tr.db.Save(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (tr *TransactionRepository) FindSuccessByUserID(userID uuid.UUID) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	if err := tr.db.Where("user_id = ?", userID).Where("status = ?", "success").Find(&transaction).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return transaction, nil
}