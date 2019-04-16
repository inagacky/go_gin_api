package main

import (
	"github.com/go_gin_sample/apps/configure/db"
	"github.com/go_gin_sample/apps/configure/logger"
	"github.com/go_gin_sample/apps/configure/routing"
)

func main() {

	// DBの初期設定
	db.Init()
	// Loggerの初期設定
	logger.Init()
	// Routingの取得
	r := routing.GetRouting()

	r.Run(":8080")
}
