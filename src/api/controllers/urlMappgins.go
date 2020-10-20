package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetURLMappgins(router *gin.Engine){
	router.HandleMethodNotAllowed = true
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true

	router.GET("/", ping)
	router.NoRoute(noRoute)
	router.NoMethod(notAllowed)

}

func ping(c *gin.Context){
	c.String(http.StatusOK, "pong")
}
func noRoute(c *gin.Context){
	c.JSON(http.StatusNotFound,map [string]string{"error":"Route not found"})
}
func notAllowed(c *gin.Context){
	c.JSON(http.StatusMethodNotAllowed,map [string]string{"error":"Method not allowed"})
}
