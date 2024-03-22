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
}

type GuideService struct {
	guideRepository repository.IGuideRepository
}

func NewGuideService(guideRepository repository.IGuideRepository) IGuideService {
	return &GuideService{
		guideRepository: guideRepository,
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

	_, err := gs.guideRepository.CreateGuide(guide)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GuideService) GetAllGuide(param model.GuideParam) ([]entity.Guide, error) {
	return gs.guideRepository.GetAllGuide(param)
}

func (gs *GuideService) GuidePatchID(param model.GuidePatch) (entity.Guide, error) {
	return gs.guideRepository.GuidePatchID(param.ID)
}

func (gs *GuideService) PatchGuide(param model.GuidePatch) error {
	existingGuide, err := gs.guideRepository.GuidePatchID(param.ID)
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
	

	return gs.guideRepository.PatchGuide(existingGuide)
}

func (gs *GuideService) GetGuideByID(param model.GuideParam) (entity.Guide, error) {
	return gs.guideRepository.GetGuideByID(param.ID)
}
