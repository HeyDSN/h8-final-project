package interfaces

import "final-project/models"

type ICommentRepo interface {
	PostComment(photo *models.Comment) error
	GetComments() ([]models.Comment, error)
	GetCommentByID(id uint) (models.Comment, error)
	UpdateComment(id uint, photo *models.Comment) (*models.Comment, error)
	DeleteComment(id uint) error
}
