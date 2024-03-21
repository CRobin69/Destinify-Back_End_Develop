package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type IPlaceRepository interface {
	CreateData(place entity.Place) (entity.Place, error)
	GetPlaceByID(id uint) (entity.Place, error)
	GetAllPlace(param model.PlaceParam) ([]entity.Place, error)
	SearchPlace(param model.SearchPlace) ([]entity.Place, error)
	GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error)
}

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) IPlaceRepository {
	return &PlaceRepository{db: db}
}

func (pr *PlaceRepository) CreateData(place entity.Place) (entity.Place, error) {
	err := pr.db.Debug().Create(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}

func (pr *PlaceRepository) GetPlaceByID(id uint) (entity.Place, error) {
	place := entity.Place{}
	err := pr.db.Debug().Where(&id).First(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}

func (pr *PlaceRepository) GetAllPlace(param model.PlaceParam) ([]entity.Place, error) {
	var place []entity.Place
	err := pr.db.Debug().Where(&param).Find(&place).Error
	if err != nil {
		return place, err
	}

	return place, nil
}

func (pr *PlaceRepository) SearchPlace(param model.SearchPlace) ([]entity.Place, error) {
	var place []entity.Place
	if err := pr.db.Where("LOWER(name) LIKE LOWER(?)", "%"+param.Name+"%").Find(&place).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Place{}, nil
		}
		return nil, err
	}

	return place, nil
}

func (pr *PlaceRepository) GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error) {
	var place []entity.Place
	err := pr.db.Debug().Where(&param).Find(&place).Error
	if err != nil {
		return place, err
	}
	return place, nil
}