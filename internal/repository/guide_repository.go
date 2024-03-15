package repository

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/model"

	"gorm.io/gorm"
)

type IGuideRepository interface {
	CreateGuide(guide entity.Guide) (entity.Guide, error)
	GetGuideByID(param model.GuideParam) (entity.Guide, error)
	GuidePatchID(id uint) (entity.Guide, error)
	GetAllGuide(param model.GuideParam) ([]entity.Guide, error)
	PatchGuide(guide entity.Guide) error
}

type GuideRepository struct {
	db *gorm.DB
}

func NewGuideRepository(db *gorm.DB) IGuideRepository {
	return &GuideRepository{db: db}
}

func (g *GuideRepository) CreateGuide(guide entity.Guide) (entity.Guide, error) {
	if err := g.db.Create(&guide).Error; err != nil {
		return guide, err
	}
	return guide, nil
}

func (gr *GuideRepository)GuidePatchID(id uint) (entity.Guide, error) {
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

func (g *GuideRepository) GetAllGuide(param model.GuideParam) ([]entity.Guide, error) {
	var guide []entity.Guide
	err := g.db.Debug().Where(&param).Find(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}

func (g *GuideRepository) GetGuideByID(param model.GuideParam) (entity.Guide, error) {
	guide := entity.Guide{}
	err := g.db.Debug().Where(&param).First(&guide).Error
	if err != nil {
		return guide, err
	}
	return guide, nil
}