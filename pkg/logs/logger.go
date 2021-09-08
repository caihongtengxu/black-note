package logs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() gin.HandlerFunc {
	// 设置json日志格式
	log.Formatter = &logrus.JSONFormatter{}

	// 设置日志文件
	currentDay := time.Now().Format("2006-01-02")
	currentPath, _ := os.Getwd()
	filename := fmt.Sprintf("%s/storage/logs/%s-gin.log", currentPath, currentDay)

	// 创建文件
	f, _ := os.OpenFile(filename, os.O_APPEND, 0666)

	// err := os.Chmod(filename, 0755)

	// 设置日志输出到文件里
	log.Out = f

	// 设置gin的日志输出也到文件里
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	// gin.DefaultWriter = io.MultiWriter(f)

	// 设置debug以上的日志全记录
	log.Level = logrus.DebugLevel

	return func(c *gin.Context) {
		c.Next()

		log.WithFields(logrus.Fields{
			"code":    200,
			"request": "reggie",
		}).Info("hello world")
	}
}

func LoggerInit(router *gin.Engine) {
	currentDay := time.Now().Format("2006-01-02")
	currentPath, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/storage/logs/%s-gin.log", currentPath, currentDay)

	// 这里需要开启读写并追加，只追加的话不行，写不进去
	f, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)

	if f == nil {
		f, _ = os.Create(filePath)
	}

	fmt.Print(f)
	// 设置日志记录文件
	// gin.DefaultWriter = io.MultiWriter(f)
	// 设置记录文件并输出控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 格式化记录的格式，还需要优化
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}
