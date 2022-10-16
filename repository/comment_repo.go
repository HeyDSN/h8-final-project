package repository

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo struct {
	Conn *gorm.DB
}

func (r *CommentRepo) PostComment(comment *models.Comment) error {
	return r.Conn.Create(comment).Error
}

func (r *CommentRepo) GetComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.Conn.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (r *CommentRepo) GetCommentByID(id uint) (models.Comment, error) {
	var comment models.Comment
	err := r.Conn.Preload("User").Preload("Photo").First(&comment, id).Error
	return comment, err
}

func (r *CommentRepo) UpdateComment(id uint, comment *models.Comment) (*models.Comment, error) {
	commentResult := &models.Comment{}
	err := r.Conn.Model(&commentResult).Clauses(clause.Returning{}).Where("id = ?", id).Updates(comment).Error
	if err != nil {
		return nil, err
	}
	return commentResult, nil
}

func (r *CommentRepo) DeleteComment(id uint) error {
	return r.Conn.Delete(&models.Comment{}, id).Error
}
