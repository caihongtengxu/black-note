package routes

import (
	"github.com/caihongtengxu/black-note/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func InitApiV1Router(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", controllers.Login)
	}
}
