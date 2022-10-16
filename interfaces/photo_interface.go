package interfaces

import "final-project/models"

type IPhotoRepo interface {
	PostPhoto(photo *models.Photo) error
	GetPhotos() ([]models.Photo, error)
	GetPhotoByID(id uint) (models.Photo, error)
	UpdatePhoto(id uint, photo *models.Photo) (*models.Photo, error)
	DeletePhoto(id uint) error
}
