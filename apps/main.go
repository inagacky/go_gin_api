package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin_sample/apps/controllers"
	logger "github.com/sirupsen/logrus"
	"log"
	"os"
	"reflect"
	"strconv"
)

const (
	logFileName = "output.log"
)

var (
	appLog = logger.New()
)

// ロガーの設定
func LogSetting() (*os.File, error) {
	//ログフォーマットの指定
	appLog.Formatter = new(logger.TextFormatter)

	appLog.Level = logger.InfoLevel
	logPath := os.Getenv("GO_GIN_LOG_PATH")
	if logPath == "" {
		logPath = "/Users/d-inagaki/go/src/github.com/go_gin_sample/apps/log"
	}

	// ログファイルの作成・追記の準備
	logfile, err := os.OpenFile(logPath + "/" + logFileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		appLog.Info(err)
	} else {
		log.SetOutput(logfile)
	}

	// ログの書き込み先
	appLog.Out = logfile

	return logfile, err
}

func main() {

	LogSetting()

	router := gin.Default()
	// ユーザー取得
	router.GET("users/:id", func(c *gin.Context) {

		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			appLog.Error(err)
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			appLog.Error(err)
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}
		// データを処理する
		userCtrl := controllers.NewUser()
		result := userCtrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			appLog.Error(err)
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	router.Run(":8080")
}