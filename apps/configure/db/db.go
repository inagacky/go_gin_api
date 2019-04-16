package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)
var engine *xorm.Engine

func Init() {
	var err error

	// データソースの定義
	dataSource := "root:sample@tcp(127.0.0.1:33006)/sample?parseTime=true&charset=utf8"
	engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}
}
// DBを返却
func GetDB() *xorm.Engine {
	return engine
}