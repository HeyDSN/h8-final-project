package services

import (
	"final-project/interfaces"
	"final-project/models"
	"final-project/repository"
)

type UserSvc struct {
	UserRepo repository.UserRepo
}

var IUserRepo interfaces.IUserRepo

func (s *UserSvc) CreateUser(user *models.User) error {
	IUserRepo = &s.UserRepo
	return IUserRepo.CreateUser(user)
}

func (s *UserSvc) GetUserByEmail(user *models.User) error {
	IUserRepo = &s.UserRepo
	return IUserRepo.GetUserByEmail(user)
}

func (s *UserSvc) UpdateUser(id uint, user *models.User) (*models.User, error) {
	IUserRepo = &s.UserRepo
	return IUserRepo.UpdateUser(id, user)
}

func (s *UserSvc) DeleteUser(id uint) error {
	IUserRepo = &s.UserRepo
	return IUserRepo.DeleteUser(id)
}
