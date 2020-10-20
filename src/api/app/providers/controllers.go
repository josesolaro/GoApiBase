package providers

import (
	"ApiBase/src/api/controllers"
	"ApiBase/src/api/services"
)

func UserController(s *services.UserService) *controllers.UserController{
	return &controllers.UserController{UserService: s}
}