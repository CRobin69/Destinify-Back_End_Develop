package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type ICityService interface {
	GetCity(param model.CityParam) (entity.City, error)
	GetAllCity(param model.CityParam) ([]entity.City, error)
}

type CityService struct {
	cr repository.ICityRepository
}

func NewCityService(cityRepository repository.ICityRepository) ICityService {
	// supabase supabase.Interface
	return &CityService{
		cr: cityRepository,
	}
	// supabase: supabase,
}

func (cs *CityService) GetCity(param model.CityParam) (entity.City, error) {
	return cs.cr.GetCity(param)
}

func (cs *CityService) GetAllCity(param model.CityParam) ([]entity.City, error) {
	return cs.cr.GetAllCity(param)
}