package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// 设置json日志格式
	log.Formatter = &logrus.JSONFormatter{}

	// 设置日志文件
	currentDay := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("./storage/logs/%s-gin.log", currentDay)

	// 创建文件
	f, _ := os.OpenFile(filename, os.O_APPEND, 0666)

	// 设置日志输出到文件里
	log.Out = f

	// 设置gin的日志输出也到文件里
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out

	// 设置debug以上的日志全记录
	log.Level = logrus.DebugLevel
}

func New() *logrus.Logger {
	return log
}

// Test::自定义格式化info
func Info(message string, jsonStr string) {
	var fields map[string]interface{}

	if len(jsonStr) > 0 {
		err := json.Unmarshal([]byte(jsonStr), &fields)

		if err != nil {
			log.Panic("日志参数格式转换错误")
			return
		}
	}

	log.WithFields(logrus.Fields(fields)).Info(message)
}
