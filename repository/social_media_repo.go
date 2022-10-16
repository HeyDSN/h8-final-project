package repository

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepo struct {
	Conn *gorm.DB
}

func (r *SocialMediaRepo) AddSocialMedia(socialMedia *models.SocialMedia) error {
	return r.Conn.Create(socialMedia).Error
}

func (r *SocialMediaRepo) GetSocialMedias() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := r.Conn.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (r *SocialMediaRepo) GetSocialMediaByID(id uint) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.Conn.Preload("User").First(&socialMedia, id).Error
	return socialMedia, err
}

func (r *SocialMediaRepo) UpdateSocialMedia(id uint, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	socialMediaResult := &models.SocialMedia{}
	err := r.Conn.Model(&socialMediaResult).Clauses(clause.Returning{}).Where("id = ?", id).Updates(socialMedia).Error
	if err != nil {
		return nil, err
	}
	return socialMediaResult, nil
}

func (r *SocialMediaRepo) DeleteSocialMedia(id uint) error {
	return r.Conn.Delete(&models.SocialMedia{}, id).Error
}
