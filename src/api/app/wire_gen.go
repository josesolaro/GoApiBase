// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"ApiBase/src/api/app/providers"
	"ApiBase/src/api/controllers"
	"ApiBase/src/api/repositories"
	"ApiBase/src/api/services"
	"ApiBase/src/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injectors from app.go:

func Start() (*gin.Engine, error) {
	userRepository := providers.UserRepository()
	userService := providers.UserService(userRepository)
	logger := providers.Logger()
	userController := providers.UserController(userService, logger)
	engine := providers.ProviderRouter(userController)
	return engine, nil
}

// app.go:

var userRouteSet = wire.NewSet(providers.Logger, wire.Bind(new(utils.ILogger), new(*utils.Logger)), providers.UserRepository, wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)), providers.UserService, wire.Bind(new(services.IUserService), new(*services.UserService)), providers.UserController, wire.Bind(new(controllers.IUserController), new(*controllers.UserController)))

var routerSet = wire.NewSet(userRouteSet, providers.ProviderRouter)
