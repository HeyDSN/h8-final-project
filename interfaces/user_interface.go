package interfaces

import "final-project/models"

type IUserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User) error
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}
