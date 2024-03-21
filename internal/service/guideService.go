package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type IGuideService interface {
	CreateGuide(param model.CreateGuide) error
	GuidePatchID(param model.GuidePatch) (entity.Guide, error)
	GetGuideByID(param model.GuideParam) (entity.Guide, error)
	GetAllGuide(param model.GuideParam) ([]entity.Guide, error)
	PatchGuide(param model.GuidePatch) error
	BookGuideByID(param model.GuideBook) (entity.Guide, error)
}

type GuideService struct {
	gr repository.IGuideRepository
}

func NewGuideService(guideRepository repository.IGuideRepository) IGuideService {
	return &GuideService{
		gr: guideRepository,
	}
}

func (gs *GuideService) CreateGuide(param model.CreateGuide) error {
	guide := entity.Guide{
		ID:           param.ID,
		PlaceID:      param.PlaceID,
		Name:         param.Name,
		GuideDesc:    param.GuideDesc,
		GuidePrice:   param.GuidePrice,
		GuidePhoto:   param.GuidePhoto,
		GuideAddress: param.GuideAddress,
		GuideContact: param.GuideContact,
	}

	_, err := gs.gr.CreateGuide(guide)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GuideService) GetAllGuide(param model.GuideParam) ([]entity.Guide, error) {
	return gs.gr.GetAllGuide(param)
}

func (gs *GuideService) GuidePatchID(param model.GuidePatch) (entity.Guide, error) {
	return gs.gr.GuidePatchID(param.ID)
}

func (gs *GuideService) PatchGuide(param model.GuidePatch) error {
	existingGuide, err := gs.gr.GuidePatchID(param.ID)
	if err != nil {
		return err
	}

	if param.Name != "" {
		existingGuide.Name = param.Name
	}
	if param.GuideDesc != "" {
		existingGuide.GuideDesc = param.GuideDesc
	}
	if param.GuidePrice != 0 {
		existingGuide.GuidePrice = param.GuidePrice
	}
	if param.GuidePhoto != "" {
		existingGuide.GuidePhoto = param.GuidePhoto
	}
	if param.GuideAddress != "" {
		existingGuide.GuideAddress = param.GuideAddress
	}
	if param.GuideContact != "" {
		existingGuide.GuideContact = param.GuideContact
	}
	

	return gs.gr.PatchGuide(existingGuide)
}

func (gs *GuideService) GetGuideByID(param model.GuideParam) (entity.Guide, error) {
	return gs.gr.GetGuideByID(param.ID)
}

func (gs *GuideService) BookGuideByID(param model.GuideBook) (entity.Guide, error) {
	guider, err := gs.gr.GuidePatchID(param.ID)
	if err != nil {
		return guider, err
	}

	guider.Booked = true
	return guider, gs.gr.PatchGuide(guider)
}
