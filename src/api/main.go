package main

import (
	"github.com/inagacky/go_gin_api/src/api/configure/db"
	"github.com/inagacky/go_gin_api/src/api/configure/di"
	"github.com/inagacky/go_gin_api/src/api/configure/logger"
	"github.com/inagacky/go_gin_api/src/api/configure/routing"
	"log"
)

func main() {

	// DBの初期設定
	// Loggerの初期設定
	_, dbErr:= db.Init()
	if dbErr != nil {
		log.Panic(dbErr)
	}

	// Loggerの初期設定
	_, lErr:= logger.Init()
	if lErr != nil {
		log.Panic(lErr)
	}

	container, cErr:= di.Init()
	if cErr != nil {
		log.Panic(cErr)
	}
	// Routingの取得
	r := routing.GetRouting(container)

	r.Run(":8080")
}
