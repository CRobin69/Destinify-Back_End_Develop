package service

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/repository"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
)

type ICulinaryService interface {
	CreateData(param model.CulinaryCreate) error
	GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error)
	GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error)
	SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error)
	GetCulinaryByCityID(param model.CulinaryParam) ([]entity.Culinary, error)
}

type CulinarService struct {
	culinaryRepository repository.ICulinaryRepository
}

func NewCulinaryService(culinaryRepository repository.ICulinaryRepository) ICulinaryService {
	return &CulinarService{
		culinaryRepository: culinaryRepository,
	}
}

func (cs *CulinarService) CreateData(param model.CulinaryCreate) error {
	culinary := entity.Culinary{
		ID:                 param.ID,
		Name:               param.Name,
		CityID:             param.CityID,
		CulinaryDesc:       param.CulinaryDesc,
		CulinaryAddress:    param.CulinaryAddress,
		CulinaryPriceRange: param.CulinaryPriceRange,
		CulinaryOpen:       param.CulinaryOpen,
		CulinaryClose:      param.CulinaryClose,
		CulinaryImage:      param.CulinaryImage,
		CulinaryAward:      param.CulinaryAward,
	}

	_, err := cs.culinaryRepository.CreateData(culinary)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CulinarService) GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error) {
	return cs.culinaryRepository.GetCulinaryByID(param)
}

func (cs *CulinarService) GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error) {
	return cs.culinaryRepository.GetAllCulinary(param)
}

func (cs *CulinarService) SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error) {
	return cs.culinaryRepository.SearchCulinary(param)
}

func (cs *CulinarService) GetCulinaryByCityID(param model.CulinaryParam) ([]entity.Culinary, error) {
	return cs.culinaryRepository.GetCulinaryByCityID(param)
}