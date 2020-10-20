package repositories

import (
	"ApiBase/src/api/configuration"
	"ApiBase/src/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

func DataBaseConnection(c *configuration.DB)(*gorm.DB, error){
	var (
		db *gorm.DB
		err error
	)
	switch os.Getenv("SCOPE"){
	case "production":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.Username, c.Password, c.Host,c.Name)
		db, err = gorm.Open("mysql",dsn)
	case "local":
		db, err = gorm.Open("sqlite3", ":memory:")
		break
	default:
		db, err = gorm.Open("sqlite3", ":memory:")
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
