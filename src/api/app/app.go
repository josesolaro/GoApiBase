package app

import (
	"github.com/gin-gonic/gin"
)

func Start() (*gin.Engine, error){
	router := gin.Default()
	return router, nil
}