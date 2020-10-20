package services

import (
	"ApiBase/src/api/models"
	"ApiBase/src/api/repositories"
)

type IUserService interface {
	GetUserById(id int) (*models.User, error)
	SaveUser(user *models.User) error
}

type UserService struct{
	UserRepository repositories.IUserRepository
}

func (u *UserService) GetUserById(id int)(*models.User, error){
	var (
		user *models.User
		err error
	)
	user, err = u.UserRepository.Get(id)
	return user, err
}
func (u *UserService) SaveUser(user *models.User)error{
	var err error

	user, err = u.UserRepository.Save(user)
	return err
}
