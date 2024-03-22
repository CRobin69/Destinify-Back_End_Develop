package service

import (
	"INTERN_BCC/entity"
	"INTERN_BCC/internal/repository"
	"INTERN_BCC/model"
)

type ICommentService interface {
	CreateComment(param model.CommentCreate) error
	FindCommentByPlaceID(param model.CommentParam) ([]entity.Comment, error)
	FindCommentByUserID(param model.CommentParam) ([]entity.Comment, error)
	UpdateComment(comment model.CommentCreate) error
}

type CommentService struct {
	commentRepository repository.ICommentRepository
}

func NewCommentService(commentRepository repository.ICommentRepository) ICommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (cs *CommentService) CreateComment(param model.CommentCreate) error {
	comment := entity.Comment{
		ID:         param.ID,
		UserID:     param.UserID,
		StarReview: param.StarReview,
		View:       param.View,
		Feedback:   param.Feedback,
		Opinion:    param.Opinion,
		PlaceID:    param.PlaceID,
	}

	_, err := cs.commentRepository.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CommentService) FindCommentByPlaceID(param model.CommentParam) ([]entity.Comment, error) {
	return cs.commentRepository.FindCommentByPlaceID(param)
}

func (cs *CommentService) FindCommentByUserID(param model.CommentParam) ([]entity.Comment, error) {
	return cs.commentRepository.FindCommentByUserID(param)
}

func (cs *CommentService) UpdateComment(param model.CommentCreate) error {
	existingComment, err := cs.commentRepository.FindCommentByID(param.ID)
	if err != nil {
		return err
	}
	if param.PlaceID == existingComment.PlaceID{
		return err
	}
	if param.StarReview != 0 {
		existingComment.StarReview = param.StarReview
	}
	if param.View != "" {
		existingComment.View = param.View
	}
	if param.Feedback != "" {
		existingComment.Feedback = param.Feedback
	}
	if param.Opinion != "" {
		existingComment.Opinion = param.Opinion
	}

	return cs.commentRepository.UpdateComment(existingComment)
}
