package configuration

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type DB struct{
	Username string
	Password string
	Host string
	Name string
}

type Connection struct{
	MaxIdleConnections int
	MaxOpenConnections int
	ConnMaxLifeTime time.Duration
}

func ConnectionConfig() *Connection{
	return &Connection{
		MaxIdleConnections: 10,
		MaxOpenConnections: 50,
		ConnMaxLifeTime: time.Second * 15,
	}
}

func GetConnectionConfiguration(scope string) *DB{
	var db DB
	switch scope{
		case "local":
			db.Name = "test"
			db.Host = "localhost:3306"
			db.Username = "root"
			db.Password = "root"
			break
	 	default:
			db.Name = "test"
			db.Host = "localhost:3306"
			db.Username = "root"
			db.Password = "root"
	}
	return &db
}