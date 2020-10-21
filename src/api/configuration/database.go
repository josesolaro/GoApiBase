package configuration

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func DatabaseMiddleware() gin.HandlerFunc {
	// Create connection pool
	var (
		db *gorm.DB
		err error
	)

	scope := os.Getenv("SCOPE")
	c := GetConnectionConfiguration(scope)
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
	if err != nil {
		fmt.Println(fmt.Sprintf(`[event:db_connection_error][error:"%s"]`, err.Error()))
		panic(err)
	}
	config := ConnectionConfig()
	db.DB().SetMaxIdleConns(config.MaxIdleConnections)
	db.DB().SetMaxOpenConns(config.MaxOpenConnections)
	db.DB().SetConnMaxLifetime(config.ConnMaxLifeTime)

	// Expose connection pool into request context
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}