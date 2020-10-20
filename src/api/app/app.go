// +build wireinject

package app

import (
	"ApiBase/src/api/app/providers"
	"ApiBase/src/api/controllers"
	"ApiBase/src/api/repositories"
	"ApiBase/src/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var userRouteSet = wire.NewSet(
	providers.UserRepository,
	wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepository)),
	providers.UserService,
	wire.Bind(new(services.IUserService), new(*services.UserService)),
	providers.UserController,
	wire.Bind(new(controllers.IUserController), new(*controllers.UserController)),
)

var routerSet = wire.NewSet(userRouteSet, providers.ProviderRouter)
func Start() (*gin.Engine, error){
	panic(wire.Build(routerSet))
	return nil, nil
}