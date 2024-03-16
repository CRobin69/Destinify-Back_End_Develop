package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type ICityRepository interface {
	CreateCity(city entity.City) (entity.City, error)
	GetCity(param model.CityParam) (entity.City, error)
	GetAllCity(param model.CityParam) ([]entity.City, error)
	SearchCity(param model.SearchCity) ([]entity.City, error)
}

type CityRepository struct {
	db *gorm.DB
}


func NewCityRepository(db *gorm.DB) ICityRepository {
	return &CityRepository{db: db}
}

func (c *CityRepository) CreateCity(city entity.City) (entity.City, error) {
	err := c.db.Debug().Create(&city).Error
	if err != nil {
		return city, err
	}
	return city, nil
}

func (c *CityRepository) GetCity(param model.CityParam) (entity.City, error) {
	city := entity.City{}
	err := c.db.Debug().Where(&param).First(&city).Error
	if err != nil {
		return city,err
	}
	return city, nil
}

func (c *CityRepository) GetAllCity(param model.CityParam) ([]entity.City, error) {
	var city []entity.City
	err := c.db.Debug().Where(&param).Find(&city).Error
	if err != nil {
		return city, err
	}
	return city, nil
}

func (c *CityRepository) SearchCity(param model.SearchCity) ([]entity.City, error) {
	var city []entity.City
	if err := c.db.Where("LOWER(name) LIKE LOWER(?)", "%"+param.Name+"%").Find(&city).Error; err != nil {
		return nil, err
	}
	return city, nil
}