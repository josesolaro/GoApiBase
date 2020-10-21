package providers

import (
	"ApiBase/src/api/controllers"
	"ApiBase/src/api/services"
	"ApiBase/src/api/utils"
)

func UserController(s *services.UserService, l *utils.Logger) *controllers.UserController{
	return &controllers.UserController{UserService: s, Logger: l}
}