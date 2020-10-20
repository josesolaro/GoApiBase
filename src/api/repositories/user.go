package repositories

import (
	"ApiBase/src/api/models"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface{
	Get(id int)(*models.User, error)
	Save(user *models.User)(*models.User, error)
}

type UserRepository struct{
	Db *gorm.DB
}

func (u *UserRepository) Get(id int)(*models.User, error){
	var (
		user models.User
	)

	res := u.Db.First(&user,"id=?",id)
	return &user, res.Error
}

func (u *UserRepository) Save(user *models.User)(*models.User, error){
	res := u.Db.Save(user)

	return user, res.Error
}