package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type ICityService interface {
	CreateCity(param model.CityCreate) error
	GetCity(param model.CityParam) (entity.City, error)
	GetAllCity(param model.CityParam) ([]entity.City, error)
	SearchCity(param model.SearchCity) ([]entity.City, error)
}

type CityService struct {
	cityRepository repository.ICityRepository
}

func NewCityService(cityRepository repository.ICityRepository) ICityService {
	return &CityService{
		cityRepository: cityRepository,
	}
	
}

func (cs *CityService) CreateCity(param model.CityCreate) error {
	city := entity.City{
		ID:       param.ID,
		Name: 	  param.Name,
	}
	_, err := cs.cityRepository.CreateCity(city)
	if err != nil {
		return err
	}
	return nil
}
func (cs *CityService) GetCity(param model.CityParam) (entity.City, error) {
	return cs.cityRepository.GetCity(param)
}

func (cs *CityService) GetAllCity(param model.CityParam) ([]entity.City, error) {
	return cs.cityRepository.GetAllCity(param)
}

func (cs *CityService) SearchCity(param model.SearchCity) ([]entity.City, error) {
	return cs.cityRepository.SearchCity(param)
}
