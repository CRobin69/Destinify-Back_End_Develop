package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type ICulinaryService interface {
	CreateData(param model.CulinaryCreate) error
	GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error)
	GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error)
	//SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error)
}

type CulinarService struct {
	cr repository.ICulinaryRepository
}

func NewCulinaryService(culinaryRepository repository.ICulinaryRepository) ICulinaryService {
	return &CulinarService{
		cr: culinaryRepository,
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

	_, err := cs.cr.CreateData(culinary)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CulinarService) GetCulinaryByID(param model.CulinaryParam) (entity.Culinary, error) {
	return cs.cr.GetCulinaryByID(param)
}

func (cs *CulinarService) GetAllCulinary(param model.CulinaryParam) ([]entity.Culinary, error) {
	return cs.cr.GetAllCulinary(param)
}

// func (cs *CulinarService) SearchCulinary(param model.SearchCulinary) ([]entity.Culinary, error) {
// 	return cs.cr.SearchCulinary(param)
// }
