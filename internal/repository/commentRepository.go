package repository

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/entity"
	"github.com/CRobin69/Destinify-Back_End_Develop/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment entity.Comment) (entity.Comment, error)
	FindCommentByID(id uint) (entity.Comment, error)
	FindCommentByPlaceID(param model.CommentParam) ([]entity.Comment, error)
	FindCommentByUserID(param model.CommentParam) ([]entity.Comment, error)
	UpdateComment(comment entity.Comment) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository{
	return &CommentRepository{db : db}
}

func (cr *CommentRepository) CreateComment(comment entity.Comment)(entity.Comment, error){
	err := cr.db.Debug().Create(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}
func (cr *CommentRepository) FindCommentByID(id uint) (entity.Comment, error){
	var comment entity.Comment
	err := cr.db.Debug().Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (cr *CommentRepository) FindCommentByPlaceID(param model.CommentParam) ([]entity.Comment, error){
	var comment []entity.Comment
	err := cr.db.Debug().Where(&param).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (cr *CommentRepository) FindCommentByUserID(param model.CommentParam) ([]entity.Comment, error){
	var comment []entity.Comment
	err := cr.db.Debug().Where(&param).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (cr *CommentRepository) UpdateComment(comment entity.Comment) error{
	err := cr.db.Debug().Model(&comment).Where(comment.ID).Updates(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

