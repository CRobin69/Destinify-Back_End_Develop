package service

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/internal/repository"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
)

type IPlaceService interface {
	CreateData(param model.PlaceCreate) (error)
	GetPlaceByID(param model.PlaceParam) (entity.Place, error)
	GetAllPlace(param model.PlaceParam) ([]entity.Place, error)
	SearchPlace(param model.SearchPlace) ([]entity.Place, error)
	GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error)
}

type PlaceService struct {
	placeRepository repository.IPlaceRepository
}

func NewPlaceService(placeRepository repository.IPlaceRepository) IPlaceService {
	return &PlaceService{
		placeRepository: placeRepository,
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

	_, err := ps.placeRepository.CreateData(place)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PlaceService) GetPlaceByID(param model.PlaceParam) (entity.Place, error) {
	return ps.placeRepository.GetPlaceByID(param.ID)
}

func (ps *PlaceService) GetAllPlace(param model.PlaceParam) ([]entity.Place, error) {
	return ps.placeRepository.GetAllPlace(param)
}

func (ps *PlaceService) SearchPlace(param model.SearchPlace) ([]entity.Place, error) {
	return ps.placeRepository.SearchPlace(param)
}

func (ps *PlaceService) GetPlaceByCityID(param model.PlaceParam) ([]entity.Place, error) {
	return ps.placeRepository.GetPlaceByCityID(param)
}