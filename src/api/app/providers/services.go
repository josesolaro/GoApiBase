package providers

import (
	"ApiBase/src/api/repositories"
	"ApiBase/src/api/services"
)

func UserService(r *repositories.UserRepository) *services.UserService{
	return &services.UserService{UserRepository: r}
}