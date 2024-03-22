package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type IGuideRepository interface {
	CreateGuide(guide entity.Guide) (entity.Guide, error)
	GetGuideByID(id uint) (entity.Guide, error)
	GuidePatchID(id uint) (entity.Guide, error)
	GetGuideByPlaceID(param model.GuideParam) ([]entity.Guide,error)
	GetAllGuide(param model.GuideParam) ([]entity.Guide, error)
	PatchGuide(guide entity.Guide) error
	BookGuideByID(id uint) (entity.Guide, error)
}

type GuideRepository struct {
	db *gorm.DB
}

func NewGuideRepository(db *gorm.DB) IGuideRepository {
	return &GuideRepository{db: db}
}

func (gr *GuideRepository) CreateGuide(guide entity.Guide) (entity.Guide, error) {
	if err := gr.db.Create(&guide).Error; err != nil {
		return guide, err
	}
	return guide, nil
}

func (gr *GuideRepository) GuidePatchID(id uint) (entity.Guide, error) {
	var guide entity.Guide
	err := gr.db.First(&guide, id).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}

func (gr *GuideRepository) PatchGuide(guide entity.Guide) error {
	return gr.db.Model(&guide).Updates(guide).Error
}

func (gr *GuideRepository) GetAllGuide(param model.GuideParam) ([]entity.Guide, error) {
	var guide []entity.Guide
	err := gr.db.Debug().Where(&param).Find(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}

func (gr *GuideRepository) GetGuideByID(id uint) (entity.Guide, error) {
	var guide entity.Guide
	if id == 0 {
		return entity.Guide{}, nil
	} 
	err := gr.db.Debug().Where("id = ?", id).First(&guide).Error
	if err != nil {
		return entity.Guide{}, err
	}
	return guide, nil
}

func (gr *GuideRepository) BookGuideByID(id uint) (entity.Guide, error) {
	var guide entity.Guide
	if err := gr.db.First(&guide, id).Error; err != nil {
		return entity.Guide{}, err
	}

	guide.Booked = true

	if err := gr.db.Save(&guide).Error; err != nil {
		return entity.Guide{}, err
	}

	if err := gr.db.First(&guide, id).Error; err != nil {
		return entity.Guide{}, err
	}

	return guide, nil
}

func (gr *GuideRepository)GetGuideByPlaceID(param model.GuideParam) ([]entity.Guide,error){
	var guide []entity.Guide
	err := gr.db.Debug().Where(&param).Find(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}
