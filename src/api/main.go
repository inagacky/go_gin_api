package main

import (
	"github.com/inagacky/go_gin_api/src/api/configure/db"
	"github.com/inagacky/go_gin_api/src/api/configure/di"
	"github.com/inagacky/go_gin_api/src/api/configure/logger"
	"github.com/inagacky/go_gin_api/src/api/configure/routing"

	)

func main() {

	// DBの初期設定
	db.Init()
	// Loggerの初期設定
	logger.Init()

	container, err:= di.Init()
	if err != nil {
		panic(err)
	}

	// Routingの取得
	r := routing.GetRouting(container)

	r.Run(":8080")
}
