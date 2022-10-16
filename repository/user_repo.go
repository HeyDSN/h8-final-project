package repository

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	Conn *gorm.DB
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.Conn.Create(user).Error
}

func (r *UserRepo) GetUserByEmail(user *models.User) error {
	return r.Conn.Where("email = ?", user.Email).First(user).Error
}

func (r *UserRepo) UpdateUser(id uint, user *models.User) (*models.User, error) {
	userResult := &models.User{}
	err := r.Conn.Model(&userResult).Clauses(clause.Returning{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (r *UserRepo) DeleteUser(id uint) error {
	// data := r.Conn.Where("id = ?", id).Delete(&models.User{})
	data := r.Conn.Select(clause.Associations).Delete(&models.User{ID: id})
	if data.Error != nil {
		return data.Error
	} else if data.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
