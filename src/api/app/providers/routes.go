package providers

import (
	"ApiBase/src/api/controllers"
	"github.com/gin-gonic/gin"
)

func ProviderRouter(userController *controllers.UserController) *gin.Engine{
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/:id",userController.GetUser)
		user.POST("",userController.SaveUser)
	}
	return router
}