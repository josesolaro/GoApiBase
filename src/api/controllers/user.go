package controllers

import (
	"ApiBase/src/api/models"
	"ApiBase/src/api/services"
	"ApiBase/src/api/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface{
	GetUser(c *gin.Context)
	SaveUser(c *gin.Context)
}

type UserController struct{
	UserService services.IUserService
	Logger utils.ILogger
}
func (u *UserController) GetUser(c *gin.Context){
	var (
		userId int
		err error
	)

	if id := c.Param("id");len(id)==0{
		u.Logger.Error("Id is null ", nil)
		c.JSON(http.StatusBadRequest, map [string]string{"error":"id not set"})
		return
	}else{
		userId, err = strconv.Atoi(id)
	}

	if err != nil{
		u.Logger.Error("Id Not Number ", err)
		c.JSON(http.StatusBadRequest, map [string]string{"error":"id not number"})
		return
	}

	data, err := u.UserService.GetUserById(userId)
	if err != nil{
		u.Logger.Error("Id Not Number ", err)
		c.JSON(http.StatusNotFound,map [string]string{"error":"value not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (u *UserController) SaveUser(c *gin.Context){
	var user models.User
	body, _ := c.GetRawData()
	err := json.Unmarshal(body,&user)

	if err !=nil{
		u.Logger.Error("Could not parse body ", err)
		c.JSON(http.StatusBadRequest, map [string]string{"error":"could not parse body"})
	}
	err = u.UserService.SaveUser(&user)
	if err != nil{
		u.Logger.Error("Could not save ", err)
		c.JSON(http.StatusInternalServerError, map [string]string{"error":"could not save"})
	}

	c.JSON(http.StatusOK, map [string]string{"user_id":fmt.Sprintf("%v",user.ID)})
}