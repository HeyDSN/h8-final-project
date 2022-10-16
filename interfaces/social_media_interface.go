package interfaces

import "final-project/models"

type ISocialMediaRepo interface {
	AddSocialMedia(socialMedia *models.SocialMedia) error
	GetSocialMedias() ([]models.SocialMedia, error)
	GetSocialMediaByID(id uint) (models.SocialMedia, error)
	UpdateSocialMedia(id uint, socialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(id uint) error
}
