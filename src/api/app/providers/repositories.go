package providers

import (
	"ApiBase/src/api/configuration"
	"ApiBase/src/api/repositories"
	"os"
)

func UserRepository() *repositories.UserRepository{
	scope := os.Getenv("SCOPE")
	db, _ := repositories.DataBaseConnection(configuration.GetConnectionConfiguration(scope),scope)
	return &repositories.UserRepository{Db: db}
}