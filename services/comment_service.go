package services

import (
	"final-project/interfaces"
	"final-project/models"
	"final-project/repository"
)

type CommentSvc struct {
	CommentRepo repository.CommentRepo
}

var ICommentRepo interfaces.ICommentRepo

func (s *CommentSvc) PostComment(comment *models.Comment) error {
	ICommentRepo = &s.CommentRepo
	return ICommentRepo.PostComment(comment)
}

func (s *CommentSvc) GetComments() ([]models.Comment, error) {
	ICommentRepo = &s.CommentRepo
	return ICommentRepo.GetComments()
}

func (s *CommentSvc) GetCommentByID(id uint) (models.Comment, error) {
	ICommentRepo = &s.CommentRepo
	return ICommentRepo.GetCommentByID(id)
}

func (s *CommentSvc) UpdateComment(id uint, comment *models.Comment) (*models.Comment, error) {
	ICommentRepo = &s.CommentRepo
	return ICommentRepo.UpdateComment(id, comment)
}

func (s *CommentSvc) DeleteComment(id uint) error {
	ICommentRepo = &s.CommentRepo
	return ICommentRepo.DeleteComment(id)
}
