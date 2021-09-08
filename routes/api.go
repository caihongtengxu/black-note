package routes

import (
	"net/http"

	"github.com/caihongtengxu/black-note/config"
	"github.com/caihongtengxu/black-note/pkg/logs"

	"github.com/gin-gonic/gin"
)

func InitApiRoute() *gin.Engine {
	router := gin.Default()

	// 设置日志中间件
	logs.LoggerInit(router)

	// 初始化路由
	InitApiV1Router(router)

	// 设置全局404请求
	router.NoRoute(func(c *gin.Context) {
		noRouteData := gin.H{"code": "NO_ROUTE_FIND", "message": "no page found"}
		c.JSON(http.StatusNotFound, noRouteData)
	})

	router.Run(":" + config.Param("APP_PORT"))

	return router
}
