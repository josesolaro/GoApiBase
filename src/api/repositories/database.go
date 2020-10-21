package repositories

import (
	"ApiBase/src/api/configuration"
	"ApiBase/src/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DataBaseConnection(c *configuration.DB,scope string)(*gorm.DB, error){
	var (
		db *gorm.DB
		err error
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.Username, c.Password, c.Host,c.Name)
	switch scope{
	case "production":
		db, err = gorm.Open("mysql",dsn)
	case "local":
		db, err = gorm.Open("mysql",dsn)
		break
	default:
		db, err = gorm.Open("mysql",dsn)
	}

	if err != nil{
		return db, err
	}
	connConfig := configuration.ConnectionConfig()

	db.DB().SetConnMaxLifetime(connConfig.ConnMaxLifeTime)
	db.DB().SetMaxIdleConns(connConfig.MaxIdleConnections)
	db.DB().SetMaxOpenConns(connConfig.MaxOpenConnections)

	db.AutoMigrate(&models.User{})

	return db, err
}
