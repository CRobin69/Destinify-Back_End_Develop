package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type IPlaceRepository interface {
	CreateData(place entity.Place) (entity.Place, error)
	GetPlaceByID(param model.PlaceParam) (entity.Place, error)
	GetAllPlace(param model.PlaceParam) ([]entity.Place, error)
}

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) IPlaceRepository {
	return &PlaceRepository{db: db}
}

func (p *PlaceRepository) CreateData(place entity.Place) (entity.Place, error) {
	err := p.db.Debug().Create(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}

func (p *PlaceRepository) GetPlaceByID(param model.PlaceParam) (entity.Place, error) {
	place := entity.Place{}
	err := p.db.Debug().Where(&param).First(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}

func (p *PlaceRepository) GetAllPlace(param model.PlaceParam) ([]entity.Place, error) {
	var place []entity.Place
	err := p.db.Debug().Where(&param).Find(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}