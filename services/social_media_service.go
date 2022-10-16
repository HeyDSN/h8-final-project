package services

import (
	"final-project/interfaces"
	"final-project/models"
	"final-project/repository"
)

type SocialMediaSvc struct {
	SocialMediaRepo repository.SocialMediaRepo
}

var ISocialMediaRepo interfaces.ISocialMediaRepo

func (s *SocialMediaSvc) AddSocialMedia(socialMedia *models.SocialMedia) error {
	ISocialMediaRepo = &s.SocialMediaRepo
	return ISocialMediaRepo.AddSocialMedia(socialMedia)
}

func (s *SocialMediaSvc) GetSocialMedias() ([]models.SocialMedia, error) {
	ISocialMediaRepo = &s.SocialMediaRepo
	return ISocialMediaRepo.GetSocialMedias()
}

func (s *SocialMediaSvc) GetSocialMediaByID(id uint) (models.SocialMedia, error) {
	ISocialMediaRepo = &s.SocialMediaRepo
	return ISocialMediaRepo.GetSocialMediaByID(id)
}

func (s *SocialMediaSvc) UpdateSocialMedia(id uint, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	ISocialMediaRepo = &s.SocialMediaRepo
	return ISocialMediaRepo.UpdateSocialMedia(id, socialMedia)
}

func (s *SocialMediaSvc) DeleteSocialMedia(id uint) error {
	ISocialMediaRepo = &s.SocialMediaRepo
	return ISocialMediaRepo.DeleteSocialMedia(id)
}
