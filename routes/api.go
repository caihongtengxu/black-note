package routes

import (
	"net/http"

	"github.com/caihongtengxu/black-note/config"
	"github.com/gin-gonic/gin"
)

func InitApiRoute() *gin.Engine {
	router := gin.Default()

	InitApiV1Router(router)

	// 设置全局404请求
	router.NoRoute(func(c *gin.Context) {
		noRouteData := gin.H{"code": "NO_ROUTE_FIND", "message": "no page found"}
		c.JSON(http.StatusNotFound, noRouteData)
	})

	router.Run(":" + config.Param("APP_PORT"))

	return router
}
