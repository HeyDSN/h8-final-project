package services

import (
	"final-project/interfaces"
	"final-project/models"
	"final-project/repository"
)

type PhotoSvc struct {
	PhotoRepo repository.PhotoRepo
}

var IPhotoRepo interfaces.IPhotoRepo

func (s *PhotoSvc) PostPhoto(photo *models.Photo) error {
	IPhotoRepo = &s.PhotoRepo
	return IPhotoRepo.PostPhoto(photo)
}

func (s *PhotoSvc) GetPhotos() ([]models.Photo, error) {
	IPhotoRepo = &s.PhotoRepo
	return IPhotoRepo.GetPhotos()
}

func (s *PhotoSvc) GetPhotoByID(id uint) (models.Photo, error) {
	IPhotoRepo = &s.PhotoRepo
	return IPhotoRepo.GetPhotoByID(id)
}

func (s *PhotoSvc) UpdatePhoto(id uint, photo *models.Photo) (*models.Photo, error) {
	IPhotoRepo = &s.PhotoRepo
	return IPhotoRepo.UpdatePhoto(id, photo)
}

func (s *PhotoSvc) DeletePhoto(id uint) error {
	IPhotoRepo = &s.PhotoRepo
	return IPhotoRepo.DeletePhoto(id)
}
