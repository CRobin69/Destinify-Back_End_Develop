package repository

import (
	"INTERN_BCC/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	FindByID(id uint) (entity.Transaction, error)
	FindByUserID(userID uuid.UUID) ([]entity.Transaction, error)
	FindByOrderID(orderID string) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
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

    if err := tr.db.Model(&transaction).Preload("Orders").Error; err != nil {
        return entity.Transaction{}, err
    }

    return transaction, nil
}

func (tr *TransactionRepository) FindAll() ([]entity.Transaction, error) {
	var data []entity.Transaction
	if err := tr.db.Find(&data).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return data, nil
}

func (tr *TransactionRepository) FindByID(id uint) (entity.Transaction, error) {
	var data entity.Transaction
	if err := tr.db.Where("id = ?", id).First(&data).Error; err != nil {
		return entity.Transaction{}, err
	}
	return data, nil
}

func (tr *TransactionRepository) FindByUserID(userID uuid.UUID) ([]entity.Transaction, error) {
	var data []entity.Transaction
	if err := tr.db.Where("user_id = ?", userID).Where("status = ?", "pending").Find(&data).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return data, nil
}

func (tr *TransactionRepository) FindByOrderID(orderID string) (entity.Transaction, error) {
	var data entity.Transaction
	if err := tr.db.Where("order_id = ?", orderID).First(&data).Error; err != nil {
		return entity.Transaction{}, err
	}
	return data, nil
}

func (tr *TransactionRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	if err := tr.db.Save(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}
