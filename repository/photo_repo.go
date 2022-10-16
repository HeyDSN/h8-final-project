package repository

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepo struct {
	Conn *gorm.DB
}

func (r *PhotoRepo) PostPhoto(photo *models.Photo) error {
	return r.Conn.Create(photo).Error
}

func (r *PhotoRepo) GetPhotos() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.Conn.Preload("User").Find(&photos).Error
	return photos, err
}

func (r *PhotoRepo) GetPhotoByID(id uint) (models.Photo, error) {
	var photo models.Photo
	err := r.Conn.Preload("User").First(&photo, id).Error
	return photo, err
}

func (r *PhotoRepo) UpdatePhoto(id uint, photo *models.Photo) (*models.Photo, error) {
	photoResult := &models.Photo{}
	err := r.Conn.Model(&photoResult).Clauses(clause.Returning{}).Where("id = ?", id).Updates(photo).Error
	if err != nil {
		return nil, err
	}
	return photoResult, nil
}

func (r *PhotoRepo) DeletePhoto(id uint) error {
	return r.Conn.Delete(&models.Photo{}, id).Error
}
