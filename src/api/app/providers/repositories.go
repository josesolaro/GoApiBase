package providers

import (
	"ApiBase/src/api/configuration"
	"ApiBase/src/api/repositories"
)

func UserRepository() *repositories.UserRepository{
	db, _ := repositories.DataBaseConnection(configuration.GetConnection())
	return &repositories.UserRepository{Db: db}
}