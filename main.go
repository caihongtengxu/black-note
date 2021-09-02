package main

import (
	"github.com/caihongtengxu/black-note/pkg/logs"

	"github.com/caihongtengxu/black-note/routes"
)

var logger = logs.New()

func main() {
	logger.Info("开始初始化")
	routes.InitApiRoute()
	logger.Info("启动路由成功")
}
