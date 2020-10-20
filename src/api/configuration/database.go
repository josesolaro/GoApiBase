package configuration

import (
	"os"
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
		ConnMaxLifeTime: time.Minute * 2,
	}
}

func GetConnection() *DB{
	var db DB
	scope := os.Getenv("SCOPE")
	switch scope{
		case "local":
			db.Name = "test.db"
			db.Host = "localhost:3306"
			db.Username = "root"
			db.Password = "root"
			break
	 	default:
			db.Name = "test.db"
			db.Host = "localhost:3306"
			db.Username = "root"
			db.Password = "root"
	}
	return &db
}