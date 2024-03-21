package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type IPlaceService interface {
	CreateData(param model.PlaceCreate) (error)
	GetPlaceByID(param model.PlaceParam) (entity.Place, error)
	GetAllPlace(param model.PlaceParam) ([]entity.Place, error)
	SearchPlace(param model.SearchPlace) ([]entity.Place, error)
	GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error)
}

type PlaceService struct {
	pr repository.IPlaceRepository
}

func NewPlaceService(placeRepository repository.IPlaceRepository) IPlaceService {
	return &PlaceService{
		pr: placeRepository,
	}
}

func (ps *PlaceService) CreateData(param model.PlaceCreate) (error) {
	place := entity.Place{
		ID:            param.ID,
		Name:          param.Name,
		CityID:        param.CityID,
		PlaceDesc:     param.PlaceDesc,
		PlaceAddress:  param.PlaceAddress,
		PlaceHistory:  param.PlaceHistory,
		PlaceFasil:    param.PlaceFasil,
		PlaceActivity: param.PlaceActivity,
		PlaceBestTime: param.PlaceBestTime,
		PlaceOpen:     param.PlaceOpen,
		PlacePrice:    param.PlacePrice,
		PlaceRules:    param.PlaceRules,
		PlaceEvent:    param.PlaceEvent,
		PlaceAward:    param.PlaceAward,
		PlaceImage:    param.PlaceImage,
	}

	_, err := ps.pr.CreateData(place)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PlaceService) GetPlaceByID(param model.PlaceParam) (entity.Place, error) {
	return ps.pr.GetPlaceByID(param.ID)
}

func (ps *PlaceService) GetAllPlace(param model.PlaceParam) ([]entity.Place, error) {
	return ps.pr.GetAllPlace(param)
}

func (ps *PlaceService) SearchPlace(param model.SearchPlace) ([]entity.Place, error) {
	return ps.pr.SearchPlace(param)
}

func (ps *PlaceService) GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error) {
	return ps.pr.GetPlaceByCityID(param)
}