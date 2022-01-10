package router

import (
	"github.com/gin-gonic/gin"
	"niceBackend/internal/resources/api/user"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("user/create", user.Register)
	}
}